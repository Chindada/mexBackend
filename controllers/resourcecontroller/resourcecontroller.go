package resourcecontroller

import (
	"encoding/json"
	"ligang/controllers"
	"ligang/models"
	"ligang/services/resourceservice"
	"net/http"
	"strconv"
)

// ResourceController 權限
type ResourceController struct {
	controllers.BaseController
}

// CreateResource CreateResource
func (c *ResourceController) CreateResource() {
	var new models.Resource
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &new); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := resourceservice.CreateResource(new); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// GetAllResource GetAllResource
func (c *ResourceController) GetAllResource() {
	if allresource, err := resourceservice.GetAllResource(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = allresource
		c.ServeJSON()
	}
}

// UpdateResource UpdateResource
func (c *ResourceController) UpdateResource() {
	resource := models.Resource{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &resource); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := resourceservice.UpdateResource(resource); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// DeleteResource DeleteResource
func (c *ResourceController) DeleteResource() {
	rid := c.Ctx.Input.Header("rid")
	if ridint, err := strconv.Atoi(rid); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := resourceservice.DeleteResource(ridint); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}
