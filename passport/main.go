package main

import (
    "github.com/changqing98/cqzone/user/api/web"
    "github.com/changqing98/cqzone/user/application/command/service"
    _ "github.com/changqing98/cqzone/user/docs"
    "github.com/changqing98/cqzone/user/infrastructure/data"
    "github.com/changqing98/cqzone/user/infrastructure/handler"
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
    r := gin.Default()
    engine := xorm.GetXormEngine()
    _ = engine.CreateTables(&data.UserDO{})
    repo := &persistence.UserRepositoryImpl{
        Engine: engine,
    }
    userApplicationService := &service.UserApplicationService{
        UserRepository: repo,
    }
    userController := web.UserController{
        UserApplicationService: userApplicationService,
    }
    r.Use(handler.Recover)
    r.POST("/v1/users", userController.Register)
    r.POST("/v1/users:passwordLogin", userController.PasswordLogin)
    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
    r.Run()
}
