package model

type Customer struct {
	Name         string `json:"cName"`
	Telephone    int64  `json:"cTel"`
	Address      string `json:"cAddress"`
	Id           int64  `json:"cID"`
	RegisterDate string `json:"cRegisterDate"`
}

type CustomerInsertResp struct {
	Name         string `json:"cName"`
	Telephone    int64  `json:"cTel"`
	Address      string `json:"cAddress"`
	Id           int64  `json:"cID"`
	RegisterDate string `json:"cRegisterDate"`
	Msg          string `json:"msg"`
}

type CustomerRetrieveResp struct {
	Size      int64      `json:"size"`
	Customers []Customer `json:"customers"`
	Msg       string     `json:"msg"`
}
