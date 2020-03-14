package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Storage Storage
type Storage struct {
	ID        int       `orm:"column(id)"`
	IsActive  bool      `orm:"default(true)"`
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// StorageTBName StorageTBName
const StorageTBName = ManufacturePrefix + "storage"

func init() {
	orm.RegisterModel(new(Storage))
}

// TableName TableName
func (c *Storage) TableName() string {
	return StorageTBName
}
