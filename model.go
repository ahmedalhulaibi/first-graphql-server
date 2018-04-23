package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AntiOrder struct {
	AntiOrderID     int32 `gorm:"column:AntiOrderID;primary_key;" json:"AntiOrderID" `
	AntiOrderNumber int32 `gorm:"column:AntiOrderNumber;" json:"AntiOrderNumber" `
	PersonID        int32 `gorm:"column:PersonID;" json:"PersonID" `
}

type Order struct {
	OrderID     int32 `gorm:"column:OrderID;primary_key;" json:"OrderID" `
	OrderNumber int32 `gorm:"column:OrderNumber;" json:"OrderNumber" `
	PersonID    int32 `gorm:"column:PersonID;" json:"PersonID" `
}

type Person struct {
	Address   string    `gorm:"column:Address;" json:"Address" `
	AntiOrder AntiOrder `gorm:"ForeignKey:PersonID;AssociationForeignKey:ID;" json:"AntiOrder" `
	City      string    `gorm:"column:City;" json:"City" `
	FirstName string    `gorm:"column:FirstName;" json:"FirstName" `
	ID        int32     `gorm:"column:ID;primary_key;" json:"ID" `
	LastName  string    `gorm:"column:LastName;" json:"LastName" `
	Orders    []Order   `gorm:"ForeignKey:PersonID;AssociationForeignKey:ID;" json:"Orders" `
}

func (AntiOrder) TableName() string {
	return "AntiOrders"
}

func (Order) TableName() string {
	return "Orders"
}

func (Person) TableName() string {
	return "Persons"
}
