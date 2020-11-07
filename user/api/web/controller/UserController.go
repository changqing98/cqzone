package controller

import (
    "github.com/changqing98/cqzone/user/application/command"
    appSvc "github.com/changqing98/cqzone/user/application/service"
    "github.com/gin-gonic/gin"
)

// UserController 用户接扣
type UserController struct {
    userApplicationService appSvc.UserApplicationService
}

func NewUserController(userApplicationService appSvc.UserApplicationService) UserController {
    return UserController{userApplicationService}
}

// SendSmsVerificationCode 发送手机验证码
// @Summary 发送手机验证码
// @Accept json
// @Produce  json
// @Param  command body command.SendSmsVerificationCodeCommand true "Mobile"
// @Success 200 {bool} bool	"true已存在,false不存在"
// @Failure 409 {object} web.Error "Can not find ID"
// @Router /users:sendVerificationCode [post]
func (userController UserController) SendSmsVerificationCode(c *gin.Context) {
    smsVerificationCodeCommand := command.SendSmsVerificationCodeCommand{}
    c.BindJSON(&smsVerificationCodeCommand)
    verificationCode := userController.userApplicationService.SendSmsVerificationCode(smsVerificationCodeCommand)
    c.JSON(200, verificationCode)
}

// Register 用户注册
// @Summary 用户注册
// @Accept json
// @Produce  json
// @Param  command body command.EmailRegisterCommand true "用户注册请求"
// @Success 200 {object} dto.AuthenticationDTO "注册成功"
// @Failure 409 {object} web.Error "Can not find ID"
// @Router /users [post]
func (userController UserController) Register(c *gin.Context) {
    userApplicationService := userController.userApplicationService
    var userRegisterCommand command.EmailRegisterCommand
    c.ShouldBind(&userRegisterCommand)
    result := userApplicationService.Register(userRegisterCommand)
    c.JSON(200, result)
}

// PasswordLogin 账号密码登录
// @Summary 账号密码登录
// @Accept json
// @Produce  json
// @Param  mobile  body string true "Mobile"
// @Param password body string true "password"
// @Success 200 {object} dto.AuthenticationDTO	"登录成功返回认证token"
// @Failure 401 {object} web.Error "密码输入错误"
// @Router /users:passwordLogin [post]
func (userController UserController) PasswordLogin(c *gin.Context) {

}
