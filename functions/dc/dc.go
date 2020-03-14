package dc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"ligang/functions/arptable"
	"ligang/models"
	"ligang/services/dcservice"
	"ligang/services/dcstatusservice"
	"ligang/services/diservice"
	"ligang/utils"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego"
)

var tr *http.Transport
var iptable *sync.Map
var dcFetchLockMap sync.Map
var diMap sync.Map
var lastStatusTimeMap *sync.Map
var lastDiTimeMap sync.Map
var allDc []models.Dc

// InitDc InitDc
func InitDc() {
	fmt.Println("DC Init")
	tr = &http.Transport{
		DisableKeepAlives: true,
		MaxIdleConns:      -1,
	}
	getDcIP()
	go initAllDc()
	if lastStatusTime, err := dcservice.GetLastStatusTime(); err != nil {
		utils.LogCritical(err)
	} else {
		lastStatusTimeMap = lastStatusTime
	}
	initDiMap()
	for _, k := range allDc {
		lastditime, _ := lastStatusTimeMap.Load(k.MacAddress)
		lastDiTimeMap.Store(k.MacAddress, lastditime.(int64))
		diToStatus(k)
	}
}

func initAllDc() {
	for {
		dcs, err := dcservice.GetAllDC()
		if err != nil {
			utils.LogCritical(err)
		}
		for _, k := range dcs {
			dcFetchLockMap.Store(k.MacAddress, false)
		}
		allDc = dcs
		time.Sleep(60 * time.Second)
	}
}

func getDcIP() {
	iptable = arptable.IPTable()
	iptable.Store("00D0C9FD7469", "220.130.131.251:9091")
	iptable.Store("00D0C9E43B8C", "220.130.131.251:9092")
	iptable.Store("00D0C9FD7A1A", "220.130.131.251:9093")
	iptable.Store("00D0C9E43B84", "220.130.131.251:9094")

}

func initDiMap() {
	for _, k := range allDc {
		interfaceTime, _ := lastStatusTimeMap.LoadOrStore(k.MacAddress, 0)
		afterTime, _ := interfaceTime.(int64)
		if afterTime != 0 {
			unCleanDi, err := diservice.GetAllUnCleanDi(k, afterTime)
			if err != nil {
				utils.LogError(err)
				return
			}
			diMap.Store(k.MacAddress, unCleanDi)
		}
	}
}

// Loop Loop
func Loop() {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		for _, k := range allDc {
			interfaceLock, _ := dcFetchLockMap.Load(k.MacAddress)
			dclock, _ := interfaceLock.(bool)
			interfaceIP, _ := iptable.Load(k.MacAddress)
			ip, _ := interfaceIP.(string)
			if !dclock && ip != "" {
				go func(k models.Dc) {
					Fetch(k, ip)
					diToStatus(k)
				}(k)
			}
		}
	}
}

// Fetch Fetch
func Fetch(dc models.Dc, IP string) {
	dcFetchLockMap.Store(dc.MacAddress, true)
	interfaceTime, _ := lastDiTimeMap.LoadOrStore(dc.MacAddress, 0)
	startTime, _ := interfaceTime.(int64)
	stst := startTime/1000 + 1
	if stst == 1 {
		stst = time.Now().Add(-1802 * time.Second).Unix()
	}
	defer func() {
		dcFetchLockMap.Store(dc.MacAddress, false)
	}()
	min, max, err := checkgetDcLog(dc, IP)
	if err != nil {
		return
	} else if stst < min {
		stst = min
	}
	stend := max
	if stend-stst > dc.MaxInterval {
		stend = stst + dc.MaxInterval
	}
	if stst > max {
		return
	}
	if err := patchDc(dc, IP, stst, stend); err != nil {
		return
	} else if err := getDcLog(dc, IP); err != nil {
		return
	} else if returnData, err := getDcData(dc, IP); err != nil {
		return
	} else {
		alldata := convertLogToDi(returnData)
		if len(alldata) != 0 {
			if err := diservice.InsertMultiDi(alldata); err != nil {
				return
			}
			lastDiTimeMap.Store(dc.MacAddress, stend*1000)
			val, _ := diMap.Load(dc.MacAddress)
			originalArray, _ := val.([]models.Di)
			originalArray = append(originalArray, alldata...)
			diMap.Store(dc.MacAddress, originalArray)
		}
	}
}

