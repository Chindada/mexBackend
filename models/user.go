package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User User
type User struct {
	ID        int    `orm:"column(id)"`
	Username  string `orm:"unique"`
	Password  string
	IsActive  bool `orm:"default(true)"`
	IsSuper   bool
	Group     int       `orm:"null"`
	Role      *Role     `orm:"null;rel(fk);on_delete(set_null)"`
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// UserTBName UserTBName
const UserTBName = AuthPrefix + "user"

func init() {
	orm.RegisterModel(new(User))
}

// TableName TableName
func (c *User) TableName() string {
	return UserTBName
}
