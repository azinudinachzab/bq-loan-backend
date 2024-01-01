package model

type LoanType struct {
	ID           uint32  `validate:"-" json:"id"`
	Name         string  `validate:"required,min=3,max=100" json:"name"`
	Margin       float64 `validate:"gte=0" json:"margin"`
	IsAddBalance int     `validate:"required,gt=0,lt=4" json:"is_add_balance"`
	CreatedAt    string  `validate:"-" json:"created_at"`
	UpdatedAt    string  `validate:"-" json:"updated_at"`
}
