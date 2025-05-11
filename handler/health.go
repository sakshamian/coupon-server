package handler

import (
	"coupon-system/pkg/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(ctx *gin.Context) {
	resp := map[string]any{
		"health": "OK",
	}
	ctx.JSON(http.StatusOK, responses.ResponseSuccess(resp))
}
