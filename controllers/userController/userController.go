package usercontroller

import (
	"encoding/json"
	"ligang/controllers"
	"ligang/models"
	"ligang/services/userservice"
	"net/http"
	"strconv"
)

// UserController UserController
type UserController struct {
	controllers.BaseController
}

// NewUser NewUser
func (c *UserController) NewUser() {
	var new models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &new); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := userservice.CreateUser(new); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
}

// GetAllUser GetAllUser
func (c *UserController) GetAllUser() {
	if alluser, err := userservice.GetAllUser(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = alluser
		c.ServeJSON()
	}
}

// UpdateUser UpdateUser
func (c *UserController) UpdateUser() {
	user := models.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := userservice.UpdateUser(user); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
}

// DeleteUser DeleteUser
func (c *UserController) DeleteUser() {
	uid := c.Ctx.Input.Header("uid")
	if uidint, err := strconv.Atoi(uid); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else if err := userservice.DeleteUser(uidint); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
}

// GetUserResourceByUsername GetUserResourceByUsername
func (c *UserController) GetUserResourceByUsername() {
	username := c.Ctx.Input.Header("username")
	if resource, err := userservice.GetUserResourceByUsername(username); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = resource
		c.ServeJSON()
	}
}
