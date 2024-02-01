package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"math"
	"time"
)

const (
	loanDetailTab = "loan_details"
)

//func (p *PgRepository) UpdateLoanDetail(ctx context.Context, detail model.LoanDetail) error {
//	tx, err := p.dbCore.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
//	if err != nil {
//		return err
//	}
//	defer tx.Rollback()
//
//	details, err := p.GetLoanDetails(ctx, detail.LoanGeneralID, tx)
//	if err != nil {
//		return err
//	}
//
//	var (
//		unpaid, count float64
//	)
//	for _, val := range details {
//		if val.Status == 0 {
//			unpaid += val.Amount
//			count++
//		}
//	}
//	adjustedUnpaidLoan := (unpaid) / (count - 1)
//
//	if err := tx.Commit(); err != nil {
//		return err
//	}
//
//	return nil
//}

func (p *PgRepository) DeleteLoanDetail(ctx context.Context, id, lgid uint32) error {
	tx, err := p.dbCore.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	details, err := p.GetLoanDetails(ctx, lgid, tx)
	if err != nil {
		return err
	}

	var (
		unpaid, count float64
	)
	for _, val := range details {
		if val.ID == id && val.Status == 1 {
			return errors.New("paid loan cannot be deleted")
		}

		if val.Status == 0 {
			unpaid += val.Amount
			count++
		}
	}
	if count == 1 {
		return errors.New("last loan cannot be deleted")
	}
	adjustedUnpaidLoan := math.Ceil(unpaid / (count - 1))

	q := fmt.Sprintf(`UPDATE %s SET amount=? WHERE loan_general_id=? AND status=?`, loanDetailTab)
	if _, err := tx.ExecContext(ctx, q, adjustedUnpaidLoan, lgid, 0); err != nil {
		return err
	}

	q = fmt.Sprintf(`UPDATE %s SET tenor=? WHERE id=?`, loanGeneralTab)
	if _, err := tx.ExecContext(ctx, q, len(details)-1, lgid); err != nil {
		return err
	}

	q = fmt.Sprintf(`DELETE FROM %s WHERE id=?`, loanDetailTab)
	if _, err := tx.ExecContext(ctx, q, id); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *PgRepository) GetLoanDetail(ctx context.Context, id uint32) (model.LoanDetail, error) {
	q := fmt.Sprintf(`SELECT id,loan_general_id,amount,datetime,status,created_at,updated_at FROM %s WHERE id=?`, loanDetailTab)
	var (
		ids, lgid    uint32
		amount       sql.NullFloat64
		status       sql.NullInt32
		dt, cat, uat sql.NullTime
	)
	if err := p.dbCore.QueryRowContext(ctx, q, id).Scan(&ids, &lgid, &amount, &dt, &status, &cat, &uat); err != nil {
		return model.LoanDetail{}, err
	}
	return model.LoanDetail{
		ID:            ids,
		LoanGeneralID: lgid,
		Amount:        amount.Float64,
		Datetime:      dt.Time.Format(time.DateTime),
		Status:        int(status.Int32),
		CreatedAt:     cat.Time.Format(time.RFC3339),
		UpdatedAt:     uat.Time.Format(time.RFC3339),
	}, nil
}

func (p *PgRepository) GetLoanDetails(ctx context.Context, generalID uint32, tx *sql.Tx) ([]model.LoanDetail, error) {
	q := fmt.Sprintf(`SELECT id, loan_general_id, amount, datetime, status,
       created_at, updated_at FROM %s WHERE loan_general_id=?`, loanDetailTab)
	executor := p.dbCore.QueryContext
	if tx != nil {
		executor = tx.QueryContext
	}
	rows, err := executor(ctx, q, generalID)
	if err != nil {
		return nil, err
	}
	defer rowClose(rows)

	loanDetails := make([]model.LoanDetail, 0)
	for rows.Next() {
		var (
			id, lgid     uint32
			amount       sql.NullFloat64
			status       sql.NullInt32
			dt, cat, uat sql.NullTime
		)

		if err := rows.Scan(&id, &lgid, &amount, &dt, &status, &cat, &uat); err != nil {
			return nil, err
		}

		loanDetails = append(loanDetails, model.LoanDetail{
			ID:            id,
			LoanGeneralID: lgid,
			Amount:        amount.Float64,
			Datetime:      dt.Time.Format(time.DateTime),
			Status:        int(status.Int32),
			CreatedAt:     cat.Time.Format(time.RFC3339),
			UpdatedAt:     uat.Time.Format(time.RFC3339),
		})
		if rows.Err() != nil {
			return nil, err
		}
	}

	return loanDetails, nil
}

func (p *PgRepository) UpdateLoanDetailStatus(ctx context.Context, id, lgid uint32, amount float64) error {
	tx, err := p.dbCore.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := fmt.Sprintf(`UPDATE %s SET status=? WHERE id=?`, loanDetailTab)
	if _, err := tx.ExecContext(ctx, q, 1, id); err != nil {
		return err
	}

	q = `INSERT INTO income (loan_general_id,amount,datetime) VALUES (?,?,?)`
	if _, err := tx.ExecContext(ctx, q, lgid, amount, time.Now().Format(time.DateTime)); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *PgRepository) BulkSaveLoanDetail(ctx context.Context, dt []model.LoanDetail) error {
	tx, err := p.dbCore.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `INSERT INTO loan_details (loan_general_id, amount, status, datetime) VALUES (?,?,?,?)`

	bulkString := ""
	args := make([]interface{}, 0)
	args = append(args, dt[0].LoanGeneralID, dt[0].Amount, dt[0].Status, dt[0].Datetime)
	if len(dt) > 1 {
		for i := 1; i < len(dt); i++ {
			args = append(args, dt[i].LoanGeneralID, dt[i].Amount, dt[i].Status, dt[i].Datetime)
			bulkString += ",(?,?,?,?)"
		}
	}
	bulkString += ";"
	q += bulkString

	if _, err := tx.ExecContext(ctx, q, args...); err != nil {
		return err
	}

	q = `UPDATE loan_generals SET status=? WHERE id=?`
	if _, err := tx.ExecContext(ctx, q, 1, dt[0].LoanGeneralID); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *PgRepository) AddBalance(ctx context.Context, id, uid uint32, balance string, c, amount float64) error {
	tx, err := p.dbCore.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := fmt.Sprintf(`UPDATE users SET %s=? WHERE id=?`, balance)
	args := []interface{}{
		c + amount, uid,
	}
	if balance == "social" {
		q = `INSERT INTO social_funds (user_id, fund_type, amount) VALUES (?,?,?)`
		args = []interface{}{
			uid, 0, amount,
		}
	}
	if _, err := tx.ExecContext(ctx, q, args...); err != nil {
		return err
	}

	q = `UPDATE loan_generals SET status=? WHERE id=?`
	if _, err := tx.ExecContext(ctx, q, 1, id); err != nil {
		return err
	}

	if balance == "balance" {
		q = `INSERT INTO income (loan_general_id,amount,datetime) VALUES (?,?,?)`
		if _, err := tx.ExecContext(ctx, q, id, amount, time.Now().Format(time.DateTime)); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
