package web

import "github.com/gin-gonic/gin"

// Authorize 授权，redirect to /{redirectUri}?code={code}&state={state}
// @Summary
// @Accept json
// @Produce  json
// @Param AuthorizeRequest body protocol.AuthorizeRequest true "授权请求"
// @Success 302
// @Router /oauth2.0/authorize [get]
func (userController UserController) Authorize(c *gin.Context) {}

// FetchToken 获取access token
// @Summary
// @Accept json
// @Produce  json
// @Param AccessTokenRequest body protocol.AccessTokenRequest true "获取access token"
// @Success 200 {object} protocol.AccessTokenResponse "获取token成功"
// @Router /oauth2.0/token [post]
func (userController UserController) FetchToken(c *gin.Context) {}
