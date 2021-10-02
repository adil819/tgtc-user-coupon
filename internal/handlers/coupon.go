package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/kelompok-1-tgtc/tgtc-user-coupon/configs/database"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/graph/model"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/internal/models"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func CreateCouponHandler(ctx context.Context, newCoupon model.NewCoupon) (*models.Coupon, error) {
	parsedBeginDate, _ := time.Parse(layoutISO, newCoupon.BeginDate)
	parsedExpiredDate, _ := time.Parse(layoutISO, newCoupon.ExpiredDate)

	coupon := models.Coupon{
		Title:                newCoupon.Title,
		CouponType:           newCoupon.CouponType,
		BeginDate:            parsedBeginDate,
		ExpiredDate:          parsedExpiredDate,
		Category:             newCoupon.Category,
		Discount:             newCoupon.Discount,
		MaxDiscountAmount:    newCoupon.MaxDiscountAmount,
		MinTransactionAmount: newCoupon.MinTransactionAmount,
		PaymentMethod:        newCoupon.PaymentMethod,
		MemberType:           newCoupon.MemberType,
		ImageURL:             newCoupon.ImageURL,
		Description:          newCoupon.Description,
	}

	res := database.DB.Create(&coupon)
	if res.Error != nil {
		return nil, res.Error
	}

	return &coupon, nil
}

func UpdateCouponHandler(ctx context.Context, id string) (*models.Coupon, error) {
	panic(fmt.Errorf("not implemented"))
}

func DeleteCouponHandler(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func MyCouponsHandler(ctx context.Context, title string, memberType string) ([]*models.Coupon, error) {
	panic(fmt.Errorf("not implemented"))
}
