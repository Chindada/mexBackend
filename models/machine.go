package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Machine Machine
type Machine struct {
	ID        int    `orm:"column(id)"`
	Number    string `orm:"unique"`
	Model     string
	Mold      []*Mold   `orm:"rel(m2m);rel_table(rel_machine_mold)"`
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// MachineTBName MachineTBName
const MachineTBName = BasicPrefix + "machine"

func init() {
	orm.RegisterModel(new(Machine))
}

// TableName TableName
func (c *Machine) TableName() string {
	return MachineTBName
}
