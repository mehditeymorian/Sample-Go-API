package db

import (
	"GoServer/model"
	"encoding/json"
	"time"
)

type ListDB struct {
	customers []model.Customer
	currentId int
}

func CreateListDB() ListDB {
	return ListDB{
		customers: nil,
		currentId: 0,
	}
}

func (db ListDB) Init() {

}

func (db ListDB) Insert(customer *model.Customer) model.Customer {
	currentTime := time.Now()
	insertVal := model.Customer{
		Name:         customer.Name,
		Telephone:    customer.Telephone,
		Address:      customer.Address,
		Id:           json.Number(db.currentId),
		RegisterDate: currentTime,
	}
	db.customers = append(db.customers, insertVal)
	db.currentId++
	return insertVal
}

func (db ListDB) Edit(customer model.Customer) model.Customer {
	return model.Customer{}
}

func (db ListDB) Delete() bool {
	return false
}

func (db ListDB) RetrieveAll() []model.Customer {
	return nil
}
