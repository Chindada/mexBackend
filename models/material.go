package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Material Material
type Material struct {
	ID        int       `orm:"column(id)"`
	IsActive  bool      `orm:"default(true)"`
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// MaterialTBName MaterialTBName
const MaterialTBName = BasicPrefix + "material"

func init() {
	orm.RegisterModel(new(Material))
}

// TableName TableName
func (c *Material) TableName() string {
	return MaterialTBName
}
