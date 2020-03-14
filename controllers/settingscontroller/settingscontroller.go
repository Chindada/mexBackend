package settingscontroller

import (
	"encoding/json"
	"ligang/controllers"
	"ligang/models"
	"ligang/services/settingsservice"
	"net/http"
	"strconv"
)

// SettingsController SettingsController
type SettingsController struct {
	controllers.BaseController
}

// CreateNewSetting CreateNewSetting
func (c *SettingsController) CreateNewSetting() {
	var newsettings models.Setting
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &newsettings); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := settingsservice.CreateNewSetting(newsettings); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// GetAllSetting GetAllSetting
func (c *SettingsController) GetAllSetting() {
	if allsettings, err := settingsservice.GetAllSetting(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = allsettings
		c.ServeJSON()
	}
}

// UpdateSetting UpdateSetting
func (c *SettingsController) UpdateSetting() {
	var editdsettings models.Setting
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &editdsettings); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := settingsservice.UpdateSetting(editdsettings); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// DeleteSetting DeleteSetting
func (c *SettingsController) DeleteSetting() {
	sid := c.Ctx.Input.Header("sid")
	if sidunit, err := strconv.ParseUint(sid, 10, 32); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := settingsservice.DeleteSetting(uint(sidunit)); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}
