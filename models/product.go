package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Product Product
type Product struct {
	ID        int       `orm:"column(id)"`
	IsActive  bool      `orm:"default(true)"`
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// ProductTBName ProductTBName
const ProductTBName = BasicPrefix + "product"

func init() {
	orm.RegisterModel(new(Product))
}

// TableName TableName
func (c *Product) TableName() string {
	return ProductTBName
}
