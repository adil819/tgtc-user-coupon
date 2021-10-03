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
	id := fmt.Sprintf("USR-%d", rand.Intn(1000))

	user := models.User{
		ID:         id,
		Name:       newUser.Name,
		MemberType: newUser.MemberType,
		CreatedAt:  time.Now(),
	}

	res := database.DB.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
