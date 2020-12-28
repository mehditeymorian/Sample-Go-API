package model

import (
	"encoding/json"
	"time"
)

type Customer struct {
	Name         string      `json:"c_name"`
	Telephone    json.Number `json:"cTel"`
	Address      string      `json:"cAddress"`
	Id           json.Number `json:"cId"`
	RegisterDate time.Time   `json:"cRegisterDate"`
}
