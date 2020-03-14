package sysinit

import (
	"ligang/models"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	// sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Settings Settings
var Settings models.Setting

// sqlite3 file position and name
const (
	Sqlite3file string = "./conf.db"
)

// InitDefaultSettings InitDefaultSettings
func InitDefaultSettings() models.Setting {
	db, err := gorm.Open("sqlite3", Sqlite3file)
	if err != nil {
		beego.Alert(err)
	}
	defer db.Close()
	db.LogMode(false)
	db.AutoMigrate(&models.Setting{})

	var settings models.Setting

	db.Last(&settings)
	if settings.Httpport == "" {
		db.Create(&models.Setting{})
		db.Last(&settings)
	}
	Settings = settings
	return settings
}
