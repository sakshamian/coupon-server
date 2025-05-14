package service

import (
	"coupon-system/models/coupon"
	"coupon-system/models/request"
	"coupon-system/pkg/resterrors"
)

func CreateCoupon(req *request.CreateCoupon) resterrors.RestErr {
	coupon := coupon.ConvertToCouponDTO(req)
	return coupon.Create()
}

func GetCouponList() ([]coupon.Coupon, resterrors.RestErr) {
	return coupon.GetCouponList()
}

func DeleteCoupon(couponId uint16) resterrors.RestErr {
	return coupon.Delete(uint(couponId))
}
