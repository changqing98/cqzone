package controller

import (
	"user/application/command"
	"user/application/service"

	"github.com/gin-gonic/gin"
)

// UserController 用户接扣
type UserController struct {
	userApplicationService service.UserApplicationService
}

var userController *UserController

// CreateUserController 创建用户控制器
func CreateUserController() *UserController {
	if userController == nil {
		userController = &UserController{
			userApplicationService: service.CreateUserApplicationService(),
		}
	}
	return userController
}

// SendSmsVerificationCode 发送手机验证码
// @Summary 发送手机验证码
// @Accept json
// @Produce  json
// @Param  command body command.SendSmsVerificationCodeCommand true "Mobile"
// @Success 200 {bool} bool	"true已存在,false不存在"
// @Failure 409 {object} web.Error "Can not find ID"
// @Router /users:sendVerificationCode [post]
func (userController *UserController) SendSmsVerificationCode(c *gin.Context) {
	smsVerificationCodeCommand := command.SendSmsVerificationCodeCommand{}
	c.BindJSON(&smsVerificationCodeCommand)
	verificationCode := userController.userApplicationService.SendSmsVerificationCode(smsVerificationCodeCommand)
	c.JSON(200, verificationCode)
}

// Register 用户注册
// @Summary 用户注册
// @Accept json
// @Produce  json
// @Param  mobile  body string true "Mobile"
// @Success 200 {bool} bool	"注册成功"
// @Failure 409 {object} web.Error "Can not find ID"
// @Router /users [post]
func (userController *UserController) Register(c *gin.Context) {
}

// PasswordLogin 账号密码登录
// @Summary 账号密码登录
// @Accept json
// @Produce  json
// @Param  mobile  body string true "Mobile"
// @Param password body string true "password"
// @Success 200 {bool} dto.AuthenticationDTO	"登录成功返回认证token"
// @Failure 401 {object} web.Error "密码输入错误"
// @Router /users:passwordLogin [post]
func (userController *UserController) PasswordLogin(c *gin.Context) {

}
