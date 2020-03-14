package roledao

import (
	"ligang/models"

	"github.com/astaxie/beego/orm"
)

// CreateNewRole CreateNewRole
func CreateNewRole(new models.Role, o orm.Ormer) error {
	_, err := o.Insert(&new)
	return err
}

// GetAllRole GetAllRole
func GetAllRole(o orm.Ormer) ([]models.Role, error) {
	var allrole []models.Role
	_, err := o.QueryTable(models.RoleTBName).All(&allrole)
	return allrole, err
}

// GetRoleByID GetRoleByID
func GetRoleByID(id int, o orm.Ormer) (models.Role, error) {
	role := models.Role{ID: id}
	err := o.Read(&role)
	return role, err
}

// UpateRole UpateRole
func UpateRole(role models.Role, o orm.Ormer) error {
	if _, err := o.Update(&role); err != nil {
		return err
	}
	return nil
}

// DeleteRole DeleteRole
func DeleteRole(id int, o orm.Ormer) error {
	if _, err := o.Delete(&models.Role{ID: id}); err != nil {
		return err
	}
	return nil
}
