package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/azinudinachzab/bq-loan-be-v2/model"
)

const (
	loanTypeTab = "loan_types"
)

func (p *PgRepository) CreateLoanType(ctx context.Context, loanType model.LoanType) error {
	q := fmt.Sprintf(`INSERT INTO %s (name,margin,is_add_balance,created_at,updated_at) VALUES (?,?,?,?,?);`, loanTypeTab)
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, loanType.Name, loanType.Margin, loanType.IsAddBalance, now, now); err != nil {
		return err
	}

	return nil
}
func (p *PgRepository) GetLoanTypes(ctx context.Context) ([]model.LoanType, error) {
	q := fmt.Sprintf(`SELECT id, name, margin,is_add_balance,created_at, updated_at FROM %s`, loanTypeTab)
	rows, err := p.dbCore.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rowClose(rows)

	loanTypes := make([]model.LoanType, 0)
	for rows.Next() {
		var (
			id       uint32
			name     sql.NullString
			margin   sql.NullFloat64
			cat, uat sql.NullTime
			iab      sql.NullInt32
		)

		if err := rows.Scan(&id, &name, &margin, &iab, &cat, &uat); err != nil {
			return nil, err
		}

		loanTypes = append(loanTypes, model.LoanType{
			ID:           id,
			Name:         name.String,
			Margin:       margin.Float64,
			IsAddBalance: int(iab.Int32),
			CreatedAt:    cat.Time.Format(time.RFC3339),
			UpdatedAt:    uat.Time.Format(time.RFC3339),
		})
		if rows.Err() != nil {
			return nil, err
		}
	}

	return loanTypes, nil
}
func (p *PgRepository) GetLoanType(ctx context.Context, id uint32) (model.LoanType, error) {
	q := fmt.Sprintf(`SELECT id,name,margin,is_add_balance,created_at,updated_at FROM %s WHERE id=?`, loanTypeTab)
	var (
		ids      uint32
		name     sql.NullString
		margin   sql.NullFloat64
		cat, uat sql.NullTime
		iab      sql.NullInt32
	)
	if err := p.dbCore.QueryRowContext(ctx, q, id).Scan(&ids, &name, &margin, &iab, &cat, &uat); err != nil {
		return model.LoanType{}, err
	}
	return model.LoanType{
		ID:           ids,
		Name:         name.String,
		Margin:       margin.Float64,
		IsAddBalance: int(iab.Int32),
		CreatedAt:    cat.Time.Format(time.RFC3339),
		UpdatedAt:    uat.Time.Format(time.RFC3339),
	}, nil
}
func (p *PgRepository) UpdateLoanType(ctx context.Context, id uint32, lt model.LoanType) error {
	q := fmt.Sprintf(`UPDATE %s SET name=?, margin=?, is_add_balance=?, updated_at=? WHERE id=?;`, loanTypeTab)
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, lt.Name, lt.Margin, lt.IsAddBalance, now, id); err != nil {
		return err
	}

	return nil
}
func (p *PgRepository) DeleteLoanType(ctx context.Context, id uint32) error {
	q := fmt.Sprintf(`DELETE FROM %s WHERE id=?;`, loanTypeTab)

	if _, err := p.dbCore.ExecContext(ctx, q, id); err != nil {
		return err
	}

	return nil
}
