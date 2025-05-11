package request

import (
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
	IsActive              int       `json:"is_active,omitempty"`
}
