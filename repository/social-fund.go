package repository

import (
	"context"
	"fmt"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"time"
)

const (
	socialFundsTab = "social_funds"
)

func (p *PgRepository) CreateSocialFundRequest(ctx context.Context, sf model.SocialFund) error {
	q := fmt.Sprintf(`INSERT INTO %s (user_id,title,fund_type,amount,status,created_at,updated_at) VALUES (?,?,?,?,?,?,?);`, socialFundsTab)
	now := time.Now().Format(time.RFC3339)

	if _, err := p.dbCore.ExecContext(ctx, q, sf.UserID, sf.Title, 1, sf.Amount, 0, now, now); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) UpdateSocialFundStatus(ctx context.Context, id uint32) error {
	q := fmt.Sprintf(`UPDATE %s SET status=? WHERE id=?;`, socialFundsTab)
	if _, err := p.dbCore.ExecContext(ctx, q, 1, id); err != nil {
		return err
	}

	return nil
}
