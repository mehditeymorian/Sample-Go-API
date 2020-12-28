package db

import (
	"GoServer/model"
	"errors"
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

	for i, each := range db.customers {
		if each.Id == customerId {
			db.customers[i].Name = customer.Name
			db.customers[i].Address = customer.Address
			db.customers[i].Telephone = customer.Telephone
			return db.customers[i], nil
		}
	}

	return model.Customer{}, errors.New("customer not found")
}

func (db *ListDB) Delete(customerId int64) error {

}

func (db *ListDB) RetrieveAll() ([]model.Customer, error) {
	if len(db.customers) == 0 {
		return nil, errors.New("no customer is stored.")
	}
	return db.customers, nil
}
