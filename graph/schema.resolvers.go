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

func (r *queryResolver) MyCoupons(ctx context.Context, title string, userID string) ([]*models.Coupon, error) {
	return handlers.MyCouponsHandler(ctx, title, userID)
}

func (r *queryResolver) GetPaginationCoupons(ctx context.Context, input model.Pagination) (*model.PaginationResultCoupon, error) {
	panic(fmt.Errorf("not implemented"))
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetPaginationCoupon(ctx context.Context, input model.Pagination) (*model.PaginationResultCoupon, error) {
	panic(fmt.Errorf("not implemented"))
}
