package resourceservice

import (
	"ligang/daos/resourcedao"
	"ligang/models"
	"sync"

	"github.com/astaxie/beego/orm"
)

// CreateResource CreateResource
func CreateResource(new models.Resource) error {
	o := orm.NewOrm()
	err := resourcedao.CreateResource(new, o)
	return err
}

// GetAllResource GetAllResource
func GetAllResource() ([]models.Resource, error) {
	o := orm.NewOrm()
	allresource, err := resourcedao.GetAllResource(o)
	return allresource, err
}

// StoreAllResourceMap StoreAllResourceMap
func StoreAllResourceMap() (*sync.Map, error) {
	var allresourceMap sync.Map
	o := orm.NewOrm()
	allresource, err := resourcedao.GetAllResource(o)
	for _, k := range allresource {
		allresourceMap.Store(k.Title, true)
	}
	return &allresourceMap, err
}

// UpdateResource UpdateResource
func UpdateResource(resource models.Resource) error {
	o := orm.NewOrm()
	if err := resourcedao.UpdateResource(resource, o); err != nil {
		return err
	}
	return nil
}

// DeleteResource DeleteResource
func DeleteResource(id int) error {
	o := orm.NewOrm()
	if err := resourcedao.DeleteResource(id, o); err != nil {
		return err
	}
	return nil
}
