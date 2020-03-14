package resourcedao

import (
	"ligang/daos/roledao"
	"ligang/models"

	"github.com/astaxie/beego/orm"
)

// GetResourceByRoleID GetResourceByRoleID
func GetResourceByRoleID(id int, o orm.Ormer) (map[string]bool, error) {
	resourceMap := make(map[string]bool)
	if role, err := roledao.GetRoleByID(id, o); err != nil {
		return nil, err
	} else if _, err := o.LoadRelated(&role, "Resource"); err != nil {
		return nil, err
	} else {
		for _, k := range role.Resource {
			resourceMap[k.Title] = true
		}
	}
	return resourceMap, nil
}

// CreateResource CreateResource
func CreateResource(new models.Resource, o orm.Ormer) error {
	_, err := o.Insert(&new)
	return err
}

// GetAllResource GetAllResource
func GetAllResource(o orm.Ormer) ([]models.Resource, error) {
	var allresource []models.Resource
	_, err := o.QueryTable(models.ResourceTBName).All(&allresource)
	return allresource, err
}

// UpdateResource UpdateResource
func UpdateResource(resource models.Resource, o orm.Ormer) error {
	if _, err := o.Update(&resource); err != nil {
		return err
	}
	return nil
}

// DeleteResource DeleteResource
func DeleteResource(id int, o orm.Ormer) error {
	if _, err := o.Delete(&models.Resource{ID: id}); err != nil {
		return err
	}
	return nil
}
