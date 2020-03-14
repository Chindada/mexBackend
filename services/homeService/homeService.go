package homeservice

import (
	"errors"
	"ligang/daos/userdao"
	"ligang/utils"

	"github.com/astaxie/beego/orm"
)

// CheckUserPass CheckUserPass
func CheckUserPass(username, password string) (isSuper bool, err error) {
	o := orm.NewOrm()
	md5papsswd := utils.String2md5(password)
	k, err := userdao.UserOneByName(username, o)
	if err != nil {
		return false, err
	}
	if md5papsswd != k.Password {
		return false, errors.New("Password Is Wrong")
	} else if !k.IsActive {
		return false, errors.New("User Is Not Active")
	}
	return k.IsSuper, nil
}
