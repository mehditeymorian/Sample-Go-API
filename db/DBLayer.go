package db

import "GoServer/model"

//var Database = CreatePostgresDB()
var Database = CreateListDB()

type Layer interface {
	// for initializing anything after database created
	Init()
	// insert a customer to database and return the inserted version
	Insert(customer *model.Customer) model.Customer
	// edit customer with id = <customerId> and use field of <customer> to update the fields
	Edit(customerId int64, customer *model.Customer) (model.Customer, error)
	// delete customer with id = <customerId>
	Delete(customerId int64) error
	// return all customers
	RetrieveAll() ([]model.Customer, error)
}
