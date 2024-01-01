package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"time"
)

func (p *PgRepository) IsEmailExists(ctx context.Context, email string) (bool, error) {
	q := `SELECT email FROM users WHERE email = ?;`

	var emailDB string
	err := p.dbCore.QueryRowContext(ctx, q, email).Scan(&emailDB)
	if errors.Is(err, sql.ErrNoRows) {
		return false, model.ErrNotFound
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *PgRepository) StoreUser(ctx context.Context, regData model.RegistrationRequest) error {
	q := `INSERT INTO users (email, name, password, role, is_active, created_at, updated_at) VALUES (?,?,?,?,?,?,?);`
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, &regData.Email, &regData.Name, &regData.Password, &regData.Role, 0,
		now, now); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	q := `SELECT id,name,email,password,role,is_active,is_leader,balance,vbalance,created_at,updated_at FROM users WHERE email=?`
	var (
		id                 uint32
		name, em, password sql.NullString
		role, isa, isl     sql.NullInt32
		balance, vbalance  sql.NullFloat64
		cat, uat           sql.NullTime
	)

	if err := p.dbCore.QueryRowContext(ctx, q, email).Scan(&id, &name, &em, &password, &role, &isa, &isl, &balance,
		&vbalance, &cat, &uat); err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:        id,
		Name:      name.String,
		Email:     em.String,
		Password:  password.String,
		Role:      int(role.Int32),
		IsActive:  int(isa.Int32),
		IsLeader:  int(isl.Int32),
		Balance:   balance.Float64,
		VBalance:  vbalance.Float64,
		CreatedAt: cat.Time.Format(time.RFC3339),
		UpdatedAt: uat.Time.Format(time.RFC3339),
	}, nil
}

func (p *PgRepository) GetUser(ctx context.Context, id uint32) (model.User, error) {
	q := `SELECT id,name,email,password,role,is_active,is_leader,balance,vbalance,created_at,updated_at FROM users WHERE id=?`
	var (
		ids                uint32
		name, em, password sql.NullString
		role, isa, isl     sql.NullInt32
		balance, vbalance  sql.NullFloat64
		cat, uat           sql.NullTime
	)

	if err := p.dbCore.QueryRowContext(ctx, q, id).Scan(&ids, &name, &em, &password, &role, &isa, &isl, &balance,
		&vbalance, &cat, &uat); err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:        ids,
		Name:      name.String,
		Email:     em.String,
		Password:  password.String,
		Role:      int(role.Int32),
		IsActive:  int(isa.Int32),
		IsLeader:  int(isl.Int32),
		Balance:   balance.Float64,
		VBalance:  vbalance.Float64,
		CreatedAt: cat.Time.Format(time.RFC3339),
		UpdatedAt: uat.Time.Format(time.RFC3339),
	}, nil
}

func (p *PgRepository) UpdateIsActive(ctx context.Context, id uint32, isActive int) error {
	q := `UPDATE users SET is_active=?, updated_at=? WHERE id=?;`
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, isActive, now, id); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) GetUsersByFilter(ctx context.Context, filter map[string]string) ([]model.User, error) {
	query := `SELECT id,name,email,role,is_active,is_leader,created_at,updated_at FROM users`
	if len(filter) > 0 {
		query += ` WHERE `
	}

	args := make([]interface{}, 0)
	idx := 1
	for key, val := range filter {
		query += fmt.Sprintf("%v LIKE ?", key)

		if idx != len(filter) {
			query += " AND "
		}
		args = append(args, "%"+val+"%")
		idx += 1
	}

	rows, err := p.dbCore.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userData := make([]model.User, 0)
	for rows.Next() {
		var (
			id             uint32
			name, em       sql.NullString
			role, isa, isl sql.NullInt32
			cat, uat       sql.NullTime
		)
		if err := rows.Scan(&id, &name, &em, &role, &isa, &isl, &cat, &uat); err != nil {
			return nil, err
		}

		userData = append(userData, model.User{
			ID:        id,
			Name:      name.String,
			Email:     em.String,
			Role:      int(role.Int32),
			IsActive:  int(isa.Int32),
			IsLeader:  int(isl.Int32),
			CreatedAt: cat.Time.Format(time.RFC3339),
			UpdatedAt: uat.Time.Format(time.RFC3339),
		})
	}

	return userData, nil
}

func (p *PgRepository) UpdateUser(ctx context.Context, id uint32, usr model.User) error {
	q := fmt.Sprintf(`UPDATE users SET name=?, email=?, role=?, is_active=?, is_leader=?, balance=?, vbalance=?, updated_at=? WHERE id=?;`)
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, usr.Name, usr.Email, usr.Role, usr.IsActive, usr.IsLeader, usr.Balance,
		usr.VBalance, now, id); err != nil {
		return err
	}

	return nil
}
func (p *PgRepository) DeleteUser(ctx context.Context, id uint32) error {
	q := fmt.Sprintf(`DELETE FROM users WHERE id=?;`)

	if _, err := p.dbCore.ExecContext(ctx, q, id); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) UpdateUserPassword(ctx context.Context, id uint32, newPw string) error {
	q := fmt.Sprintf(`UPDATE users SET password=?, updated_at=? WHERE id=?;`)
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, newPw, now, id); err != nil {
		return err
	}
	return nil
}
