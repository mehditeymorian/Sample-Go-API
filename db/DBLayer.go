package db

import "GoServer/model"

var Database = CreateListDB()

type Layer interface {
	Init()
	Insert(customer *model.Customer) model.Customer
	Edit(customerId int64, customer *model.Customer) (model.Customer, error)
	Delete(customerId int64) error
	RetrieveAll() ([]model.Customer, error)
}
