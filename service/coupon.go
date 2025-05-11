package service

import (
	"coupon-system/models/coupon"
	"coupon-system/models/request"
	"coupon-system/pkg/resterrors"
)

func GetApplicableCoupons(req request.ApplicableCouponRequest) ([]coupon.Coupon, resterrors.RestErr) {
	// res, err := coupon.GetApplicableCoupons(req)
	// if err != nil {
	// 	return []coupon.Coupon{}, err
	// }

	// couponRes := make(res.C)
	return coupon.GetApplicableCoupons(req)
}
