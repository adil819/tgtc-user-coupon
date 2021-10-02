package handlers

import (
	"context"
	"fmt"

	"github.com/kelompok-1-tgtc/tgtc-user-coupon/graph/model"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/internal/models"
)

func CreateCouponHandler(ctx context.Context, newCoupon model.NewCoupon) (*models.Coupon, error) {
	a := &models.Coupon{
		Title:                "Hehe",
		CouponType:           "Hehe",
		Category:             "Hehe",
		Discount:             25,
		MaxDiscountAmount:    25000,
		MinTransactionAmount: 500000,
		PaymentMethod:        "VA BCA",
		MemberType:           "User",
		ImageURL:             "www.google.com",
		Description:          "Hehe",
	}
	return a, nil
}

func UpdateCouponHandler(ctx context.Context, id string) (*models.Coupon, error) {
	panic(fmt.Errorf("not implemented"))
}

func DeleteCouponHandler(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func ReadAllCouponsHandler(ctx context.Context, userID string) ([]*models.Coupon, error) {
	panic(fmt.Errorf("not implemented"))
}
