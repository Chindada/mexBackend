package dcservice

import (
	"ligang/daos/dcdao"
	"ligang/models"
	"sync"

	"github.com/astaxie/beego/orm"
)

// CreateDC CreateDC
func CreateDC(new models.Dc) error {
	o := orm.NewOrm()
	err := dcdao.CreateDC(new, o)
	return err
}

// GetAllDC GetAllDC
func GetAllDC() ([]models.Dc, error) {
	o := orm.NewOrm()
	alldc, err := dcdao.GetAllDC(o)
	return alldc, err
}

// UpdateDC UpdateDC
func UpdateDC(dc models.Dc) error {
	o := orm.NewOrm()
	if err := dcdao.UpdateDC(dc, o); err != nil {
		return err
	}
	return nil
}

// DeleteDC DeleteDC
func DeleteDC(id int) error {
	o := orm.NewOrm()
	if err := dcdao.DeleteDC(id, o); err != nil {
		return err
	}
	return nil
}

// GetLastStatusTime GetLastStatusTime
func GetLastStatusTime() (*sync.Map, error) {
	o := orm.NewOrm()
	lastStatusTime, err := dcdao.GetLastStatusTime(o)
	return lastStatusTime, err
}

// GetLastDiTime GetLastDiTime
func GetLastDiTime() (*sync.Map, error) {
	o := orm.NewOrm()
	lastDiTime, err := dcdao.GetLastDiTime(o)
	return lastDiTime, err
}
