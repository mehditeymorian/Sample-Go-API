package db

import (
	"GoServer/model"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// use to configure the connection between API and database server
var host = "localhost"
var user = "postgres"
var pass = "admin"
var dbname = "customers"
var port = "5432"

type PostgresDB struct {
	connection *gorm.DB
}

func CreatePostgresDB() *PostgresDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return &PostgresDB{
		connection: db,
	}
}

func (db *PostgresDB) Init() {
	_ = db.connection.AutoMigrate(&model.Customer{})
}

func (db *PostgresDB) Insert(customer *model.Customer) model.Customer {

	insertVal := model.Customer{
		Name:         customer.Name,
		Telephone:    customer.Telephone,
		Address:      customer.Address,
		RegisterDate: time.Now().Format("2006-01-02"),
	}
	err := db.connection.Create(&insertVal).Error
	fmt.Print(err)
	return insertVal
}

func (db *PostgresDB) Edit(customerId int64, customer *model.Customer) (model.Customer, error) {

	edited := model.Customer{
		ID: customerId,
	}
	result := db.connection.First(&edited)

	if result.RowsAffected == 0 {
		return model.Customer{}, errors.New("customer not found")
	}

	edited.Name = customer.Name
	edited.Address = customer.Address
	edited.Telephone = customer.Telephone

	db.connection.Save(&edited)

	return edited, nil
}

func (db *PostgresDB) Delete(customerId int64) error {

	deleted := model.Customer{
		ID: customerId,
	}
	result := db.connection.Delete(&deleted)

	if result.RowsAffected == 0 {
		return errors.New("customer not found")
	}

	return nil
}

func (db *PostgresDB) RetrieveAll() ([]model.Customer, error) {
	var all []model.Customer

	db.connection.Find(&all)

	if len(all) == 0 {
		return nil, errors.New("no customer is stored.")
	}
	return all, nil
}
