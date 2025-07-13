package posting_usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
	"github.com/samber/lo"
)

var errEmptyCarts = errors.New("no available carts to create posting")

type CreateReq struct {
	CartID int64
}

// TODO пока что один постинг на заказ, нужно придумать как реализовать доставку товаров через разные постинги при необходимости
// TODO также нужно сделать так, чтобы в одном постинге могли ехать товары из разных заказов
func (u *UseCase) Create(ctx context.Context, req CreateReq) error {
	carts, err := u.cartRepo.FindAllByFilter(ctx, dto.FindCartsFilter{
		ID:       lo.ToPtr(req.CartID),
		Status:   lo.ToPtr(internal_entities.CartStatus_BUILDING),
		IsActive: lo.ToPtr(true),
	})
	if err != nil {
		return fmt.Errorf("cartRepo.FindAllByFilter :%w", err)
	}
	if len(carts) == 0 {
		return fmt.Errorf("cartRepo.FindAllByFilter :%w by cart id %d", errEmptyCarts, req.CartID)
	}

	posting := internal_entities.NewPosting(req.CartID)
	posting.SetCartAndPositionsIDs(carts[0].ID, carts[0].PositionsIDs)
	posting.ChangePostingStatus(internal_entities.PostingStatus_BUILT)

	err = u.postingRepo.CreatePosting(ctx, posting)
	if err != nil {
		return fmt.Errorf("postingRepo.CreatePosting :%w", err)
	}

	return nil
}
