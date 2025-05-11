package response

type ApplicableCouponResponse struct {
	ApplicableCoupons []Coupons `json:"applicable_coupons"`
}

type Coupons struct {
	CouponCode    string  `json:"coupon_code"`
	DiscountValue float64 `json:"discount_value"`
}
