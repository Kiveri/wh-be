package service_provider

import (
	"github.com/Kiveri/wh-be/internal/adapter/postgres/carts"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/clients"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/companies"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/employees"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/orders"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/positions"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/postings"
	"github.com/Kiveri/wh-be/internal/config"
	"github.com/Kiveri/wh-be/internal/pkg"
	"github.com/Kiveri/wh-be/internal/usecases/cart_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/client_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/company_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/employee_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/order_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/position_usecase"
	"github.com/Kiveri/wh-be/internal/usecases/posting_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	cartRepo      *carts.Repo
	clientRepo    *clients.Repo
	companiesRepo *companies.Repo
	employeesRepo *employees.Repo
	ordersRepo    *orders.Repo
	positionsRepo *positions.Repo
	postingsRepo  *postings.Repo

	cartUseCase     *cart_usecase.UseCase
	clientUseCase   *client_usecase.UseCase
	companyUseCase  *company_usecase.UseCase
	employeeUseCase *employee_usecase.UseCase
	orderUseCase    *order_usecase.UseCase
	positionUseCase *position_usecase.UseCase
	postingUseCase  *posting_usecase.UseCase

	timer *pkg.Timer
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
