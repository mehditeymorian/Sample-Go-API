package db

import "GoServer/model"

var Database = CreateListDB()

type Layer interface {
	Init()
	Insert(customer *model.Customer) (model.Customer, error)
	Edit(customerId int64, customer *model.Customer) (model.Customer, error)
	Delete() bool
	RetrieveAll() []model.Customer
}
