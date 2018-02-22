package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Order struct {
	OrderID     int32 `gorm:"column:OrderID;primary_key;" json:"OrderID" `
	OrderNumber int32 `gorm:"column:OrderNumber;" json:"OrderNumber" `
	PersonID    int32 `gorm:"column:PersonID;" json:"PersonID" `
}

func (Order) TableName() string {
	return "Orders"
}

type Person struct {
	AntiOrder AntiOrder `gorm:"ForeignKey:PersonID;AssociationForeignKey:ID;" json:"AntiOrder" `
	Orders    []Order   `gorm:"ForeignKey:PersonID;AssociationForeignKey:ID;" json:"Orders" `
	ID        int32     `gorm:"column:ID;primary_key;" json:"ID" `
	LastName  string    `gorm:"column:LastName;" json:"LastName" `
	FirstName string    `gorm:"column:FirstName;" json:"FirstName" `
	Address   string    `json:"Address" gorm:"column:Address;" `
	City      string    `gorm:"column:City;" json:"City" `
}

func (Person) TableName() string {
	return "Persons"
}

type AntiOrder struct {
	AntiOrderID     int32 `gorm:"column:AntiOrderID;primary_key;" json:"AntiOrderID" `
	AntiOrderNumber int32 `gorm:"column:AntiOrderNumber;" json:"AntiOrderNumber" `
	PersonID        int32 `gorm:"column:PersonID;" json:"PersonID" `
}

func (AntiOrder) TableName() string {
	return "AntiOrders"
}
