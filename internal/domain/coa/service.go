package coa

import (
	"context"
	"database/sql"

	"ledger-service/internal/pkg/app"
	"ledger-service/pb/ledgers"
)

type Service struct {
	Db *sql.DB
	ledgers.UnimplementedChartOfAccountServiceServer
}

func (a *Service) Create(ctx context.Context, in *ledgers.ChartOfAccount) (*ledgers.ChartOfAccount, error) {
	repo := Repo{db: a.Db, pb: ledgers.ChartOfAccount{
		CompanyId:   ctx.Value(app.Ctx("company_id")).(string),
		AccountType: in.AccountType,
		Parent:      in.Parent,
		Name:        in.Name,
	}}

	err := repo.Create(ctx)
	if err != nil {
		return nil, err
	}

	return &repo.pb, nil
}

func (a *Service) Update(ctx context.Context, in *ledgers.ChartOfAccount) (*ledgers.ChartOfAccount, error) {
	repo := Repo{db: a.Db, pb: ledgers.ChartOfAccount{
		Id:          in.Id,
		CompanyId:   ctx.Value(app.Ctx("company_id")).(string),
		AccountType: in.AccountType,
		Parent:      in.Parent,
		Name:        in.Name,
	}}

	err := repo.Update(ctx)
	if err != nil {
		return nil, err
	}

	return &repo.pb, nil
}

func (a *Service) List(ctx context.Context, in *ledgers.EmptyMessage) (*ledgers.ChartOfAccounts, error) {
	repo := Repo{db: a.Db}
	return repo.List(ctx)
}

func (a *Service) Delete(ctx context.Context, in *ledgers.Id) (*ledgers.BoolMessage, error) {
	repo := Repo{db: a.Db, pb: ledgers.ChartOfAccount{
		Id: in.Id,
	}}

	isTrue, err := repo.Delete(ctx)
	if err != nil {
		return nil, err
	}

	return &ledgers.BoolMessage{IsTrue: isTrue}, nil
}
