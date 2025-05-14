package coupon

import (
	"coupon-system/constants"
	"coupon-system/db"

	"coupon-system/models/request"
	"coupon-system/pkg/resterrors"

	"github.com/go-sql-driver/mysql"
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

	return coupons, nil
}

func GetCouponByCouponCode(couponCode string) (Coupon, resterrors.RestErr) {
	var coupon Coupon
	result := db.DB.Debug().Table(Coupon{}.TableName()).Where("coupon_code = ?", couponCode).First(&coupon)
	if result.Error != nil {
		return Coupon{}, resterrors.NewInternalServerError(constants.MESSAGE_SOMETHING_WENT_WRONG, result.Error)
	}
	if result.RowsAffected == 0 {
		return Coupon{}, resterrors.NewBadRequestError(constants.MESSAGE_NO_APPLICABLE_COUPON)
	}

	return coupon, nil
}

// admin APIs
func (c Coupon) Create() resterrors.RestErr {
	result := db.DB.Table(Coupon{}.TableName()).Create(&c)
	if mysqlError, ok := result.Error.(*mysql.MySQLError); ok && mysqlError.Number == 1062 {
		return resterrors.NewConflictError(constants.MESSAGE_COUPON_ALREADY_EXISTS)
	}
	if result.Error != nil {
		return resterrors.NewInternalServerError(constants.MESSAGE_SOMETHING_WENT_WRONG, result.Error)
	}
	return nil
}

func Delete(couponId uint) resterrors.RestErr {
	result := db.DB.Table(Coupon{}.TableName()).Where("id = ?", couponId).Update("is_active", 0)
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
