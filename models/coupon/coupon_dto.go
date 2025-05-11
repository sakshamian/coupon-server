package coupon

import (
	"coupon-system/models/request"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	ID                    uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CouponCode            string          `gorm:"type:varchar(50);not null;unique" json:"coupon_code"`
	UsageType             string          `gorm:"type:enum('one_time','multi_use','time_based');not null" json:"usage_type"`
	ApplicableMedicineIDs json.RawMessage `gorm:"type:json;default:null" json:"applicable_medicine_ids"`
	ApplicableCategories  json.RawMessage `gorm:"type:json;default:null" json:"applicable_categories"`
	MinOrderValue         float64         `gorm:"type:decimal(10,2);not null" json:"min_order_value"`
	ValidFrom             time.Time       `gorm:"default:null" json:"valid_from"`
	ValidTo               time.Time       `gorm:"default:null" json:"valid_to"`
	TermsAndConditions    string          `gorm:"type:text" json:"terms_and_conditions"`
	DiscountType          string          `gorm:"type:enum('inventory','charges');not null" json:"discount_type"`
	DiscountValue         float64         `gorm:"type:decimal(10,2);not null" json:"discount_value"`
	MaxUsagePerUser       int             `gorm:"type:integer" json:"max_usage_per_user"`
	IsActive              int             `gorm:"default:1" json:"is_active"`
	CreatedAt             time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt             time.Time       `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt             gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

func (Coupon) TableName() string {
	return "coupon"
}

func ConvertToCouponDTO(req *request.CreateCoupon) Coupon {
	medIDsJSON, _ := json.Marshal(req.ApplicableMedicineIDs)
	categoriesJSON, _ := json.Marshal(req.ApplicableCategories)

	return Coupon{
		CouponCode:            req.CouponCode,
		UsageType:             req.UsageType,
		ApplicableMedicineIDs: medIDsJSON,
		ApplicableCategories:  categoriesJSON,
		MinOrderValue:         req.MinOrderValue,
		ValidFrom:             req.ValidFrom,
		ValidTo:               req.ValidTo,
		TermsAndConditions:    req.TermsAndConditions,
		DiscountType:          req.DiscountType,
		MaxUsagePerUser:       req.MaxUsagePerUser,
		IsActive:              req.IsActive,
	}
}

// func ConvertToCouponResponse(req *request.CreateCoupon) Coupon {
// 	medIDsJSON, _ := json.Marshal(req.ApplicableMedicineIDs)
// 	categoriesJSON, _ := json.Marshal(req.ApplicableCategories)

// 	return Coupon{
// 		CouponCode:            req.CouponCode,
// 		UsageType:             req.UsageType,
// 		ApplicableMedicineIDs: medIDsJSON,
// 		ApplicableCategories:  categoriesJSON,
// 		MinOrderValue:         req.MinOrderValue,
// 		ValidFrom:             req.ValidFrom,
// 		ValidTo:               req.ValidTo,
// 		TermsAndConditions:    req.TermsAndConditions,
// 		DiscountType:          req.DiscountType,
// 		MaxUsagePerUser:       req.MaxUsagePerUser,
// 		IsActive:              req.IsActive,
// 	}
// }
