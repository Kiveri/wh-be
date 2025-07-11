package position_usecase

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model"
	"github.com/Kiveri/wh-be/internal/pkg"
	"github.com/samber/lo"
	"time"
)

type CreatePositionReq struct {
	ExternalID     int64
	Barcode        int64
	Name           string
	Manufacturer   string
	Price          int64
	Type           model.PositionType
	ProductionDate *pkg.Date
	ExpirationDate *pkg.Date
}

func (u *UseCase) CreatePosition(ctx context.Context, req CreatePositionReq) error {
	position := model.NewPosition(req.ExternalID, req.Barcode, req.Name, req.Manufacturer, req.Price, req.Type)

	if req.ProductionDate != nil {
		prodDate := time.Date(req.ProductionDate.Year, time.Month(req.ProductionDate.Month), req.ProductionDate.Day, 0, 0, 0, 0, nil)
		position.ProductionDate = lo.ToPtr(prodDate)
	}
	if req.ExpirationDate != nil {
		expDate := time.Date(req.ExpirationDate.Year, time.Month(req.ExpirationDate.Month), req.ExpirationDate.Day, 0, 0, 0, 0, nil)
		position.ExpirationDate = lo.ToPtr(expDate)
	}

	err := u.positionRepo.CreatePosition(ctx, position)
	if err != nil {
		return fmt.Errorf("positionRepo.CreatePosition: %w", err)
	}

	return nil
}
