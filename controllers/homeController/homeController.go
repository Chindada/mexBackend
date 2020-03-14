package homecontroller

import (
	"ligang/controllers"
	"ligang/functions/restartserver"
	"ligang/services/homeservice"
	"ligang/services/userservice"
	"ligang/utils"
	"net/http"

	"github.com/astaxie/beego"
)

// Homecontroller Homecontroller
type Homecontroller struct {
	controllers.BaseController
}

// CookieJWTage CookieJWTage
const CookieJWTage int64 = 3600

// Login Login
func (c *Homecontroller) Login() {
	username := c.Ctx.Input.Header("username")
	password := c.Ctx.Input.Header("password")
	var res loginRes
	isSuper, err := homeservice.CheckUserPass(username, password)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	token := utils.GenToken(username, CookieJWTage)
	c.Ctx.SetCookie("jwt", token, CookieJWTage)
	beego.Informational("Checking User Resource")
	if isSuper {
		supermap := make(map[string]bool)
		supermap["isSuper"] = true
		c.SetSession("usersession", supermap)
		res = supermap
	} else if resource, err := userservice.GetUserResourceByUsername(username); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		res = resource
		if resource[c.ActionName] {
			c.SetSession("usersession", resource)
		}
	}
	c.Data["json"] = res
	c.ServeJSON()
}

// Restart Restart
func (c *Homecontroller) Restart() {
	restartserver.Restart()
}

type loginRes map[string]bool
