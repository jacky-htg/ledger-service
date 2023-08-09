package acc_configuration

import (
	"context"
	"database/sql"
	"ledger-service/pb/ledgers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Db *sql.DB
	ledgers.UnimplementedAccConfigurationServiceServer
}

func (a *Service) List(in *ledgers.EmptyMessage, stream ledgers.AccConfigurationService_ListServer) error {
	repo := Repo{db: a.Db}
	rows, err := repo.List(stream.Context())
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var detail ledgers.AccConfigurationDetail
		err = rows.Scan(
			&detail.Id, &detail.AccConfiguration.Id, &detail.AccConfiguration.Name,
			&detail.ChartOfAccount.Id, &detail.ChartOfAccount.Code, &detail.ChartOfAccount.Name,
		)

		if err != nil {
			return status.Errorf(codes.Internal, "Scan List acc configuration: %v", err)
		}

		if err := stream.Send(&detail); err != nil {
			return err
		}
	}

	return nil
}

func (a *Service) Create(ctx context.Context, in *ledgers.AccConfigurationDetail) (*ledgers.AccConfigurationDetail, error) {
	repo := Repo{db: a.Db, pb: ledgers.AccConfigurationDetail{
		AccConfiguration: &ledgers.AccConfiguration{Id: in.AccConfiguration.Id},
		ChartOfAccount:   &ledgers.ChartOfAccount{Id: in.ChartOfAccount.Id},
	}}

	err := repo.Create(ctx)
	if err != nil {
		return nil, err
	}

	return &repo.pb, nil
}

func (a *Service) Update(ctx context.Context, in *ledgers.AccConfigurationDetail) (*ledgers.AccConfigurationDetail, error) {

	repo := Repo{db: a.Db, pb: ledgers.AccConfigurationDetail{
		Id:               in.Id,
		AccConfiguration: &ledgers.AccConfiguration{Id: in.AccConfiguration.Id},
		ChartOfAccount:   &ledgers.ChartOfAccount{Id: in.ChartOfAccount.Id},
	}}

	err := repo.Update(ctx)
	if err != nil {
		return nil, err
	}

	return &repo.pb, nil
}

func (a *Service) Delete(ctx context.Context, in *ledgers.Id) (*ledgers.BoolMessage, error) {
	repo := Repo{db: a.Db, pb: ledgers.AccConfigurationDetail{
		Id: in.Id,
	}}

	isTrue, err := repo.Delete(ctx)
	if err != nil {
		return nil, err
	}

	return &ledgers.BoolMessage{IsTrue: isTrue}, nil
}
