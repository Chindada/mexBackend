package rolecontroller

import (
	"encoding/json"
	"ligang/controllers"
	"ligang/models"
	"ligang/services/roleservice"
	"net/http"
	"strconv"
)

// RoleController RoleController
type RoleController struct {
	controllers.BaseController
}

// CreateNewRole CreateNewRole
func (c *RoleController) CreateNewRole() {
	new := models.Role{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &new); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := roleservice.CreateNewRole(new); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// GetAllRole GetAllRole
func (c *RoleController) GetAllRole() {
	if allrole, err := roleservice.GetAllRole(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = allrole
		c.ServeJSON()
	}
}

// GetOneRole GetOneRole
func (c *RoleController) GetOneRole() {

}

// UpateRole UpateRole
func (c *RoleController) UpateRole() {
	role := models.Role{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &role); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := roleservice.UpateRole(role); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// DeleteRole DeleteRole
func (c *RoleController) DeleteRole() {
	rid := c.Ctx.Input.Header("rid")
	if ridint, err := strconv.Atoi(rid); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := roleservice.DeleteRole(ridint); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// AddRoleResourceRel AddRoleResourceRel
func (c *RoleController) AddRoleResourceRel() {
	add := models.Role{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &add); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := roleservice.AddRoleResourceRel(add); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}
