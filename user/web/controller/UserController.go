package controller

import "github.com/gin-gonic/gin"

type UserController struct{}

var userController *UserController

func newUserController() *UserController {
	if userController == nil {
		userController = &UserController{}
	}
	return userController
}

// @Summary 校验手机号是否存在
// @Description description
// @Accept json
// @Produce  json
// @Param  mobile  body string true "Mobile"
// @Success 200 {bool} bool	"true已存在,false不存在"
// @Failure 404 {object} web.Error "Can not find ID"
// @Router /users:checkEmail [post]
func (userController *UserController) checkMobile(c *gin.Context) {
	c.Request.Body.Read()
}

func (userController *UserController) register(c *gin.Context) {
}
