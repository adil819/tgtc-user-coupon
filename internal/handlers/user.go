package handlers

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/kelompok-1-tgtc/tgtc-user-coupon/configs/database"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/graph/model"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/internal/models"
)

func CreateUserHandler(ctx context.Context, newUser *model.NewUser) (*models.User, error) {
	var userCoupon []models.Coupon

	for _, v := range newUser.Coupons {
		id := fmt.Sprintf("TKP-%d", rand.Intn(1000))

		parsedBeginDate, _ := time.Parse(layoutISO, v.BeginDate)
		parsedExpiredDate, _ := time.Parse(layoutISO, v.ExpiredDate)

		tempCoupon := models.Coupon{
			ID:                   id,
			Title:                v.Title,
			CouponType:           v.CouponType,
			BeginDate:            parsedBeginDate,
			ExpiredDate:          parsedExpiredDate,
			Category:             v.Category,
			Discount:             v.Discount,
			MaxDiscountAmount:    v.MaxDiscountAmount,
			MinTransactionAmount: v.MinTransactionAmount,
			PaymentMethod:        v.PaymentMethod,
			MemberType:           v.MemberType,
			ImageURL:             v.ImageURL,
			Description:          v.Description,
		}
		userCoupon = append(userCoupon, tempCoupon)
	}

	id := fmt.Sprintf("USR-%d", rand.Intn(1000))

	user := models.User{
		ID:         id,
		Name:       newUser.Name,
		MemberType: newUser.MemberType,
		Coupons:    userCoupon,
		CreatedAt:  time.Now(),
	}

	res := database.DB.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
