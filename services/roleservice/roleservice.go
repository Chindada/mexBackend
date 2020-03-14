package roleservice

import (
	"ligang/daos/resourcedao"
	"ligang/daos/roledao"
	"ligang/models"

	"github.com/astaxie/beego/orm"
)

// CreateNewRole CreateNewRole
func CreateNewRole(new models.Role) error {
	o := orm.NewOrm()
	daoErr := roledao.CreateNewRole(new, o)
	return daoErr
}

// GetAllRole GetAllRole
func GetAllRole() ([]models.Role, error) {
	o := orm.NewOrm()
	allrole, daoErr := roledao.GetAllRole(o)
	return allrole, daoErr
}

// GetOneRole GetOneRole
func GetOneRole() {

}

// UpateRole UpateRole
func UpateRole(role models.Role) error {
	o := orm.NewOrm()
	if err := roledao.UpateRole(role, o); err != nil {
		return err
	}
	return nil
}

// DeleteRole DeleteRole
func DeleteRole(id int) error {
	o := orm.NewOrm()
	if err := roledao.DeleteRole(id, o); err != nil {
		return err
	}
	return nil
}

// AddRoleResourceRel AddRoleResourceRel
func AddRoleResourceRel(role models.Role) error {
	if role.Resource == nil {
		return nil
	}
	o := orm.NewOrm()
	m2m := o.QueryM2M(&role, "Resource")
	var resourceIDArry []int
	if role.Resource != nil {
		for _, k := range role.Resource {
			resourceIDArry = append(resourceIDArry, k.ID)
		}
	}

	var resource []models.Resource
	originalresource, daoErr := resourcedao.GetResourceByRoleID(role.ID, o)
	if daoErr != nil {
		return daoErr
	}
	o.QueryTable(models.ResourceTBName).Filter("ID__in", resourceIDArry).All(&resource)
	var addresource []models.Resource
	for _, k := range resource {
		if !originalresource[k.Title] {
			addresource = append(addresource, k)
		}
	}
	if addresource != nil {
		m2m.Add(addresource)
	}
	return nil
}