func diToStatus(dc models.Dc) {
	var allstatus []models.DcStatus
	now := time.Now().Unix()
	status := models.DcStatus{}
	interfaceDis, _ := diMap.Load(dc.MacAddress)
	dis, _ := interfaceDis.([]models.Di)
	cycle := cycleStep{}
	for _, k := range dis {
		interfaceTime, _ := lastStatusTimeMap.LoadOrStore(k.MacAddress, 0)
		thisTime, _ := interfaceTime.(int64)
		if k.Timestamp < thisTime {
			continue
		}
		if cycle.s1.Timestamp != 0 && cycle.s2.Timestamp != 0 && cycle.s3.Timestamp != 0 {
			status.CycleTime = float64(cycle.s3.Timestamp-cycle.s1.Timestamp) / float64(1000)
			status.Timestamp = int64(cycle.s3.Timestamp)
			status.MacAddress = k.MacAddress
			status.Status = 5
			allstatus = append(allstatus, status)
			lastStatusTimeMap.Store(k.MacAddress, cycle.s3.Timestamp)
			beego.Informational("CT", k.MacAddress, status.CycleTime)
			cycle.s1 = cycle.s3
			cycle.s2.Timestamp, cycle.s3.Timestamp = 0, 0
		}
		if cycle.s1.Timestamp == 0 && k.Di2 == 1 {
			cycle.s1 = k
		}
		if cycle.s2.Timestamp == 0 && cycle.s1.Timestamp != 0 && k.Di1 == 1 {
			if cycle.s2.Timestamp-cycle.s1.Timestamp < dc.IdleTime {
				cycle.s2 = k
			} else {
				status.CycleTime = 0
				status.Timestamp = int64(cycle.s2.Timestamp)
				status.MacAddress = k.MacAddress
				status.Status = 3
				allstatus = append(allstatus, status)
				cycle.s1.Timestamp, cycle.s2.Timestamp = 0, 0
				lastStatusTimeMap.Store(k.MacAddress, cycle.s2.Timestamp)
			}
		}
		if cycle.s3.Timestamp == 0 && cycle.s1.Timestamp != 0 && cycle.s2.Timestamp != 0 && k.Di2 == 1 {
			if cycle.s3.Timestamp-cycle.s2.Timestamp < dc.IdleTime {
				cycle.s3 = k
			} else {
				status.CycleTime = 0
				status.Timestamp = int64(cycle.s3.Timestamp)
				status.MacAddress = k.MacAddress
				status.Status = 3
				allstatus = append(allstatus, status)
				cycle.s1 = k
				cycle.s2.Timestamp = 0
				lastStatusTimeMap.Store(k.MacAddress, cycle.s3.Timestamp)
			}
		}
	}
	var remainedDiArray []models.Di
	if cycle.s1.Timestamp != 0 && cycle.s2.Timestamp != 0 && cycle.s3.Timestamp != 0 {
		return
	}
	if cycle.s2.Timestamp == 0 && cycle.s3.Timestamp == 0 && len(dis) > 0 {
		if now-cycle.s1.Timestamp < dc.MaxInterval {
			remainedDiArray = append(remainedDiArray, cycle.s1)
		} else {
			status.CycleTime = 0
			status.Timestamp = now * 1000
			status.MacAddress = dc.MacAddress
			status.Status = 3
			allstatus = append(allstatus, status)
			lastStatusTimeMap.Store(dc.MacAddress, now*1000)
		}
	} else if cycle.s2.Timestamp != 0 && cycle.s3.Timestamp == 0 {
		if now-cycle.s2.Timestamp < dc.MaxInterval {
			remainedDiArray = append(remainedDiArray, cycle.s1, cycle.s2)
		} else {
			status.CycleTime = 0
			status.Timestamp = now * 1000
			status.MacAddress = dc.MacAddress
			status.Status = 3
			allstatus = append(allstatus, status)
			lastStatusTimeMap.Store(dc.MacAddress, now*1000)
		}
	}
	if len(allstatus) != 0 {
		if err := dcstatusservice.InsertMultiStatus(allstatus); err != nil {
			return
		}
		// diMap.Delete(dc.MacAddress)
		diMap.Store(dc.MacAddress, remainedDiArray)
	}
}

