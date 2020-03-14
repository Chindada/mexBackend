package models

import "github.com/jinzhu/gorm"

// prefix
const (
	AuthPrefix        string = "auth_"
	BasicPrefix       string = "basic_"
	DataPrefix        string = "data_"
	ManufacturePrefix string = "mes_"
)

// Setting Setting
type Setting struct {
	gorm.Model
	DbType    string `gorm:"default:'mysql'"`
	DbAlias   string `gorm:"default:'default'"`
	DbName    string `gorm:"default:'mex'"`
	DbUser    string `gorm:"default:'root'"`
	DbPwd     string `gorm:"default:'82589155'"`
	DbHost    string `gorm:"default:'123.194.35.130'"`
	DbPort    string `gorm:"default:'8881'"`
	DbCharset string `gorm:"default:'utf8mb4'"`
	Httpport  string `gorm:"default:'8088'"`
	Runmode   string `gorm:"default:'dev'"`
	LogLevel  string `gorm:"default:'7'"`
	RedisHost string `gorm:"default:'127.0.0.1:6379'"`
}
