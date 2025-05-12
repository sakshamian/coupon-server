package coupon

import (
	"coupon-system/constants"
	"coupon-system/db"
	"encoding/json"
	"slices"

	"coupon-system/models/request"
	"coupon-system/pkg/resterrors"
)

func GetApplicableCoupons(req request.ApplicableCouponRequest) ([]Coupon, resterrors.RestErr) {
	var coupons []Coupon
	result := db.DB.Debug().Table(Coupon{}.TableName()).
		Where("valid_from <= ? AND valid_to >= ?", req.Timestamp, req.Timestamp).
		Where("min_order_value <= ?", req.OrderTotal).
		Where("is_active = ?", 1).
		Find(&coupons)
	if result.Error != nil {
		return []Coupon{}, resterrors.NewInternalServerError(constants.MESSAGE_SOMETHING_WENT_WRONG, result.Error)
	}
	if result.RowsAffected == 0 {
		return []Coupon{}, resterrors.NewBadRequestError(constants.MESSAGE_NO_APPLICABLE_COUPON)
	}

	var couponRes []Coupon
	for _, item := range coupons {
		var medIDs, catIDs []string
		_ = json.Unmarshal(item.ApplicableMedicineIDs, &medIDs)
		_ = json.Unmarshal(item.ApplicableCategories, &catIDs)

		for _, c := range req.CartItems {
			if slices.Contains(medIDs, c.Id) && slices.Contains(catIDs, c.Category) {
				couponRes = append(couponRes, item)
				break
			}
		}
	}
	return couponRes, nil
}

func ApplyCoupon(req request.ApplyCoupon) (Coupon, resterrors.RestErr) {
	var coupon Coupon
	result := db.DB.Debug().Table(Coupon{}.TableName()).
		Where("coupon_code = ?", req.CouponCode).
		Where("valid_from <= ? AND valid_to >= ?", req.Timestamp, req.Timestamp).
		Where("min_order_value <= ?", req.OrderTotal).
		Where("is_active = ?", 1).
		First(&coupon)
	if result.Error != nil {
		return Coupon{}, resterrors.NewInternalServerError(constants.MESSAGE_SOMETHING_WENT_WRONG, result.Error)
	}
	if result.RowsAffected == 0 {
		return Coupon{}, resterrors.NewBadRequestError(constants.MESSAGE_NO_APPLICABLE_COUPON)
	}

	var medIDs, catIDs []string
	_ = json.Unmarshal(coupon.ApplicableMedicineIDs, &medIDs)
	_ = json.Unmarshal(coupon.ApplicableCategories, &catIDs)

	counter := 0
	for _, c := range req.CartItems {
		if slices.Contains(medIDs, c.Id) && slices.Contains(catIDs, c.Category) {
			counter++
		}
	}
	if counter == 0 {
		return Coupon{}, resterrors.NewBadRequestError(constants.MESSAGE_COUPON_NOT_APPLIED)
	}

	return coupon, nil
}

// admin APIs
func (c Coupon) Create() resterrors.RestErr {
	result := db.DB.Table(Coupon{}.TableName()).Create(&c)
	if result.Error != nil {
		return resterrors.NewInternalServerError(constants.MESSAGE_SOMETHING_WENT_WRONG, result.Error)
	}
	return nil
}

func Delete(couponId uint) resterrors.RestErr {
	result := db.DB.Table(Coupon{}.TableName()).Update("is_active = ?", 0).Where("id = ?", couponId)
	if result.RowsAffected == 0 {
		resterrors.NewNotFoundError(constants.MESSAGE_NO_COUPON_FOUND)
	}
	if result.Error != nil {
		return resterrors.NewInternalServerError(constants.MESSAGE_SOMETHING_WENT_WRONG, result.Error)
	}
	return nil
}

func GetCouponList() ([]Coupon, resterrors.RestErr) {
	var coupon []Coupon
	result := db.DB.Table(Coupon{}.TableName()).Find(&coupon)
	if result.Error != nil {
		return []Coupon{}, resterrors.NewInternalServerError(constants.MESSAGE_SOMETHING_WENT_WRONG, result.Error)
	}
	return coupon, nil
}
