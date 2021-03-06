package handlers

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
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

	randomID := fmt.Sprintf("TKP-%d", rand.Intn(1000))

	coupon := models.Coupon{
		ID:                   randomID,
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

func UpdateCouponHandler(ctx context.Context, id string, newCoupon model.NewCoupon) (*models.Coupon, error) {
	var coupon models.Coupon

	parsedBeginDate, _ := time.Parse(layoutISO, newCoupon.BeginDate)
	parsedExpiredDate, _ := time.Parse(layoutISO, newCoupon.ExpiredDate)

	res := database.DB.Model(&models.Coupon{}).Where("id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	// Validation
	if time.Now().After(coupon.BeginDate) && time.Now().Before(coupon.ExpiredDate) {
		return nil, errors.New("tidak dapat melakukan update kupon yang sedang live")
	}

	coupon = models.Coupon{
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

	res = database.DB.Model(&models.Coupon{}).Where("id = ?", id).Updates(coupon)
	if res.Error != nil {
		return nil, res.Error
	}

	return &coupon, nil
}

func DeleteCouponHandler(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func MyCouponsHandler(ctx context.Context, title string, userID string) ([]*models.Coupon, error) {
	var user models.User
	database.DB.Where("id = ?", userID).First(&user)
	userMemberType := user.MemberType

	var coupons []models.Coupon
	database.DB.Where("member_type = ?", userMemberType).Find(&coupons)

	var res []*models.Coupon

	for _, v := range coupons {
		res = append(res, &v)
	}

	return res, nil
}