func convertLogToDi(data WiseOriginalData) []models.Di {
	var dis []models.Di
	for _, k := range data.LogMsg {
		di := models.Di{}
		di.MacAddress = k.UID[10:22]
		logTimeInt64, _ := strconv.ParseInt(k.TIM, 10, 64)
		di.Timestamp = logTimeInt64*1000 + k.SysTk%1000/100*100
		di.Di0 = k.Record[0][3]
		di.Di1 = k.Record[1][3]
		di.Di2 = k.Record[2][3]
		di.Di3 = k.Record[3][3]
		di.Di4 = k.Record[4][3]
		di.Di5 = k.Record[5][3]
		di.Di6 = k.Record[6][3]
		di.Di7 = k.Record[7][3]
		dis = append(dis, di)
	}
	return dis
}

func checkgetDcLog(dc models.Dc, IP string) (min, max int64, err error) {
	url := "http://" + IP + "/log_output"
	req, _ := http.NewRequest("GET", url, nil)
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", dc.Token)
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	} else if err != nil {
		utils.LogError(err)
		return 0, 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.LogError(err)
		return 0, 0, err
	}
	var data getlogBody
	if err := json.Unmarshal(body, &data); err != nil {
		utils.LogError(err)
		return 0, 0, err
	}
	min = data.TLst
	max = data.TFst
	return min, max, nil
}

func putDc(dc models.Dc, IP string) error {
	url := "http://" + IP + "/log_output"
	puBody := putBody{1, 0, 0, 1}
	putcontent, _ := json.Marshal(puBody)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(putcontent))
	if err != nil {
		utils.LogError(err)
		return err
	}
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", dc.Token)
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	} else if err != nil {
		utils.LogError(err)
		return err
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}

func patchDc(dc models.Dc, IP string, stst, stend int64) error {
	url := "http://" + IP + "/log_output"
	pBody := patchBody{stst, stend}
	patchcontent, _ := json.Marshal(pBody)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(patchcontent))
	if err != nil {
		utils.LogError(err)
		return err
	}
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", dc.Token)
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	} else if err != nil {
		utils.LogError(err)
		return err
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}

func getDcLog(dc models.Dc, IP string) error {
	url := "http://" + IP + "/log_output"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		utils.LogError(err)
		return err
	}
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", dc.Token)
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	} else if err != nil {
		utils.LogError(err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.LogError(err)
		return err
	}
	var data getlogBody
	if err := json.Unmarshal(body, &data); err != nil {
		utils.LogError(err)
	}
	if data.UID != 1 || data.MAC != 0 || data.TmF != 0 || data.Fltr != 1 {
		if err := putDc(dc, IP); err != nil {
			return err
		}
	}
	return nil
}

func getDcData(dc models.Dc, IP string) (WiseOriginalData, error) {
	var data WiseOriginalData
	url := "http://" + IP + "/log_message"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		utils.LogError(err)
		return data, err
	}
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", dc.Token)
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	} else if err != nil {
		utils.LogError(err)
		return data, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.LogError(err)
		return data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		utils.LogError(err)
		return data, err
	}
	return data, nil
}

type patchBody struct {
	TSt  int64
	TEnd int64
}

type putBody struct {
	UID  int
	MAC  int
	TmF  int
	Fltr int
}

type getlogBody struct {
	UID   int64
	MAC   int64
	TmF   int64
	SysTk int64
	Fltr  int64
	TSt   int64
	TEnd  int64
	Amt   int64
	Total int64
	TLst  int64
	TFst  int64
}

// WiseOriginalData WiseOriginalData
type WiseOriginalData struct {
	LogMsg []struct {
		PE     int
		UID    string
		TIM    string
		SysTk  int64
		Record [][]int
	}
}

type cycleStep struct {
	s1 models.Di
	s2 models.Di
	s3 models.Di
}
