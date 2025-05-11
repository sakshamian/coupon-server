package handler

import (
	"coupon-system/constants"
	"coupon-system/models/request"
	"coupon-system/pkg/responses"
	"coupon-system/pkg/resterrors"
	"coupon-system/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCoupon(ctx *gin.Context) {
	var req request.CreateCoupon

	if err := ctx.ShouldBindJSON(&req); err != nil {
		restError := resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
		ctx.JSON(restError.Status(), restError)
		return
	}

	err := service.CreateCoupon(&req)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, responses.ResponseCreated(nil))
}

// func UpdateAdminUser(ctx *gin.Context) {
// 	var req request.UpdateAdminUser
// 	var restError resterrors.RestErr

// 	if err := ctx.ShouldBindJSON(&req); err != nil || req.Id < 1 {
// 		restError = resterrors.NewBadRequestError(constants.MESSAGE_INVALID_INPUTS)
// 		ctx.JSON(restError.Status(), restError)
// 		return
// 	}

// 	restError = req.Validate(ctx)
// 	if restError != nil {
// 		ctx.JSON(restError.Status(), restError)
// 		return
// 	}

// 	restError = services.UpdateAdminUser(&req)
// 	if restError != nil {
// 		ctx.JSON(restError.Status(), restError)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, responses.ResponseSuccess(nil))
// }

func GetCouponList(ctx *gin.Context) {
	coupons, restError := service.GetCouponList()
	if restError != nil {
		ctx.JSON(restError.Status(), restError)
		return
	}

	couponListResponse := map[string]any{
		"coupons": coupons,
	}

	ctx.JSON(http.StatusOK, responses.ResponseSuccess(couponListResponse))
}

func DeleteCoupons(ctx *gin.Context) {
	var restError resterrors.RestErr

	couponId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || couponId < 1 {
		restError = resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
		ctx.JSON(restError.Status(), restError)
		return
	}

	restError = service.DeleteCoupon(uint16(couponId))
	if restError != nil {
		ctx.JSON(restError.Status(), restError)
		return
	}

	ctx.JSON(http.StatusOK, responses.ResponseSuccess(nil))
}
