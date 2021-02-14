package web

import (
    "github.com/changqing98/cqzone/user/application/command/protocol"
    "github.com/changqing98/cqzone/user/application/command/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

// UserController 用户接口
type UserController struct {
    UserApplicationService *service.UserApplicationService
}

// Register 用户注册
// @Summary 用户注册
// @Accept json
// @Produce  json
// @Param CreateUserCommand body protocol.CreateUserCommand true "用户注册请求"
// @Success 200 {object} protocol.AccessTokenResponse "注册成功"
// @Failure 409 {object} gin.Error "Can not find ID"
// @Router /users [post]
func (userController *UserController) Register(c *gin.Context) {
    var command protocol.CreateUserCommand
    _ = c.ShouldBind(&command)
    response := userController.UserApplicationService.Register(&command)
    c.JSON(http.StatusOK, response)
}

// PasswordLogin 账号密码登录
// @Summary 账号密码登录
// @Accept json
// @Produce  json
// @Param PasswordLoginCommand body protocol.PasswordLoginCommand true "用户注册请求"
// @Success 200 {object} protocol.AccessTokenResponse	"登录成功返回认证token"
// @Failure 401 {object} gin.Error "密码输入错误"
// @Router /users:passwordLogin [post]
func (userController *UserController) PasswordLogin(c *gin.Context) {
    var command protocol.PasswordLoginCommand
    _ = c.ShouldBind(&command)
    response := userController.UserApplicationService.PasswordLogin(&command)
    c.JSON(http.StatusOK, response)
}
