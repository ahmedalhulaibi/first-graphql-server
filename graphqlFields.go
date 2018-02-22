package main

import (
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Fields = graphql.Fields{
	"Person": &graphql.Field{
		Type: personType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			person := Person{}
			DB.First(&person)
			antiOrder := AntiOrder{}
			DB.Model(&person).Association("AntiOrder").Find(&antiOrder)
			person.AntiOrder = antiOrder
			orders := []Order{}
			DB.Model(&person).Association("Orders").Find(&orders)
			person.Orders = append(person.Orders, orders...)
			return person, nil
		},
	},
	"AntiOrder": &graphql.Field{
		Type: antiOrderType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			antiOrder := AntiOrder{}
			DB.First(&antiOrder)
			return antiOrder, nil
		},
	},
	"Order": &graphql.Field{
		Type: orderType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			order := Order{}
			DB.First(&order)
			return order, nil
		},
	},
}
