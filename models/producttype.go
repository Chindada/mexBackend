package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ProductType ProductType
type ProductType struct {
	ID        int       `orm:"column(id)"`
	IsActive  bool      `orm:"default(true)"`
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// ProductTypeTBName ProductTypeTBName
const ProductTypeTBName = BasicPrefix + "product_type"

func init() {
	orm.RegisterModel(new(ProductType))
}

// TableName TableName
func (c *ProductType) TableName() string {
	return ProductTypeTBName
}
