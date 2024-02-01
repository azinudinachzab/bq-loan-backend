package model

type SocialFund struct {
	ID        uint32      `validate:"-" json:"id"`
	Title     string      `validate:"required" json:"title"`
	UserID    uint32      `validate:"required" json:"user_id"`
	Users     UserGeneral `validate:"-" json:"user"`
	Amount    float64     `validate:"required,gt=0" json:"amount"`
	Status    int         `validate:"-" json:"status"`
	FundType  int         `validate:"required,gte=0,lt=2" json:"fund_type"`
	CreatedAt string      `validate:"-" json:"created_at"`
	UpdatedAt string      `validate:"-" json:"updated_at"`
}
