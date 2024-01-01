package model

type (
	User struct {
		ID        uint32  `json:"id"`
		Name      string  `json:"name"`
		Email     string  `json:"email"`
		Password  string  `json:"password"`
		Role      int     `json:"role"`
		IsActive  int     `json:"is_active"`
		IsLeader  int     `json:"is_leader"`
		Balance   float64 `json:"balance"`
		VBalance  float64 `json:"vbalance"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
	}

	RegistrationRequest struct {
		Email    string `validate:"required,email" json:"email"`
		Name     string `validate:"required" json:"name"`
		Password string `validate:"required" json:"password"`
		Role     int    `validate:"omitempty,gt=0,lt=5" json:"role"`
	}

	LoginRequest struct {
		Email    string `validate:"required,email" json:"email"`
		Password string `validate:"required" json:"password"`
	}
)
