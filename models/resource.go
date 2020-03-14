package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Resource Resource
type Resource struct {
	ID        int       `orm:"column(id)"`
	Title     string    `orm:"unique"`
	Role      []*Role   `orm:"reverse(many)"`
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// ResourceTBName ResourceTBName
const ResourceTBName = AuthPrefix + "resource"

func init() {
	orm.RegisterModel(new(Resource))
}

// TableName TableName
func (c *Resource) TableName() string {
	return ResourceTBName
}
