package main

import (
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var antiOrderType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AntiOrder",
		Fields: graphql.Fields{

			"AntiOrderID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"AntiOrderNumber": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"PersonID": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var orderType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{

			"OrderID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"OrderNumber": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"PersonID": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var personType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{

			"Orders": &graphql.Field{
				Type: graphql.NewList(orderType),
			},
			"ID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"LastName": &graphql.Field{
				Type: graphql.String,
			},
			"FirstName": &graphql.Field{
				Type: graphql.String,
			},
			"Address": &graphql.Field{
				Type: graphql.String,
			},
			"City": &graphql.Field{
				Type: graphql.String,
			},
			"AntiOrder": &graphql.Field{
				Type: antiOrderType,
			},
		},
	},
)
