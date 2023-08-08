package coa

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
	pb ledgers.ChartOfAccount
}

func (a *Repo) Create(ctx context.Context) error {
	query := `
		INSERT INTO chart_of_accounts (company_id, account_type_id, parent_id, code, name, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, updated_at, created_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement create coa: %v", err)
	}
	defer stmt.Close()

	var parentId sql.NullString
	if len(a.pb.Parent.Id) > 0 {
		parentId.String = a.pb.Parent.Id
		parentId.Valid = true
	}

	a.pb.CreatedBy = ctx.Value(app.Ctx("user_id")).(string)
	a.pb.UpdatedBy = ctx.Value(app.Ctx("user_id")).(string)

	err = stmt.QueryRowContext(ctx,
		a.pb.CompanyId,
		a.pb.AccountType.Id,
		parentId,
		a.pb.Code,
		a.pb.Name,
		a.pb.CreatedBy,
		a.pb.UpdatedBy,
	).Scan(&a.pb.Id, &a.pb.UpdatedAt, &a.pb.CreatedAt)

	if err != nil {
		return status.Errorf(codes.Internal, "QueryRowContext Create Coa: %v", err)
	}

	return nil
}

func (a *Repo) Update(ctx context.Context) error {
	query := `
		UPDATE chart_of_accounts SET 
		account_type_id = $1,
		parent_id = $2, 
		name = $3, 
		updated_at = timezone('utc', NOW()), 
		updated_by = $4
		WHERE id = $5 AND company_id = $6
		RETURNING created_by, created_at, updated_at
		`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement update coa: %v", err)
	}
	defer stmt.Close()

	var parentId sql.NullString
	if len(a.pb.Parent.Id) > 0 {
		parentId.String = a.pb.Parent.Id
		parentId.Valid = true
	}

	a.pb.UpdatedBy = ctx.Value(app.Ctx("user_id")).(string)

	err = stmt.QueryRowContext(ctx,
		a.pb.AccountType.Id,
		parentId,
		a.pb.Name,
		a.pb.UpdatedBy,
		a.pb.Id,
		ctx.Value(app.Ctx("user_id")).(string),
	).Scan(&a.pb.CreatedBy, &a.pb.CreatedAt, &a.pb.UpdatedAt)

	if err != nil {
		return status.Errorf(codes.Internal, "QueryRowContext Update Coa: %v", err)
	}

	return nil
}

func (a *Repo) List(ctx context.Context) (*ledgers.ChartOfAccounts, error) {
	var output ledgers.ChartOfAccounts
	query := `
		SELECT 
			chart_of_accounts.id, account_type_id, account_types.name account_type_name, parent_id, code, chart_of_accounts.name, 
			created_at, created_by, updated_at, updated_by
		FROM chart_of_accounts
		JOIN account_types ON chart_of_accounts.account_type_id = account_types.id
		WHERE company_id = $1
		ORDER BY parent_id, created_at
	`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Prepare statement List coa: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, ctx.Value(app.Ctx("company_id")).(string))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "QueryContext List coa: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var coa ledgers.ChartOfAccount
		err = rows.Scan(
			&coa.Id, &coa.AccountType.Id, &coa.AccountType.Name, &coa.Parent.Id, &coa.Code, &coa.Name,
			&coa.CreatedAt, &coa.CreatedBy, &coa.UpdatedAt, &coa.UpdatedBy,
		)

		if err != nil {
			return nil, status.Errorf(codes.Internal, "Scan List coa: %v", err)
		}

		output.ChartOfAccount = append(output.ChartOfAccount, &coa)
	}

	if rows.Err() != nil {
		return nil, status.Errorf(codes.Internal, "Rows Err coa: %v", err)
	}

	return &output, nil
}

func (a *Repo) Delete(ctx context.Context) (bool, error) {
	query := `DELETE FROM chart_of_accounts WHERE id = $1 AND company_id = $2`

	stmt, err := a.tx.PrepareContext(ctx, query)
	if err != nil {
		return false, status.Errorf(codes.Internal, "Prepare statement delete coa: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, a.pb.Id, ctx.Value(app.Ctx("company_id")).(string))

	if err != nil {
		return false, status.Errorf(codes.Internal, "ExecContext Coa: %v", err)
	}

	return true, nil
}
