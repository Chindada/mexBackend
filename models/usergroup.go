package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// UserGroup UserGroup
type UserGroup struct {
	ID        int `orm:"column(id)"`
	Title     string
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// UserGroupTBName UserGroupTBName
const UserGroupTBName = AuthPrefix + "user_group"

func init() {
	orm.RegisterModel(new(UserGroup))
}

// TableName TableName
func (c *UserGroup) TableName() string {
	return UserGroupTBName
}
