package models

import (
	"time"
)

type Coupon struct {
	ID                   string
	Title                string
	CouponType           string
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

type NewCoupon struct {
	Title                string
	CouponType           string
	Category             string
	Discount             int64
	MaxDiscountAmount    int64
	MinTransactionAmount int64
	PaymentMethod        string
	MemberType           string
	ImageURL             string
	Description          string
}

type User struct {
	ID         string
	Name       string
	MemberType string
	Coupons    []Coupon `gorm:"many2many:user_coupons;"`
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

type UserHasCoupon struct {
	UserID    string `gorm:"primaryKey"`
	CouponID  string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}
