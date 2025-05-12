package request

import (
	"coupon-system/constants"
	"coupon-system/pkg/resterrors"
	"time"
)

type ApplicableCouponRequest struct {
	CartItems  []CartItem `json:"cart_items" binding:"required"`
	OrderTotal float64    `json:"order_total"  binding:"required"`
	Timestamp  time.Time  `json:"timestamp"  binding:"required"`
}

type CartItem struct {
	Id       string `json:"id"  binding:"required"`
	Category string `json:"category"  binding:"required"`
}

func (req ApplicableCouponRequest) Validate() resterrors.RestErr {
	if len(req.CartItems) == 0 {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_CART_SIZE)
	}
	return nil
}

type ApplyCoupon struct {
	CouponCode string     `json:"coupon_code" binding:"required"`
	CartItems  []CartItem `json:"cart_items" binding:"required"`
	OrderTotal float64    `json:"order_total"  binding:"required"`
	Timestamp  time.Time  `json:"timestamp"  binding:"required"`
}

func (req ApplyCoupon) Validate() resterrors.RestErr {
	if len(req.CartItems) == 0 {
		return resterrors.NewBadRequestError(constants.MESSAGE_INVALID_CART_SIZE)
	}
	return nil
}
