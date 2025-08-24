package routers

import (
	"be-uniconvert/internal/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group("/api/v1/convert")

	api.POST("audio", controllers.Audio)
}