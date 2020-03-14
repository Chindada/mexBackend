package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Role Role
type Role struct {
	ID        int         `orm:"column(id)"`
	Title     string      `orm:"unique"`
	IsActive  bool        `orm:"default(true)"`
	User      []*User     `orm:"reverse(many)"`
	Resource  []*Resource `orm:"rel(m2m);rel_table(rel_role_resource)"`
	Modified  time.Time   `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime)"`
}

// RoleTBName RoleTBName
const RoleTBName = AuthPrefix + "role"

func init() {
	orm.RegisterModel(new(Role))
}

// TableName TableName
func (c *Role) TableName() string {
	return RoleTBName
}
