package service

import (
	"coupon-system/models/coupon"
	"coupon-system/models/request"
	"coupon-system/models/response"
	"coupon-system/pkg/resterrors"
)

func GetApplicableCoupons(req request.ApplicableCouponRequest) ([]response.ApplicableCouponResponse, resterrors.RestErr) {
	res, err := coupon.GetApplicableCoupons(req)
	if err != nil {
		return []response.ApplicableCouponResponse{}, err
	}

	discountRes := make([]response.ApplicableCouponResponse, 0, len(res))
	for _, item := range res {
		discountRes = append(discountRes, coupon.ConvertToDiscountResponse(item))
	}

	return discountRes, nil
}
