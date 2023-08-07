package scheme

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Create uuid extension",
		Script:      `CREATE EXTENSION "uuid-ossp";`,
	},
	{
		Version:     2,
		Description: "Create table account_types",
		Script: `CREATE TABLE account_types(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			name varchar(45) NOT NULL
		);`,
	},
	{
		Version:     3,
		Description: "Create table chart_of_accounts",
		Script: `CREATE TABLE chart_of_accounts(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			account_type_id uuid NOT NULL,
			parent_id uuid,
			code char(8) NOT NULL,
			name varchar(45) NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, code),
			CONSTRAINT fk_chart_of_accounts_to_account_types FOREIGN KEY(account_type_id) REFERENCES account_types(id),
			CONSTRAINT fk_chart_of_accounts_to_parents FOREIGN KEY(parent_id) REFERENCES chart_of_accounts(id) 
		);`,
	},
	{
		Version:     4,
		Description: "Create table account_balances",
		Script: `CREATE TABLE account_balances(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			account_id uuid NOT NULL,
			period date NOT NULL,
			beginning_balance double NOT NULL,
			debit double NOT NULL,
			credit double NOT NULL,
			ending_balance double NOT NULL,
			closing_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			closing_by uuid NOT NULL,
			UNIQUE(company_id, period),
			CONSTRAINT fk_account_balances_to_chart_of_accounts FOREIGN KEY(account_id) REFERENCES chart_of_accounts(id)
		);`,
	},
	{
		Version:     5,
		Description: "Create table debt_balances",
		Script: `CREATE TABLE debt_balances(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			period date NOT NULL,
			beginning_balance double NOT NULL,
			amount double NOT NULL,
			ending_balance double NOT NULL,
			closing_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			closing_by uuid NOT NULL,
			UNIQUE(company_id, period)
		);`,
	},
	{
		Version:     6,
		Description: "Create table ar_balances",
		Script: `CREATE TABLE ar_balances(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			period date NOT NULL,
			beginning_balance double NOT NULL,
			amount double NOT NULL,
			ending_balance double NOT NULL,
			closing_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			closing_by uuid NOT NULL,
			UNIQUE(company_id, period)
		);`,
	},
	{
		Version:     7,
		Description: "Create table receivables",
		Script: `CREATE TABLE receivables(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			sales_id uuid NOT NULL,
			customer_id uuid NOT NULL,
			due_date date NOT NULL, 
			amount_owed double NOT NULL,
			amount_paid double NOT NULL,
			payment_status char(1) NOT NULL,
			UNIQUE(company_id, sales_id)
		);`,
	},
	{
		Version:     8,
		Description: "Create table debts",
		Script: `CREATE TABLE debts(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			purchase_id uuid NOT NULL,
			supplier_id uuid NOT NULL,
			due_date date NOT NULL, 
			amount_owed double NOT NULL,
			amount_paid double NOT NULL,
			payment_status char(1) NOT NULL,
			UNIQUE(company_id, purchase_id)
		);`,
	},
	{
		Version:     9,
		Description: "Create table cash_receipts",
		Script: `CREATE TABLE cash_receipts(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			code char(13) NOT NULL,
			date date NOT NULL,
			reference_no char(13),
			debt_id uuid,
			sales_id uuid,
			amount double NOT NULL,
			source varchar(45),
			description varchar(100),
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, code)
		);`,
	},
	{
		Version:     10,
		Description: "Create table cash_payments",
		Script: `CREATE TABLE cash_payments(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			code char(13) NOT NULL,
			date date NOT NULL,
			reference_no char(13),
			ar_id uuid,
			purchase_id uuid,
			amount double NOT NULL,
			payee varchar(45),
			description varchar(100),
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, code)
		);`,
	},
	{
		Version:     11,
		Description: "Create table bank_accounts",
		Script: `CREATE TABLE bank_accounts(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			bank_name varchar(45) NOT NULL,
			account_number char(20) NOT NULL,
			account_name varchar(45) NOT NULL,
			account_type char(1) NOT NULL,
			owner_type char(1) NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, account_number)
		);`,
	},
	{
		Version:     12,
		Description: "Create table bank_receipts",
		Script: `CREATE TABLE bank_receipts(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			code char(13) NOT NULL,
			date date NOT NULL,
			reference_no char(13),
			debt_id uuid,
			sales_id uuid,
			amount double NOT NULL,
			source varchar(45),
			sender_bank_account_id uuid NOT NULL,
			receiver_bank_account_id uuid NOT NULL,
			description varchar(100),
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, code),
			CONSTRAINT fk_bank_receipts_to_sender_bank_accounts FOREIGN KEY(sender_bank_account_id) REFERENCES bank_accounts(id),
			CONSTRAINT fk_bank_receipts_to_receiver_bank_accounts FOREIGN KEY(receiver_bank_account_id) REFERENCES bank_accounts(id)
		);`,
	},
	{
		Version:     13,
		Description: "Create table bank_payments",
		Script: `CREATE TABLE bank_payments(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			code char(13) NOT NULL,
			date date NOT NULL,
			reference_no char(13),
			ar_id uuid,
			purcahse_id uuid,
			amount double NOT NULL,
			payee varchar(45),
			sender_bank_account_id uuid NOT NULL,
			receiver_bank_account_id uuid NOT NULL,
			description varchar(100),
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, code),
			CONSTRAINT fk_bank_receipts_to_sender_bank_accounts FOREIGN KEY(sender_bank_account_id) REFERENCES bank_accounts(id),
			CONSTRAINT fk_bank_receipts_to_receiver_bank_accounts FOREIGN KEY(receiver_bank_account_id) REFERENCES bank_accounts(id)
		);`,
	},
	{
		Version:     14,
		Description: "Create table general_journals",
		Script: `CREATE TABLE general_journals(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			code char(13) NOT NULL,
			date date NOT NULL,
			transaction_type char(2) NOT NULL,
			reference_no char(13),
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, code)
		);`,
	},
	{
		Version:     15,
		Description: "Create table general_journal_details",
		Script: `CREATE TABLE general_journal_details(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			general_journal_id uuid NOT NULL,
			account_id uuid,
			name varchar(45) NOT NULL,
			amount double NOT NULL,
			debit_credit char(1) NOT NULL,
			CONSTRAINT fk_general_journal_details_to_chart_of_accounts FOREIGN KEY(account_id) REFERENCES chart_of_accounts(id),
			CONSTRAINT fk_general_journal_details_to_general_journals FOREIGN KEY(general_journal_id) REFERENCES general_journals(id)
		);`,
	},
	{
		Version:     16,
		Description: "Create table general_ledgers",
		Script: `CREATE TABLE general_ledgers(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			code char(13) NOT NULL,
			general_journal_id uuid NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, code),
			CONSTRAINT fk_general_ledgers_to_general_journals FOREIGN KEY(general_journal_id) REFERENCES general_journals(id)
		);`,
	},
	{
		Version:     17,
		Description: "Create table general_ledger_details",
		Script: `CREATE TABLE general_ledger_details(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			general_ledger_id uuid NOT NULL,
			general_journal_detail_id uuid NOT NULL,
			account_id uuid NOT NULL,
			CONSTRAINT fk_general_ledger_details_to_general_ledgers FOREIGN KEY(general_ledger_id) REFERENCES general_ledgers(id),
			CONSTRAINT fk_general_ledger_details_to_general_journal_details FOREIGN KEY(general_journal_detail_id) REFERENCES general_journal_details(id),
			CONSTRAINT fk_general_ledger_details_to_chart_of_accounts FOREIGN KEY(account_id) REFERENCES chart_of_accounts(id)
		);`,
	},
	{
		Version:     18,
		Description: "Create table journal_entries",
		Script: `CREATE TABLE journal_entries(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			code char(13) NOT NULL,
			date date NOT NULL,
			description varchar(100) NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', NOW()),
			created_by uuid NOT NULL,
			updated_by uuid NOT NULL,
			UNIQUE(company_id, code)
		);`,
	},
	{
		Version:     19,
		Description: "Create table journal_entry_details",
		Script: `CREATE TABLE journal_entry_details(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			journal_entry_id uuid NOT NULL,
			name varchar(45) NOT NULL,
			amount double NOT NULL,
			debit_credit char(1) NOT NULL,
			CONSTRAINT fk_journal_entry_details_to_journal_entries FOREIGN KEY(journal_entry_id) REFERENCES journal_entries(id)
		);`,
	},
	{
		Version:     20,
		Description: "Create table acc_configurations",
		Script: `CREATE TABLE acc_configurations(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			name varchar(45) NOT NULL UNIQUE
		);`,
	},
	{
		Version:     21,
		Description: "Create table acc_configuration_details",
		Script: `CREATE TABLE acc_configuration_details(
			id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
			company_id uuid NOT NULL,
			acc_configuration_id uuid NOT NULL,
			account_id uuid NOT NULL,
			UNIQUE(company_id, acc_configuration_id),
			UNIQUE(company_id, acc_configuration_id, account_id),
			CONSTRAINT fk_acc_configuration_details_to_acc_configurations FOREIGN KEY(acc_configuration_id) REFERENCES acc_configurations(id),
			CONSTRAINT fk_acc_configuration_details_to_chart_of_accounts FOREIGN KEY(account_id) REFERENCES chart_of_accounts(id)
		);`,
	},
}

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sql.DB) error {
	driver := darwin.NewGenericDriver(db, darwin.PostgresDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
