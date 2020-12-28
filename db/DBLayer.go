package db

import "GoServer/model"

var Database = CreateListDB()

type Layer interface {
	Init()
	Insert(customer *model.Customer) model.Customer
	Edit(customerId int64) model.Customer
	Delete() bool
	RetrieveAll() []model.Customer
}
