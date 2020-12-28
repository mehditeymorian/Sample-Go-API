package responses

import (
	"GoServer/model"
)

func EchoResp(c model.Customer, msg string) model.CustomerInsertResp {
	return model.CustomerInsertResp{
		Name:         c.Name,
		Telephone:    c.Telephone,
		Address:      c.Address,
		Id:           c.Id,
		RegisterDate: c.RegisterDate,
		Msg:          msg,
	}
}

func RetrieveAllResp(c []model.Customer, msg string) model.CustomerRetrieveResp {
	return model.CustomerRetrieveResp{
		Size:      int64(len(c)),
		Customers: c,
		Msg:       msg,
	}
}

func MsgResp(msg string) model.Msg {
	return model.Msg{
		Msg: msg,
	}
}
