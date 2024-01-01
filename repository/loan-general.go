package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/azinudinachzab/bq-loan-be-v2/model"
)

const (
	loanGeneralTab = "loan_generals"
)

func (p *PgRepository) CreateLoanGeneral(ctx context.Context, loanGeneral model.LoanGeneral) error {
	q := fmt.Sprintf(`INSERT INTO %s (title,amount,datetime,tenor,status,user_id,loan_type_id,
                created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?);`, loanGeneralTab)
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, loanGeneral.Title, loanGeneral.Amount, loanGeneral.Datetime,
		loanGeneral.Tenor, loanGeneral.Status, loanGeneral.UserID, loanGeneral.LoanTypeID, now, now); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) GetLoanGenerals(ctx context.Context, lastID, uID uint32, t string) ([]model.LoanGeneral, error) {
	qWhere := ""
	args := make([]interface{}, 0)
	args = append(args, lastID)

	if uID != 0 {
		qWhere += " AND lg.user_id=?"
		args = append(args, uID)
	}

	if t != "" {
		qWhere += " AND lg.title LIKE ?"
		args = append(args, "%"+t+"%")
	}

	//if c == true {
	//	now := time.Now()
	//	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	//	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-1 * time.Second)
	//	qWhere += " AND datetime >=? AND datetime <=?"
	//	args = append(args, startOfMonth)
	//	args = append(args, endOfMonth)
	//}

	q := fmt.Sprintf(`SELECT lg.id,lg.title,lg.amount,lg.datetime,lg.tenor,lg.status,lg.user_id,lg.loan_type_id,
       	lg.created_at,lg.updated_at,lt.name as ltname,u.name as uname FROM %s lg
		INNER JOIN loan_types lt on lg.loan_type_id = lt.id INNER JOIN users u on lg.user_id = u.id
		WHERE lg.id > ?%s ORDER BY lg.id LIMIT 25`, loanGeneralTab, qWhere)
	rows, err := p.dbCore.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rowClose(rows)

	loanGenerals := make([]model.LoanGeneral, 0)
	for rows.Next() {
		var (
			ids, ltid, uid       uint32
			title, ltname, uname sql.NullString
			amount               sql.NullFloat64
			tenor, status        sql.NullInt32
			dt, cat, uat         sql.NullTime
		)

		if err := rows.Scan(&ids, &title, &amount, &dt, &tenor, &status, &uid, &ltid,
			&cat, &uat, &ltname, &uname); err != nil {
			return nil, err
		}

		loanGenerals = append(loanGenerals, model.LoanGeneral{
			ID:         ids,
			UserID:     uid,
			Title:      title.String,
			Amount:     amount.Float64,
			Datetime:   dt.Time.Format(time.DateTime),
			Tenor:      int(tenor.Int32),
			Status:     int(status.Int32),
			LoanTypeID: ltid,
			CreatedAt:  cat.Time.Format(time.RFC3339),
			UpdatedAt:  uat.Time.Format(time.RFC3339),
			LoanTypes:  model.LoanTypeGeneral{ID: ltid, Name: ltname.String},
			Users:      model.UserGeneral{ID: uid, Name: uname.String},
		})
		if rows.Err() != nil {
			return nil, err
		}
	}

	return loanGenerals, nil
}
func (p *PgRepository) GetLoanGeneral(ctx context.Context, id uint32) (model.LoanGeneral, error) {
	q := fmt.Sprintf(`SELECT id,title,amount,datetime,tenor,status,user_id,loan_type_id,
                created_at,updated_at FROM %s WHERE id=?`, loanGeneralTab)
	var (
		ids, ltid, uid uint32
		title          sql.NullString
		amount         sql.NullFloat64
		tenor, status  sql.NullInt32
		dt, cat, uat   sql.NullTime
	)
	if err := p.dbCore.QueryRowContext(ctx, q, id).Scan(&ids, &title, &amount, &dt, &tenor, &status, &uid, &ltid,
		&cat, &uat); err != nil {
		return model.LoanGeneral{}, err
	}
	return model.LoanGeneral{
		ID:         ids,
		UserID:     uid,
		Title:      title.String,
		Amount:     amount.Float64,
		Datetime:   dt.Time.Format(time.DateTime),
		Tenor:      int(tenor.Int32),
		Status:     int(status.Int32),
		LoanTypeID: ltid,
		CreatedAt:  cat.Time.Format(time.RFC3339),
		UpdatedAt:  uat.Time.Format(time.RFC3339),
	}, nil
}

func (p *PgRepository) UpdateLoanGeneral(ctx context.Context, lt model.LoanGeneral) error {
	q := fmt.Sprintf(`UPDATE %s SET title=?, amount=?, datetime=?, tenor=?, updated_at=? WHERE id=?;`, loanGeneralTab)
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, lt.Title, lt.Amount, lt.Datetime, lt.Tenor, now, lt.ID); err != nil {
		return err
	}

	return nil
}
func (p *PgRepository) DeleteLoanGeneral(ctx context.Context, id uint32) error {
	q := fmt.Sprintf(`DELETE FROM %s WHERE id=?;`, loanGeneralTab)

	if _, err := p.dbCore.ExecContext(ctx, q, id); err != nil {
		return err
	}

	return nil
}
