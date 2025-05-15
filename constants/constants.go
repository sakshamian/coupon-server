package constants

import "time"

const (
	// error messages
	MESSAGE_INVALID_REQUEST_PARAMETERS = "Invalid Reqeust parameters"
	MESSAGE_INVALID_CART_SIZE          = "Invalid cart size"
	MESSAGE_SOMETHING_WENT_WRONG       = "Something went wrong"
	MESSAGE_NO_COUPON_FOUND            = "Coupon not found"
	MESSAGE_NO_APPLICABLE_COUPON       = "No applicable coupons"
	MESSAGE_COUPON_NOT_APPLIED         = "Coupon cannot be applied to this cart"
	MESSAGE_COUPON_ALREADY_EXISTS      = "A coupon exists with the same code"

	// discount type
	DISCOUNT_TYPE_CHARGE = "charges"
	DISCOUNT_TYPE_ITEM   = "inventory"

	// usage type
	USAGE_TYPE_SINGLE     = "one_time"
	USAGE_TYPE_MULTIPLE   = "multi_use"
	USAGE_TYPE_TIME_BASED = "time_based"

	// active flag
	ACTIVE_COUPON_TYPE   = 1
	INACTIVE_COUPON_TYPE = 0

	// cache constants
	COUPON_CACHE_KEY = "coupon_%s"
	COUPON_CACHE_TTL = time.Hour * 24

	// database
	MAX_RETRIES = 10
)
