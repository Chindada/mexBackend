package dccontroller

import (
	"encoding/json"
	"ligang/controllers"
	"ligang/models"
	"ligang/services/dcservice"
	"net/http"
	"strconv"
)

// DcController DcController
type DcController struct {
	controllers.BaseController
}

// CreateDC CreateDC
func (c *DcController) CreateDC() {
	var new models.Dc
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &new); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := dcservice.CreateDC(new); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// GetAllDC GetAllDC
func (c *DcController) GetAllDC() {
	if allDc, err := dcservice.GetAllDC(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = allDc
		c.ServeJSON()
	}
}

// UpdateDC UpdateDC
func (c *DcController) UpdateDC() {
	dc := models.Dc{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &dc); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := dcservice.UpdateDC(dc); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}

// DeleteDC DeleteDC
func (c *DcController) DeleteDC() {
	did := c.Ctx.Input.Header("did")
	if didint, err := strconv.Atoi(did); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := dcservice.DeleteDC(didint); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
}
