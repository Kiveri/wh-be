package service_provider

import (
	"context"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/carts"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/clients"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/companies"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/employees"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/orders"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/positions"
	"github.com/Kiveri/wh-be/internal/adapter/postgres/postings"
)

func (sp *ServiceProvider) getCartRepo(ctx context.Context) *carts.Repo {
	if sp.cartRepo == nil {
		sp.cartRepo = carts.NewRepo(
			sp.getDbCluster(ctx),
			sp.getTimer(),
		)
	}

	return sp.cartRepo
}

func (sp *ServiceProvider) getClientRepo(ctx context.Context) *clients.Repo {
	if sp.clientRepo == nil {
		sp.clientRepo = clients.NewRepo(
			sp.getDbCluster(ctx),
			sp.getTimer(),
		)
	}

	return sp.clientRepo
}

func (sp *ServiceProvider) getCompanyRepo(ctx context.Context) *companies.Repo {
	if sp.companiesRepo == nil {
		sp.companiesRepo = companies.NewRepo(
			sp.getDbCluster(ctx),
			sp.getTimer(),
		)
	}

	return sp.companiesRepo
}

func (sp *ServiceProvider) getEmployeeRepo(ctx context.Context) *employees.Repo {
	if sp.employeesRepo == nil {
		sp.employeesRepo = employees.NewRepo(
			sp.getDbCluster(ctx),
			sp.getTimer(),
		)
	}

	return sp.employeesRepo
}

func (sp *ServiceProvider) getOrderRepo(ctx context.Context) *orders.Repo {
	if sp.ordersRepo == nil {
		sp.ordersRepo = orders.NewRepo(
			sp.getDbCluster(ctx),
			sp.getTimer(),
		)
	}

	return sp.ordersRepo
}

func (sp *ServiceProvider) getPositionRepo(ctx context.Context) *positions.Repo {
	if sp.positionsRepo == nil {
		sp.positionsRepo = positions.NewRepo(
			sp.getDbCluster(ctx),
			sp.getTimer(),
		)
	}

	return sp.positionsRepo
}

func (sp *ServiceProvider) getPostingRepo(ctx context.Context) *postings.Repo {
	if sp.postingsRepo == nil {
		sp.postingsRepo = postings.NewRepo(
			sp.getDbCluster(ctx),
			sp.getTimer(),
		)
	}

	return sp.postingsRepo
}
