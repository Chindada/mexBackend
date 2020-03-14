package didao

import (
	"ligang/models"

	"github.com/astaxie/beego/orm"
)

// InsertMultiDi InsertMultiDi
func InsertMultiDi(dis []models.Di, o orm.Ormer) error {
	err := o.Begin()
	if _, err = o.InsertMulti(100, dis); err != nil {
		err = o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}

// GetAllUnCleanDi GetAllUnCleanDi
func GetAllUnCleanDi(dc models.Dc, afterTime int64, o orm.Ormer) ([]models.Di, error) {
	var unCleanDi []models.Di
	if _, err := o.QueryTable(models.DiTBName).Filter("MacAddress", dc.MacAddress).Filter("Timestamp__gt", afterTime).All(&unCleanDi); err != nil {
		return nil, err
	}
	return unCleanDi, nil
}
