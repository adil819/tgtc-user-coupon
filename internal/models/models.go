package models

import (
	"time"

	"github.com/kelompok-1-tgtc/tgtc-user-coupon/configs/database"
	"gorm.io/gorm"
)

type Coupon struct {
	ID                   string
	Title                string
	CouponType           string
	BeginDate            time.Time
	ExpiredDate          time.Time
	Category             string
	Discount             int64
	MaxDiscountAmount    int64
	MinTransactionAmount int64
	PaymentMethod        string
	MemberType           string
	ImageURL             string
	Description          string
	CreatedAt            time.Time
	UpdatedAt            *time.Time
}

type User struct {
	ID         string
	Name       string
	MemberType string
	Coupons    []Coupon `gorm:"many2many:user_has_coupons;"`
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

type UserHasCoupon struct {
	UserID    string `gorm:"primaryKey"`
	CouponID  string `gorm:"primaryKey"`
	Quantity  int
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (UserHasCoupon) BeforeCreate(db *gorm.DB) error {
	err := database.DB.SetupJoinTable(&User{}, "Coupons", &UserHasCoupon{})
	if err != nil {
		return err
	}

	return nil
}
