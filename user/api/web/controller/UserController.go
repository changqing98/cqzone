package controller

import (
    "github.com/changqing98/cqzone/user/application/command"
    "github.com/changqing98/cqzone/user/application/command/service"
    "github.com/gin-gonic/gin"
)

// UserController 用户接口
type UserController struct {
    userApplicationService *service.UserApplicationService
}

// Register 用户注册
// @Summary 用户注册
// @Accept json
// @Produce  json
// @Param  command body command.CreateUserCommand true "用户注册请求"
// @Success 200 {object} dto.AuthenticationDTO "注册成功"
// @Failure 409 {object} web.Error "Can not find ID"
// @Router /users [post]
func (userController UserController) Register(c *gin.Context) {
    userApplicationService := userController.userApplicationService
    var userRegisterCommand *command.CreateUserCommand
    _ = c.ShouldBind(userRegisterCommand)
    result := userApplicationService.Register(userRegisterCommand)
    c.JSON(200, result)
}

// PasswordLogin 账号密码登录
// @Summary 账号密码登录
// @Accept json
// @Produce  json
// @Param  command body command.PasswordLoginCommand true "用户注册请求"
// @Success 200 {object} dto.AuthenticationDTO	"登录成功返回认证token"
// @Failure 401 {object} web.Error "密码输入错误"
// @Router /users:passwordLogin [post]
func (userController UserController) PasswordLogin(c *gin.Context) {
    userApplicationService := userController.userApplicationService
    var passwordLoginCommand *command.PasswordLoginCommand
    _ = c.ShouldBind(passwordLoginCommand)
    result := userApplicationService.PasswordLogin(passwordLoginCommand)
    c.JSON(200, result)
}
