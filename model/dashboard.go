package model

type Dashboard struct {
	LatestLoan   float64 `json:"latest_loan"`
	BiggestLoan  float64 `json:"biggest_loan"`
	AcceptedLoan int     `json:"accepted_loan"`
}

type DashboardAdmin struct {
	TotalIncome string `json:"total_income"`
}
