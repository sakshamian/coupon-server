package service

import (
	"coupon-system/models/coupon"
	"coupon-system/models/request"
	"coupon-system/pkg/resterrors"
)

func CreateCoupon(req *request.CreateCoupon) resterrors.RestErr {
	coupon := coupon.ConvertToCouponDTO(req)
	return coupon.Create()
}

// func UpdateAdminUser(req *request.UpdateAdminUser) resterrors.RestErr {
// 	adminUser := adminuser.AdminUser{Id: req.Id, Status: req.Status, Name: req.Name}
// 	if req.Email != nil {
// 		adminUser.Email = *req.Email
// 	}
// 	if req.RoleId != nil {
// 		adminUser.RoleId = *req.RoleId
// 	}

// 	return adminUser.Update()
// }

func GetCouponList() ([]coupon.Coupon, resterrors.RestErr) {
	return coupon.GetCouponList()
}

func DeleteCoupon(couponId uint16) resterrors.RestErr {
	return coupon.Delete(uint(couponId))
}

// func GetAdminUserDetails(userId uint16) (response.UserDetails, resterrors.RestErr) {
// 	userDetails, err := adminuser.GetAdminUserDetails(userId)

// 	if err != nil {
// 		return response.UserDetails{}, err
// 	}
// 	return userDetails, nil
// }
