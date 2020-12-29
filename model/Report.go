package model

type MonthlyReport struct {
	TotalCustomers int64  `json:"totalCustomers"`
	Period         int64  `json:"period"`
	Msg            string `json:"msg"`
}

type AnnualReport struct {
	Months []MonthlyReport `json:"months"`
	Msg    string          `json:"msg"`
}
