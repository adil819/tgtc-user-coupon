package handlers

import (
	"context"
	"time"

	"github.com/kelompok-1-tgtc/tgtc-user-coupon/configs/database"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/graph/model"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/internal/models"
)

func CreateUserHandler(ctx context.Context, newUser *model.NewUser) (*models.User, error) {
	// userCoupon := []*model.NewCoupon{}
	a := models.Coupon{
		ID:                   "123",
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
	userCoupon := []models.Coupon{a}

	user := models.User{
		ID:         "hehe",
		Name:       newUser.Name,
		MemberType: newUser.MemberType,
		// Coupons:    userCoupon,
		Coupons:   userCoupon,
		CreatedAt: time.Now(),
	}

	res := database.DB.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
