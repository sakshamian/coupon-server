package service

import (
	"coupon-system/constants"
	"coupon-system/models/coupon"
	"coupon-system/models/request"
	"coupon-system/models/response"
	"coupon-system/pkg/resterrors"
	cache "coupon-system/redis"
	"encoding/json"
	"fmt"
	"slices"
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

func ApplyCoupon(req request.ApplyCoupon) (response.ApplyCoupon, resterrors.RestErr) {
	couponData, err := getCouponData(req.CouponCode)
	if err != nil {
		return response.ApplyCoupon{}, err
	}

	if validateCoupon(req, couponData) {
		return coupon.ConvertToApplyCouponResponse(couponData), nil
	}

	return response.ApplyCoupon{}, resterrors.NewBadRequestError(constants.MESSAGE_COUPON_NOT_APPLIED)
}

func getCouponData(couponCode string) (coupon.Coupon, resterrors.RestErr) {
	cacheKey := fmt.Sprintf(constants.COUPON_CACHE_KEY, couponCode)
	data, err := cache.GetFromRedis(cacheKey)
	if err == nil {
		var couponData coupon.Coupon
		if err := json.Unmarshal([]byte(*data), &couponData); err == nil {
			return couponData, nil
		}
	}

	couponData, restErr := coupon.GetCouponByCouponCode(couponCode)
	if restErr != nil {
		return coupon.Coupon{}, restErr
	}

	go cache.SetInRedis(&cache.SetRedis{
		Key:  cacheKey,
		Data: couponData,
		Exp:  constants.COUPON_CACHE_TTL,
	})
	return couponData, nil
}

func validateCoupon(req request.ApplyCoupon, couponData coupon.Coupon) bool {
	if couponData.ValidFrom.Before(req.Timestamp) && couponData.ValidTo.After(req.Timestamp) && couponData.IsActive == 1 && couponData.MinOrderValue <= req.OrderTotal {
		var medIDs, catIDs []string
		_ = json.Unmarshal(couponData.ApplicableMedicineIDs, &medIDs)
		_ = json.Unmarshal(couponData.ApplicableCategories, &catIDs)

		counter := 0
		for _, couponData := range req.CartItems {
			if slices.Contains(medIDs, couponData.Id) && slices.Contains(catIDs, couponData.Category) {
				counter++
			}
		}
		if counter > 0 {
			return true
		}
	}
	return false
}
