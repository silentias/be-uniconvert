package main

import (
	"be-uniconvert/config"
	"be-uniconvert/docs"
	"be-uniconvert/internal/routers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title uniconvert API
// @version 1.0
// @host localhost:8000
// @basePath /api/v1
func main() {
	config.LoadConfig()
	r := gin.Default()
	routers.InitRouters(r)
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(config.GetHost() + ":" + config.GetPort())
}
