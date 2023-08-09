package acc_configuration

import (
	"context"
	"database/sql"
	"ledger-service/internal/pkg/app"
	"ledger-service/pb/ledgers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repo struct {
	db *sql.DB
	tx *sql.Tx
	pb ledgers.AccConfigurationDetail
}

func (a *Repo) Create(ctx context.Context) error {
	query := `
		INSERT INTO acc_configuration_details (company_id, acc_configuration_id, account_id)
		VALUES ($1, $2, $3)
		RETURNING id
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement create acc configuration: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		ctx.Value(app.Ctx("company_id")).(string),
		a.pb.AccConfiguration.Id,
		a.pb.ChartOfAccount.Id,
	).Scan(&a.pb.Id)

	if err != nil {
		return status.Errorf(codes.Internal, "QueryRowContext Create acc configuration: %v", err)
	}

	return nil
}

func (a *Repo) Update(ctx context.Context) error {
	query := `
		UPDATE acc_configuration_details SET
		account_id = $1,
		WHERE id = $2 AND company_id = $3
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement update acc configuration: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		a.pb.ChartOfAccount.Id,
		a.pb.Id,
		ctx.Value(app.Ctx("company_id")).(string),
	)

	if err != nil {
		return status.Errorf(codes.Internal, "ExecContext Update acc configuration: %v", err)
	}

	return nil
}

func (a *Repo) List(ctx context.Context) (*sql.Rows, error) {
	query := `
		SELECT 
			acd.id, ac.id, ac.name acc_configuration_name, 
			coa.id coa_id, coa.code coa_code, coa.name coa_name 
		FROM acc_configurations ac
		JOIN acc_configuration_details acd ON ac.id = acd.acc_configuration_id
		JOIN chart_of_accounts coa ON coa.id = acd.account_id
		WHERE acd.company_id = $1 AND coa.company_id = $2
		ORDER BY ac.name, acd.created_at
	`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Prepare statement List acc configuration: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, ctx.Value(app.Ctx("company_id")).(string), ctx.Value(app.Ctx("company_id")).(string))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "QueryContext List acc configuration: %v", err)
	}

	return rows, nil
}

func (a *Repo) Delete(ctx context.Context) (bool, error) {
	query := `DELETE FROM acc_configuration_details WHERE id = $1 AND company_id = $2`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return false, status.Errorf(codes.Internal, "Prepare statement delete acc configuration: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, a.pb.Id, ctx.Value(app.Ctx("company_id")).(string))

	if err != nil {
		return false, status.Errorf(codes.Internal, "ExecContext acc configuration: %v", err)
	}

	return true, nil
}
