package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Di Di
type Di struct {
	ID         int `orm:"column(id)"`
	MacAddress string
	Timestamp  int64
	Di0        int
	Di1        int
	Di2        int
	Di3        int
	Di4        int
	Di5        int
	Di6        int
	Di7        int
	AutoInsert bool
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
}

// DiTBName DiTBName
const DiTBName = BasicPrefix + "dc_di"

func init() {
	orm.RegisterModel(new(Di))
}

// TableName TableName
func (c *Di) TableName() string {
	return DiTBName
}

// TableIndex TableIndex
func (c *Di) TableIndex() [][]string {
	return [][]string{
		[]string{"Timestamp"},
	}
}

// TableUnique TableUnique
func (c *Di) TableUnique() [][]string {
	return [][]string{
		[]string{"MacAddress", "Timestamp"},
	}
}
