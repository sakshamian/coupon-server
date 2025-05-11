package routes

import (
	"coupon-system/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("health", handler.GetHealth)

			admin := v1.Group("admin")
			{
				// TODO: middleware for admin auth
				couponManagement := admin.Group("coupon")
				{
					couponManagement.POST("", handler.CreateCoupon)
					couponManagement.GET("", handler.GetCouponList)
					couponManagement.DELETE("/:id", handler.DeleteCoupons)
				}
			}

			couponModule := v1.Group("coupons")
			{
				couponModule.GET("applicable", handler.GetApplicableCoupons)
			}
		}
	}
}
