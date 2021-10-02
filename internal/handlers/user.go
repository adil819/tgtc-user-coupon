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

	// userCoupon := []model.Coupon{

	// }

	user := models.User{
		Name:       newUser.Name,
		MemberType: newUser.MemberType,
		// Coupons:    userCoupon,
		CreatedAt: time.Now(),
	}

	res := database.DB.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
