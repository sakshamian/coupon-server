package response

type ApplicableCouponResponse struct {
	CouponCode    string  `json:"coupon_code"`
	DiscountValue float64 `json:"discount_value"`
	UsageType     string  `json:"usage_type"`
}

type ApplyCoupon struct {
	IsValid  bool     `json:"is_valid"`
	Discount Discount `json:"discount"`
}

type Discount struct {
	ItemDiscount    float64 `json:"item_discount"`
	ChargesDiscount float64 `json:"charges_discount"`
}
