package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kelompok-1-tgtc/tgtc-user-coupon/graph/generated"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/graph/model"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/internal/handlers"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/internal/models"
)

func (r *couponResolver) BeginDate(ctx context.Context, obj *models.Coupon) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *couponResolver) ExpiredDate(ctx context.Context, obj *models.Coupon) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCoupon(ctx context.Context, newCoupon model.NewCoupon) (*models.Coupon, error) {
	return handlers.CreateCouponHandler(ctx, newCoupon)
}

func (r *mutationResolver) UpdateCoupon(ctx context.Context, id string, newCoupon model.NewCoupon) (*models.Coupon, error) {
	return handlers.UpdateCouponHandler(ctx, id, newCoupon)
}

func (r *mutationResolver) DeleteCoupon(ctx context.Context, id string) (bool, error) {
	return handlers.DeleteCouponHandler(ctx, id)
}

func (r *mutationResolver) CreateUser(ctx context.Context, newUser *model.NewUser) (*models.User, error) {
	return handlers.CreateUserHandler(ctx, newUser)
}

func (r *queryResolver) MyCoupons(ctx context.Context, title string, memberType string) ([]*models.Coupon, error) {
	return handlers.MyCouponsHandler(ctx, title, memberType)
}

// Coupon returns generated.CouponResolver implementation.
func (r *Resolver) Coupon() generated.CouponResolver { return &couponResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type couponResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
