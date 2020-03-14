package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ProductOrder ProductOrder
type ProductOrder struct {
	ID        int       `orm:"column(id)"`
	IsActive  bool      `orm:"default(true)"`
	Modified  time.Time `orm:"auto_now;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

// ProductOrderTBName ProductOrderTBName
const ProductOrderTBName = ManufacturePrefix + "product_order"

func init() {
	orm.RegisterModel(new(ProductOrder))
}

// TableName TableName
func (c *ProductOrder) TableName() string {
	return ProductOrderTBName
}
