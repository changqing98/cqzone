package main

import (
    "github.com/changqing98/cqzone/user/api/web/controller"
    "github.com/changqing98/cqzone/user/application/command/service"
    _ "github.com/changqing98/cqzone/user/docs"
    "github.com/changqing98/cqzone/user/infrastructure/persistence"
    "github.com/changqing98/cqzone/user/infrastructure/xorm"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func main() {
    engine := xorm.GetXormEngine()
    userRepo := persistence.UserRepositoryImpl{Engine: engine}
    userApplicationService := service.NewUserApplicationService(userRepo)
    userController := controller.NewUserController(userApplicationService)

    r := gin.New()

    r.POST("/v1/users", userController.Register)

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
    r.Run()
}
