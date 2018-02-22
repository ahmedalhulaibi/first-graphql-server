package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreatePerson(db *gorm.DB, newPerson Person) {
	db.Create(&newPerson)
}

func GetPerson(db *gorm.DB, queryPerson Person, resultPerson *Person) {
	db.Where(&queryPerson).First(resultPerson)
}

func UpdatePerson(db *gorm.DB, oldPerson Person, newPerson Person, resultPerson *Person) {
	var oldResultPerson Person
	db.Where(&oldPerson).First(&oldResultPerson)
	if oldResultPerson.ID == newPerson.ID {
		oldResultPerson = newPerson
		db.Save(oldResultPerson)
	}
	GetPerson(db, newPerson, resultPerson)
}

func DeletePerson(db *gorm.DB, oldPerson Person) {
	db.Delete(&oldPerson)
}

func CreateAntiOrder(db *gorm.DB, newAntiOrder AntiOrder) {
	db.Create(&newAntiOrder)
}

func GetAntiOrder(db *gorm.DB, queryAntiOrder AntiOrder, resultAntiOrder *AntiOrder) {
	db.Where(&queryAntiOrder).First(resultAntiOrder)
}

func UpdateAntiOrder(db *gorm.DB, oldAntiOrder AntiOrder, newAntiOrder AntiOrder, resultAntiOrder *AntiOrder) {
	var oldResultAntiOrder AntiOrder
	db.Where(&oldAntiOrder).First(&oldResultAntiOrder)
	if oldResultAntiOrder.AntiOrderID == newAntiOrder.AntiOrderID {
		oldResultAntiOrder = newAntiOrder
		db.Save(oldResultAntiOrder)
	}
	GetAntiOrder(db, newAntiOrder, resultAntiOrder)
}

func DeleteAntiOrder(db *gorm.DB, oldAntiOrder AntiOrder) {
	db.Delete(&oldAntiOrder)
}

func CreateOrder(db *gorm.DB, newOrder Order) {
	db.Create(&newOrder)
}

func GetOrder(db *gorm.DB, queryOrder Order, resultOrder *Order) {
	db.Where(&queryOrder).First(resultOrder)
}

func UpdateOrder(db *gorm.DB, oldOrder Order, newOrder Order, resultOrder *Order) {
	var oldResultOrder Order
	db.Where(&oldOrder).First(&oldResultOrder)
	if oldResultOrder.OrderID == newOrder.OrderID {
		oldResultOrder = newOrder
		db.Save(oldResultOrder)
	}
	GetOrder(db, newOrder, resultOrder)
}

func DeleteOrder(db *gorm.DB, oldOrder Order) {
	db.Delete(&oldOrder)
}
