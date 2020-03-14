package dcstatusservice

import (
	"ligang/daos/dcstatusdao"
	"ligang/models"

	"github.com/astaxie/beego/orm"
)

// InsertMultiStatus InsertMultiStatus
func InsertMultiStatus(multiStatus []models.DcStatus) error {
	o := orm.NewOrm()
	if err := dcstatusdao.InsertMultiStatus(multiStatus, o); err != nil {
		return err
	}
	return nil
}
