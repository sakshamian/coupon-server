package request

import (
	"coupon-system/constants"
	"coupon-system/pkg/resterrors"
	"time"
)

type CreateCoupon struct {
	CouponCode            string    `json:"coupon_code" binding:"required"`
	UsageType             string    `json:"usage_type" binding:"required"`
	ApplicableMedicineIDs []string  `json:"applicable_medicine_ids,omitempty"`
	ApplicableCategories  []string  `json:"applicable_categories,omitempty"`
	MinOrderValue         float64   `json:"min_order_value" binding:"required"`
	ValidFrom             time.Time `json:"valid_from" binding:"required"`
	ValidTo               time.Time `json:"valid_to" binding:"required"`
	TermsAndConditions    string    `json:"terms_and_conditions,omitempty"`
	DiscountType          string    `json:"discount_type" binding:"required"`
	DiscountValue         float64   `json:"discount_value" binding:"required"`
	MaxUsagePerUser       int       `json:"max_usage_per_user" binding:"required"`
	IsActive              int       `json:"is_active" binding:"required"`
}

func (req CreateCoupon) Validate() resterrors.RestErr {
	if req.UsageType != constants.USAGE_TYPE_SINGLE && req.UsageType != constants.USAGE_TYPE_MULTIPLE && req.UsageType != constants.USAGE_TYPE_TIME_BASED {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
	}
	if len(req.ApplicableMedicineIDs) == 0 || len(req.ApplicableCategories) == 0 {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
	}
	if req.MinOrderValue <= 0 || req.MaxUsagePerUser <= 0 || req.DiscountValue <= 0 {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
	}
	if req.ValidTo.Before(req.ValidFrom) {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
	}
	if req.DiscountType != constants.DISCOUNT_TYPE_CHARGE && req.DiscountType != constants.DISCOUNT_TYPE_ITEM {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
	}
	if req.DiscountType != constants.DISCOUNT_TYPE_CHARGE && req.DiscountType != constants.DISCOUNT_TYPE_ITEM {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
	}
	if req.IsActive != constants.ACTIVE_COUPON_TYPE && req.IsActive != constants.INACTIVE_COUPON_TYPE {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_REQUEST_PARAMETERS)
	}
	return nil
}
