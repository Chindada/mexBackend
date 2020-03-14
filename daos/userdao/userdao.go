package userdao

import (
	"errors"
	"ligang/models"
	"ligang/utils"

	"github.com/astaxie/beego/orm"
)

// CreateUser New
func CreateUser(new models.User, o orm.Ormer) error {
	new.ID = 0
	if new.Password == "" {
		return errors.New("Password is empty")
	}
	new.Password = utils.String2md5(new.Password)
	_, err := o.Insert(&new)
	return err
}

// UserOneByName UserOneByName
func UserOneByName(username string, o orm.Ormer) (models.User, error) {
	m := models.User{Username: username}
	err := o.Read(&m, "Username")
	return m, err
}

// GetAllUser GetAllUser
func GetAllUser(o orm.Ormer) ([]models.User, error) {
	var alluser []models.User
	_, err := o.QueryTable(models.UserTBName).All(&alluser)
	return alluser, err
}

// UpdateUser New
func UpdateUser(user models.User, o orm.Ormer) error {
	user.Password = utils.String2md5(user.Password)
	if _, err := o.Update(&user); err != nil {
		return err
	}
	return nil
}

// DeleteUser New
func DeleteUser(id int, o orm.Ormer) error {
	if _, err := o.Delete(&models.User{ID: id}); err != nil {
		return err
	}
	return nil
}
