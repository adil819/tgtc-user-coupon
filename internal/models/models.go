package models

import (
	"time"
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
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}
