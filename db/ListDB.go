package db

import (
	"GoServer/model"
	"time"
)

type ListDB struct {
	customers []model.Customer
	currentId int64
}

func CreateListDB() Layer {
	return &ListDB{
		customers: []model.Customer{},
		currentId: 0,
	}
}

func (db *ListDB) Init() {

}

func (db *ListDB) Insert(customer *model.Customer) model.Customer {

	insertVal := model.Customer{
		Name:         customer.Name,
		Telephone:    customer.Telephone,
		Address:      customer.Address,
		Id:           db.currentId,
		RegisterDate: time.Now().Format("2006-01-02"),
	}
	db.customers = append(db.customers, insertVal)
	db.currentId = db.currentId + 1

	return insertVal
}

func (db *ListDB) Edit(customerId int64, customer *model.Customer) (model.Customer, error) {
	return model.Customer{}, nil
}

func (db *ListDB) Delete() bool {
	return false
}

func (db *ListDB) RetrieveAll() []model.Customer {
	return db.customers
}
