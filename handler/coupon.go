package handler

import (
	"coupon-system/constants"
	"coupon-system/models/request"
	"coupon-system/pkg/responses"
	"coupon-system/pkg/resterrors"
	"coupon-system/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApplicableCoupons(ctx *gin.Context) {
	var req request.ApplicableCouponRequest
	var restErr resterrors.RestErr

	if err := ctx.ShouldBindJSON(req); err != nil {
		restErr = resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	if restErr = req.Validate(); restErr != nil {
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	coupons, restErr := service.GetApplicableCoupons(req)
	if restErr != nil {
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	ctx.JSON(http.StatusOK, responses.ResponseSuccess(coupons))
}
