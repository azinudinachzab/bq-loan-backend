package model

type Income struct {
	ID            uint32      `validate:"-" json:"id"`
	LoanGeneralID uint32      `validate:"-" json:"loan_general_id"`
	General       LoanGeneral `validate:"-" json:"general"`
	Amount        float64     `validate:"-" json:"amount"`
	Datetime      string      `validate:"-" json:"datetime"`
	CreatedAt     string      `validate:"-" json:"created_at"`
	UpdatedAt     string      `validate:"-" json:"updated_at"`
}
