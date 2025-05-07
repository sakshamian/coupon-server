package routes

import (
	"coupon-system/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	couponApi := router.Group("api")
	{
		couponApi.GET("health", handler.GetHealth)
	}
}
