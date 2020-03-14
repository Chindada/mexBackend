package diservice

import (
	"ligang/daos/didao"
	"ligang/models"

	"github.com/astaxie/beego/orm"
)

// InsertMultiDi InsertMultiDi
func InsertMultiDi(dis []models.Di) error {
	o := orm.NewOrm()
	if err := didao.InsertMultiDi(dis, o); err != nil {
		return err
	}
	return nil
}

// GetAllUnCleanDi GetAllUnCleanDi
func GetAllUnCleanDi(dc models.Dc, afterTime int64) ([]models.Di, error) {
	o := orm.NewOrm()
	unCleanDi, err := didao.GetAllUnCleanDi(dc, afterTime, o)
	if err != nil {
		return nil, err
	}
	return unCleanDi, nil
}
