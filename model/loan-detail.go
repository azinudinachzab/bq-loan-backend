package model

type LoanDetail struct {
	ID            uint32        `validate:"required" json:"id"`
	LoanGeneralID uint32        `validate:"required" json:"loan_general_id"`
	General       GeneralDetail `validate:"-" json:"loan_general"`
	Amount        float64       `validate:"required" json:"amount"`
	Datetime      string        `validate:"-" json:"datetime"`
	Status        int           `validate:"-" json:"status"`
	CreatedAt     string        `validate:"-" json:"created_at"`
	UpdatedAt     string        `validate:"-" json:"updated_at"`
}

type GeneralDetail struct {
	Title    string  `json:"title,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Datetime string  `json:"datetime,omitempty"`
	Tenor    int     `json:"tenor,omitempty"`
	Status   int     `json:"status,omitempty"`
}
