package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateAntiOrder(db *gorm.DB, newAntiOrder AntiOrder) []error {
	return db.Create(&newAntiOrder).GetErrors()
}

func GetAntiOrder(db *gorm.DB, queryAntiOrder AntiOrder, resultAntiOrder *AntiOrder) []error {
	return db.Where(&queryAntiOrder).First(resultAntiOrder).GetErrors()
}

func GetAllAntiOrder(db *gorm.DB, queryAntiOrder AntiOrder, resultAntiOrder []AntiOrder) []error {
	return db.Where(&queryAntiOrder).Find(resultAntiOrder).GetErrors()
}

func UpdateAntiOrder(db *gorm.DB, oldAntiOrder AntiOrder, newAntiOrder AntiOrder, resultAntiOrder *AntiOrder) []error {
	var oldResultAntiOrder AntiOrder
	err := db.Where(&oldAntiOrder).First(&oldResultAntiOrder).GetErrors()
	if oldResultAntiOrder.AntiOrderID == newAntiOrder.AntiOrderID {
		oldResultAntiOrder = newAntiOrder
		err = append(err, db.Save(oldResultAntiOrder).GetErrors()...)
	}
	err = append(err, GetAntiOrder(db, newAntiOrder, resultAntiOrder)...)
	return err
}

func DeleteAntiOrder(db *gorm.DB, oldAntiOrder AntiOrder) []error {
	return db.Delete(&oldAntiOrder).GetErrors()
}

func CreateOrder(db *gorm.DB, newOrder Order) []error {
	return db.Create(&newOrder).GetErrors()
}

func GetOrder(db *gorm.DB, queryOrder Order, resultOrder *Order) []error {
	return db.Where(&queryOrder).First(resultOrder).GetErrors()
}

func GetAllOrder(db *gorm.DB, queryOrder Order, resultOrder []Order) []error {
	return db.Where(&queryOrder).Find(resultOrder).GetErrors()
}

func UpdateOrder(db *gorm.DB, oldOrder Order, newOrder Order, resultOrder *Order) []error {
	var oldResultOrder Order
	err := db.Where(&oldOrder).First(&oldResultOrder).GetErrors()
	if oldResultOrder.OrderID == newOrder.OrderID {
		oldResultOrder = newOrder
		err = append(err, db.Save(oldResultOrder).GetErrors()...)
	}
	err = append(err, GetOrder(db, newOrder, resultOrder)...)
	return err
}

func DeleteOrder(db *gorm.DB, oldOrder Order) []error {
	return db.Delete(&oldOrder).GetErrors()
}

func CreatePerson(db *gorm.DB, newPerson Person) []error {
	return db.Create(&newPerson).GetErrors()
}

func GetPerson(db *gorm.DB, queryPerson Person, resultPerson *Person) []error {
	return db.Where(&queryPerson).First(resultPerson).GetErrors()
}

func GetAllPerson(db *gorm.DB, queryPerson Person, resultPerson []Person) []error {
	return db.Where(&queryPerson).Find(resultPerson).GetErrors()
}

func UpdatePerson(db *gorm.DB, oldPerson Person, newPerson Person, resultPerson *Person) []error {
	var oldResultPerson Person
	err := db.Where(&oldPerson).First(&oldResultPerson).GetErrors()
	if oldResultPerson.ID == newPerson.ID {
		oldResultPerson = newPerson
		err = append(err, db.Save(oldResultPerson).GetErrors()...)
	}
	err = append(err, GetPerson(db, newPerson, resultPerson)...)
	return err
}

func DeletePerson(db *gorm.DB, oldPerson Person) []error {
	return db.Delete(&oldPerson).GetErrors()
}
