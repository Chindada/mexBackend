package nowtimecontroller

import (
	"ligang/controllers"
	"time"
)

// NowTimeController NowTimeController
type NowTimeController struct {
	controllers.BaseController
}

// GetTime GetTime
func (c *NowTimeController) GetTime() {
	var res res
	res.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	c.Data["json"] = res
	c.ServeJSON()
}

type res struct {
	Timestamp int64
}
