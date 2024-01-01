package model

type LoanGeneral struct {
	ID         uint32          `validate:"-" json:"id"`
	UserID     uint32          `validate:"required" json:"user_id"`
	Users      UserGeneral     `validate:"-" json:"user"`
	Title      string          `validate:"required,min=1,max=100" json:"title"`
	Amount     float64         `validate:"required,gt=0" json:"amount"`
	Datetime   string          `validate:"required" json:"datetime"`
	Tenor      int             `validate:"required,gt=0" json:"tenor"`
	Status     int             `validate:"-" json:"status"`
	LoanTypeID uint32          `validate:"required" json:"loan_type_id"`
	LoanTypes  LoanTypeGeneral `validate:"-" json:"loan_type"`
	CreatedAt  string          `validate:"-" json:"created_at"`
	UpdatedAt  string          `validate:"-" json:"updated_at"`
}

type LoanTypeGeneral struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

type UserGeneral struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}
