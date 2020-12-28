package model

import "time"

type Customer struct {
	Name         string    `json:"c_name"`
	Telephone    int32     `json:"cTel"`
	Address      string    `json:"cAddress"`
	Id           int       `json:"cId"`
	RegisterDate time.Time `json:"cRegisterDate"`
}
