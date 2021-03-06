package model

type MonthlyReport struct {
	TotalCustomers int64  `json:"totalCustomers"`
	Period         int64  `json:"period"`
	Msg            string `json:"msg"`
}

type MonthInAnnualReport struct {
	TotalCustomers int64 `json:"totalCustomers"`
	Period         int64 `json:"period"`
}

type AnnualReport struct {
	Months [12]MonthInAnnualReport `json:"months"`
	Msg    string                  `json:"msg"`
}
