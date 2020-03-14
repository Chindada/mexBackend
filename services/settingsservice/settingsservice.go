package settingsservice

import (
	"errors"
	"ligang/daos/settingsdao"
	"ligang/models"
	"ligang/sysinit"

	"github.com/jinzhu/gorm"
)

// CreateNewSetting CreateNewSetting
func CreateNewSetting(newsetting models.Setting) error {
	db, err := gorm.Open("sqlite3", sysinit.Sqlite3file)
	if err != nil {
		return errors.New("sqlite3 is panic")
	}
	defer db.Close()
	db.LogMode(false)
	if err := settingsdao.CreateNewSetting(newsetting, db); err != nil {
		return err
	}
	return nil
}

// GetAllSetting GetAllSetting
func GetAllSetting() ([]models.Setting, error) {
	db, err := gorm.Open("sqlite3", sysinit.Sqlite3file)
	if err != nil {
		return nil, errors.New("sqlite3 is panic")
	}
	defer db.Close()
	db.LogMode(false)
	allsettings, err := settingsdao.GetAllSetting(db)
	if err != nil {
		return nil, err
	}
	return allsettings, nil
}

// UpdateSetting UpdateSetting
func UpdateSetting(editdsettings models.Setting) error {
	db, err := gorm.Open("sqlite3", sysinit.Sqlite3file)
	if err != nil {
		return errors.New("sqlite3 is panic")
	}
	defer db.Close()
	db.LogMode(false)
	if err := settingsdao.UpdateSetting(editdsettings, db); err != nil {
		return err
	}
	return nil
}

// DeleteSetting DeleteSetting
func DeleteSetting(id uint) error {
	db, err := gorm.Open("sqlite3", sysinit.Sqlite3file)
	if err != nil {
		return errors.New("sqlite3 is panic")
	}
	defer db.Close()
	db.LogMode(false)
	if err := settingsdao.DeleteSetting(id, db); err != nil {
		return err
	}
	return nil
}
