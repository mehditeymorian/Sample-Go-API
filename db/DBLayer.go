package db

import "GoServer/model"

type Layer interface {
	init()
	insert(customer model.Customer) model.Customer
	edit(customer model.Customer) model.Customer
	delete() bool
	retrieveAll() []model.Customer
}
