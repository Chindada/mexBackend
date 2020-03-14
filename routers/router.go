package routers

import (
	"github.com/astaxie/beego/context"
	"ligang/controllers/dccontroller"
	"ligang/controllers/homecontroller"
	"ligang/controllers/nowtimecontroller"
	"ligang/controllers/resourcecontroller"
	"ligang/controllers/rolecontroller"
	"ligang/controllers/settingscontroller"
	"ligang/controllers/usercontroller"
	"ligang/models"
	"ligang/services/resourceservice"
	"ligang/utils"
	"strings"

	"github.com/astaxie/beego"
)

// CookieJWTage CookieJWTage
const CookieJWTage int64 = 3600

const (
	userMethod         = "post:NewUser;get:GetAllUser;put:UpdateUser;delete:DeleteUser"
	loginMethod        = "get:Login"
	restartMethod      = "get:Restart"
	userresourceMethod = "get:GetUserResourceByUsername"
	settingsMethod     = "post:CreateNewSetting;get:GetAllSetting;put:UpdateSetting;delete:DeleteSetting"
	nowtimeMethod      = "get:GetTime"
	roleMethod         = "post:CreateNewRole;get:GetAllRole;put:UpateRole;delete:DeleteRole"
	roleResourceMethod = "post:AddRoleResourceRel"
	resourceMethod     = "post:CreateResource;get:GetAllResource;put:UpdateResource;delete:DeleteResource"
	dcMethod           = "post:CreateDC;get:GetAllDC;put:UpdateDC;delete:DeleteDC"
)

func init() {
	// Filter for jwt
	beego.InsertFilter("/api/*", beego.BeforeRouter, func(ctx *context.Context) {
		cookie, err := ctx.Request.Cookie("jwt")
		if err != nil {
			utils.LogError(err)
		} else if username, ok := utils.CheckToken(cookie.Value); ok {
			token := utils.GenToken(username, CookieJWTage)
			ctx.SetCookie("jwt", token, CookieJWTage)
		} else {
			ctx.Redirect(500, "/")
			beego.Informational("Not OK")
		}
	})

	beego.Router("/auth/login",
		&homecontroller.Homecontroller{},
		loginMethod)

	beego.Router("/system/user",
		&usercontroller.UserController{},
		userMethod)

	beego.Router("/system/restart",
		&homecontroller.Homecontroller{},
		restartMethod)

	beego.Router("/api/user/getuserresource",
		&usercontroller.UserController{},
		userresourceMethod)

	beego.Router("/api/settings",
		&settingscontroller.SettingsController{},
		settingsMethod)

	beego.Router("/api/nowtime",
		&nowtimecontroller.NowTimeController{},
		nowtimeMethod)

	beego.Router("/api/role",
		&rolecontroller.RoleController{},
		roleMethod)

	beego.Router("/api/addroleresourcerel",
		&rolecontroller.RoleController{},
		roleResourceMethod)

	beego.Router("/api/resource",
		&resourcecontroller.ResourceController{},
		resourceMethod)

	beego.Router("/api/dc",
		&dccontroller.DcController{},
		dcMethod)
}

// SplitMethod SplitMethod
func SplitMethod() {
	var result []string
	allMethod := strings.Split(
		userMethod+";"+
			loginMethod+";"+
			restartMethod+";"+
			userresourceMethod+";"+
			settingsMethod+";"+
			nowtimeMethod+";"+
			roleMethod+";"+
			roleResourceMethod+";"+
			resourceMethod+";"+
			dcMethod, ";")
	for _, k := range allMethod {
		step2 := strings.Split(k, ":")
		result = append(result, step2...)
	}
	allResourceMap, err := resourceservice.StoreAllResourceMap()
	if err != nil {
		return
	}
	for _, v := range result {
		if v == "get" || v == "post" || v == "delete" || v == "patch" || v == "put" {
			continue
		}
		if _, ok := allResourceMap.Load(v); !ok {
			resource := models.Resource{}
			resource.Title = v
			if err := resourceservice.CreateResource(resource); err != nil {
				beego.Informational(err)
			}
		}
	}
}
