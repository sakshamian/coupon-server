package response

type ApplicableCouponResponse struct {
	CouponCode    string  `json:"coupon_code"`
	DiscountValue float64 `json:"discount_value"`
	UsageType     string  `json:"usage_type"`
}
