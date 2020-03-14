package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Dc Dc
type Dc struct {
	ID          int `orm:"column(id)"`
	Group       int
	MacAddress  string `orm:"unique"`
	Token       string
	MaxInterval int64
	IdleTime    int64
	Modified    time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

// DcTBName DcTBName
const DcTBName = BasicPrefix + "dc"

func init() {
	orm.RegisterModel(new(Dc))
}

// TableName TableName
func (c *Dc) TableName() string {
	return DcTBName
}

// TableUnique TableUnique
func (c *Dc) TableUnique() [][]string {
	return [][]string{
		[]string{"Group", "MacAddress"},
	}
}
