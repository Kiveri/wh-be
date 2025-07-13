package service_provider

import (
	"context"
	"github.com/Kiveri/wh-be/internal/usecases/cart_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/client_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/company_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/employee_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/order_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/position_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/posting_usecase"
)

func (sp *ServiceProvider) getCartUseCase(ctx context.Context) *cart_usecase.UseCase {
	if sp.cartUseCase == nil {
		sp.cartUseCase = cart_usecase.NewUseCase(
			sp.getCartRepo(ctx),
			sp.getPositionRepo(ctx),
			sp.getTimer(),
		)
	}

	return sp.cartUseCase
}

func (sp *ServiceProvider) getClientUseCase(ctx context.Context) *client_usecase.UseCase {
	if sp.clientUseCase == nil {
		sp.clientUseCase = client_usecase.NewUseCase(
			sp.getClientRepo(ctx),
		)
	}

	return sp.clientUseCase
}

func (sp *ServiceProvider) getCompanyUseCase(ctx context.Context) *company_usecase.UseCase {
	if sp.companyUseCase == nil {
		sp.companyUseCase = company_usecase.NewUseCase(
			sp.getCompanyRepo(ctx),
		)
	}

	return sp.companyUseCase
}

func (sp *ServiceProvider) getEmployeeUseCase(ctx context.Context) *employee_usecase.UseCase {
	if sp.employeeUseCase == nil {
		sp.employeeUseCase = employee_usecase.NewUseCase(
			sp.getEmployeeRepo(ctx),
		)
	}

	return sp.employeeUseCase
}

func (sp *ServiceProvider) getOrderUseCase(ctx context.Context) *order_usecase.UseCase {
	if sp.orderUseCase == nil {
		sp.orderUseCase = order_usecase.NewUseCase(
			sp.getOrderRepo(ctx),
		)
	}

	return sp.orderUseCase
}

func (sp *ServiceProvider) getPositionUseCase(ctx context.Context) *position_usecase.UseCase {
	if sp.positionUseCase == nil {
		sp.positionUseCase = position_usecase.NewUseCase(
			sp.getPositionRepo(ctx),
		)
	}

	return sp.positionUseCase
}

func (sp *ServiceProvider) getPostingUseCase(ctx context.Context) *posting_usecase.UseCase {
	if sp.postingUseCase == nil {
		sp.postingUseCase = posting_usecase.NewUseCase(
			sp.getPostingRepo(ctx),
		)
	}

	return sp.postingUseCase
}
