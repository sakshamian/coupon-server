package coupon

import (
	"coupon-system/constants"
	"coupon-system/db"

	"coupon-system/models/request"
	"coupon-system/pkg/resterrors"
)

func GetApplicableCoupons(req request.ApplicableCouponRequest) ([]Coupon, resterrors.RestErr) {
	var coupons []Coupon
	result := db.DB.Table(Coupon{}.TableName()).
		Where("is_active = ?", 1).
		Where("min_order_value <= ?", req.OrderTotal).
		// Where("(valid_from IS NULL OR valid_from <= ?)", currentTime).
		// Where("(valid_to IS NULL OR valid_to >= ?)", currentTime).
		Find(&coupons)
	if result.Error != nil {
		return []Coupon{}, resterrors.NewInternalServerError(constants.MESSAGE_SOMETHING_WENT_WRONG, result.Error)
	}
	if result.RowsAffected == 0 {
		return []Coupon{}, resterrors.NewBadRequestError(constants.MESSAGE_NO_APPLICABLE_COUPON)
	}
	return coupons, nil
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
