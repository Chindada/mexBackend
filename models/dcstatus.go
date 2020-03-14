package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// DcStatus DcStatus
type DcStatus struct {
	ID         int `orm:"column(id)"`
	MacAddress string
	Timestamp  int64
	Status     int
	Sct        float64
	CycleTime  float64
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
}

// DcStatusTBName DcStatusTBName
const DcStatusTBName = BasicPrefix + "dc_status"

func init() {
	orm.RegisterModel(new(DcStatus))
}

// TableName TableName
func (c *DcStatus) TableName() string {
	return DcStatusTBName
}

// TableUnique TableUnique
func (c *DcStatus) TableUnique() [][]string {
	return [][]string{
		[]string{"MacAddress", "Timestamp"},
	}
}
