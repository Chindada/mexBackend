package sysinit

import (
	"database/sql"
	"ligang/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// InitDB InitDB
func InitDB() {
	settings := InitDefaultSettings()
	CreateDatabase(settings)
	orm.RegisterDataBase(
		settings.DbAlias,
		settings.DbType,
		settings.DbUser+":"+settings.DbPwd+"@tcp("+settings.DbHost+":"+settings.DbPort+")/"+settings.DbName+"?charset="+settings.DbCharset+"&loc=Local",
		30)

	isDev := (settings.Runmode == "dev")
	orm.RunSyncdb("default", false, isDev)
	orm.Debug = !isDev
}

// CreateDatabase CreateDatabase
func CreateDatabase(settings models.Setting) {
	db, err := sql.Open(
		settings.DbType,
		settings.DbUser+":"+settings.DbPwd+"@tcp("+settings.DbHost+":"+settings.DbPort+")/")
	if err != nil {
		beego.Warning(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + settings.DbName + " CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
	if err != nil {
		beego.Warning(err)
	}
}
