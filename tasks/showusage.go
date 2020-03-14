package tasks

import (
	"fmt"
	"ligang/utils"
	"runtime"
	"strconv"

	"github.com/astaxie/beego/toolbox"
)

func init() {
	usageTask := toolbox.NewTask("usageTask", "0 */1 * * * *", printMemUsage)
	toolbox.AddTask("usageTask", usageTask)
	toolbox.StartTask()
}

// printMemUsage printMemUsage
func printMemUsage() error {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	msg := "Alloc = " + bToMb(m.Alloc) + " MiB, TotalAlloc = " + bToMb(m.TotalAlloc) + " MiB, Sys = " + bToMb(m.Sys) + " MiB, NumGC = " + fmt.Sprint(m.NumGC)
	utils.LogInfo(msg)
	return nil
}

func bToMb(b uint64) string {
	toString := strconv.FormatUint(b/1024/1024, 10)
	return toString
}
