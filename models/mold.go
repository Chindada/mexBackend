package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Mold Mold
type Mold struct {
	ID        int        `orm:"column(id)"`
	Number    string     `orm:"unique"`
	Machine   []*Machine `orm:"reverse(many)"`
	Modified  time.Time  `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time  `orm:"auto_now_add;type(datetime)"`
}

// MoldTBName MoldTBName
const MoldTBName = BasicPrefix + "mold"

func init() {
	orm.RegisterModel(new(Mold))
}

// TableName TableName
func (c *Mold) TableName() string {
	return MoldTBName
}
