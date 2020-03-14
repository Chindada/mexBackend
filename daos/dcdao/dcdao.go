package dcdao

import (
	"ligang/models"
	"sync"

	"github.com/astaxie/beego/orm"
)

// CreateDC CreateDC
func CreateDC(new models.Dc, o orm.Ormer) error {
	_, err := o.Insert(&new)
	return err
}

// GetAllDC GetAllDC
func GetAllDC(o orm.Ormer) ([]models.Dc, error) {
	var alldc []models.Dc
	_, err := o.QueryTable(models.DcTBName).All(&alldc)
	return alldc, err
}

// UpdateDC UpdateDC
func UpdateDC(dc models.Dc, o orm.Ormer) error {
	if _, err := o.Update(&dc); err != nil {
		return err
	}
	return nil
}

// DeleteDC DeleteDC
func DeleteDC(id int, o orm.Ormer) error {
	if _, err := o.Delete(&models.Dc{ID: id}); err != nil {
		return err
	}
	return nil
}

// GetLastStatusTime GetLastStatusTime
func GetLastStatusTime(o orm.Ormer) (*sync.Map, error) {
	var lastStatusTime sync.Map
	var alldcMac []string
	alldc, err := GetAllDC(o)
	if err != nil {
		return nil, err
	}
	for _, k := range alldc {
		alldcMac = append(alldcMac, k.MacAddress)
	}

	var lastStatusArray []models.DcStatus
	var lastStatus models.DcStatus
	for _, v := range alldcMac {
		o.QueryTable(models.DcStatusTBName).Filter("MacAddress__in", v).OrderBy("-Timestamp").One(&lastStatus)
		lastStatusArray = append(lastStatusArray, lastStatus)
	}

	for _, v := range lastStatusArray {
		lastStatusTime.Store(v.MacAddress, v.Timestamp)
	}
	return &lastStatusTime, nil
}

// GetLastDiTime GetLastDiTime
func GetLastDiTime(o orm.Ormer) (*sync.Map, error) {
	var lastDiTime sync.Map
	var alldcMac []string
	alldc, err := GetAllDC(o)
	if err != nil {
		return nil, err
	}
	for _, k := range alldc {
		alldcMac = append(alldcMac, k.MacAddress)
	}
	var lastDiArray []models.Di
	var lastDi models.Di
	for _, v := range alldcMac {
		o.QueryTable(models.DiTBName).Filter("MacAddress__in", v).OrderBy("-Timestamp").One(&lastDi)
		lastDiArray = append(lastDiArray, lastDi)
	}

	for _, v := range lastDiArray {
		// lastDiTime[v.MacAddress] = v.Timestamp
		lastDiTime.Store(v.MacAddress, v.Timestamp)
	}
	return &lastDiTime, nil
}
