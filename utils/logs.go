package utils

import (
	"fmt"
	"ligang/sysinit"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var consoleLogs *logs.BeeLogger
var fileLogs *logs.BeeLogger
var runmode string

// InitLogs InitLogs
func InitLogs() {
	fmt.Println("Init Logs")
	settings := sysinit.InitDefaultSettings()
	level := settings.LogLevel
	runmode = settings.Runmode
	if runmode == "" {
		runmode = "dev"
	}
	beego.Informational("Logs Level is", level)

	consoleLogs = logs.NewLogger(128)
	consoleLogs.SetLogger(logs.AdapterConsole, `{
		"level":`+level+`,
		"color":true
		}`)
	consoleLogs.Async()

	fileLogs = logs.NewLogger(1024)
	fileLogs.SetLogger(logs.AdapterFile, `{
		"filename":"Server_Log.log",
		"level":`+level+`,
		"daily":true,
		"maxdays":10
		}`)
	fileLogs.Async()
}

// LogEmergency LogEmergency
func LogEmergency(v interface{}) {
	log("emergency", v)
}

// LogAlert LogAlert
func LogAlert(v interface{}) {
	log("alert", v)
}

// LogCritical LogCritical
func LogCritical(v interface{}) {
	log("critical", v)
}

// LogError LogError
func LogError(v interface{}) {
	log("error", v)
}

// LogWarning LogWarning
func LogWarning(v interface{}) {
	log("warning", v)
}

// LogNotice LogNotice
func LogNotice(v interface{}) {
	log("notice", v)
}

// LogInfo LogInfo
func LogInfo(v interface{}) {
	log("info", v)
}

// LogDebug LogDebug
func LogDebug(v interface{}) {
	log("debug", v)
}

// LogTrace LogTrace
func LogTrace(v interface{}) {
	log("trace", v)
}

// Log Log
func log(level, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	if runmode == "dev" {
		switch level {
		case "emergency":
			fileLogs.Emergency(format, v)
		case "alert":
			fileLogs.Alert(format, v)
		case "critical":
			fileLogs.Critical(format, v)
		case "error":
			fileLogs.Error(format, v)
		case "warning":
			fileLogs.Warning(format, v)
		case "notice":
			fileLogs.Notice(format, v)
		case "info":
			fileLogs.Info(format, v)
		case "debug":
			fileLogs.Debug(format, v)
		case "trace":
			fileLogs.Trace(format, v)
		default:
			fileLogs.Debug(format, v)
		}
	}
	switch level {
	case "emergency":
		consoleLogs.Emergency(format, v)
	case "alert":
		consoleLogs.Alert(format, v)
	case "critical":
		consoleLogs.Critical(format, v)
	case "error":
		consoleLogs.Error(format, v)
	case "warning":
		consoleLogs.Warning(format, v)
	case "notice":
		consoleLogs.Notice(format, v)
	case "info":
		consoleLogs.Info(format, v)
	case "debug":
		consoleLogs.Debug(format, v)
	case "trace":
		consoleLogs.Trace(format, v)
	default:
		consoleLogs.Debug(format, v)
	}
}
