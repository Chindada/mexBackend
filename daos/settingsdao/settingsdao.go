package settingsdao

import (
	"errors"
	"ligang/models"

	"github.com/jinzhu/gorm"
)

// CreateNewSetting CreateNewSetting
func CreateNewSetting(newsettings models.Setting, db *gorm.DB) error {
	ok := db.NewRecord(newsettings)
	if !ok {
		return errors.New("PK is not zero")
	}
	result := db.Create(&newsettings)
	return result.Error
}

// GetAllSetting GetAllSetting
func GetAllSetting(db *gorm.DB) ([]models.Setting, error) {
	var settings []models.Setting
	result := db.Find(&settings)
	return settings, result.Error
}

// UpdateSetting UpdateSetting
func UpdateSetting(editdsettings models.Setting, db *gorm.DB) error {
	result := db.Save(&editdsettings)
	return result.Error
}

// DeleteSetting DeleteSetting
func DeleteSetting(id uint, db *gorm.DB) error {
	var deletedsettings models.Setting
	deletedsettings.ID = id
	result := db.Delete(&deletedsettings)
	return result.Error
}
