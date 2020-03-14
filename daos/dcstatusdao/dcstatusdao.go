package dcstatusdao

import (
	"ligang/models"

	"github.com/astaxie/beego/orm"
)

// InsertMultiStatus InsertMultiStatus
func InsertMultiStatus(multiStatus []models.DcStatus, o orm.Ormer) error {
	err := o.Begin()
	if _, err = o.InsertMulti(100, multiStatus); err != nil {
		err = o.Rollback()
		return err
	}
	err = o.Commit()
	return nil
}
