package services

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/n340r/backend-notes/toptal/internal/app/domain"
)

type CartService struct {
	cartRepo CartRepository
}

func NewCartService(repo CartRepository) CartService {
	return CartService{
		cartRepo: repo,
	}
}

func (s CartService) UpdateCartAndStocks(ctx context.Context, cart domain.Cart) (domain.Cart, error) {
	err := s.cartRepo.UpdateCartAndStocks(ctx, cart)

	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to update cart and stocks %w", err)
	}

	updatedCart, err := s.cartRepo.GetCart(ctx, cart.UserID())
	spew.Dump(err)
	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to get updated cart: %w", err)
	}

	return updatedCart, nil
}

func (s CartService) Checkout(ctx context.Context, userID int) error {
	return s.cartRepo.DeleteCart(ctx, userID)
}
