package userservice

import (
	"ligang/daos/resourcedao"
	"ligang/daos/roledao"
	"ligang/daos/userdao"
	"ligang/models"

	"github.com/astaxie/beego/orm"
)

// CreateUser CreateUser
func CreateUser(new models.User) error {
	o := orm.NewOrm()
	if err := userdao.CreateUser(new, o); err != nil {
		return err
	}
	return nil
}

// GetAllUser GetAllUser
func GetAllUser() ([]models.User, error) {
	o := orm.NewOrm()
	alluser, err := userdao.GetAllUser(o)
	if err != nil {
		return nil, err
	}
	return alluser, nil
}

// UpdateUser New
func UpdateUser(user models.User) error {
	o := orm.NewOrm()
	if err := userdao.UpdateUser(user, o); err != nil {
		return err
	}
	return nil
}

// DeleteUser New
func DeleteUser(id int) error {
	o := orm.NewOrm()
	if err := userdao.DeleteUser(id, o); err != nil {
		return err
	}
	return nil
}

// GetUserResourceByUsername GetUserResourceByUsername
func GetUserResourceByUsername(username string) (map[string]bool, error) {
	o := orm.NewOrm()
	if user, err := userdao.UserOneByName(username, o); err != nil {
		return nil, err
	} else if role, err := roledao.GetRoleByID(user.ID, o); err != nil {
		return nil, err
	} else if resourceMap, err := resourcedao.GetResourceByRoleID(role.ID, o); err != nil {
		return nil, err
	} else {
		return resourceMap, nil
	}
}
