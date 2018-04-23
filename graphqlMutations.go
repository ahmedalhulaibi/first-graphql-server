package main

import (
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MutationFields graphql.Fields

func init() {
	MutationFields = make(graphql.Fields, 1)

	MutationFields["CreateAntiOrder"] = &graphql.Field{
		Type: antiOrderType,
		Args: graphql.FieldConfigArgument{
			"AntiOrderID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"AntiOrderNumber": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"PersonID": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryAntiOrderObj := AntiOrder{}
			if val, ok := p.Args["AntiOrderID"]; ok {
				QueryAntiOrderObj.AntiOrderID = int32(val.(int))
			}
			if val, ok := p.Args["AntiOrderNumber"]; ok {
				QueryAntiOrderObj.AntiOrderNumber = int32(val.(int))
			}
			if val, ok := p.Args["PersonID"]; ok {
				QueryAntiOrderObj.PersonID = int32(val.(int))
			}

			err := CreateAntiOrder(DB, QueryAntiOrderObj)
			var ResultAntiOrderObj AntiOrder
			err = append(err, GetAntiOrder(DB, QueryAntiOrderObj, &ResultAntiOrderObj)...)
			if len(err) > 0 {
				return ResultAntiOrderObj, err[len(err)-1]
			}
			return ResultAntiOrderObj, nil
		},
	}

	MutationFields["CreateOrder"] = &graphql.Field{
		Type: orderType,
		Args: graphql.FieldConfigArgument{
			"OrderID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"OrderNumber": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"PersonID": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryOrderObj := Order{}
			if val, ok := p.Args["OrderID"]; ok {
				QueryOrderObj.OrderID = int32(val.(int))
			}
			if val, ok := p.Args["OrderNumber"]; ok {
				QueryOrderObj.OrderNumber = int32(val.(int))
			}
			if val, ok := p.Args["PersonID"]; ok {
				QueryOrderObj.PersonID = int32(val.(int))
			}

			err := CreateOrder(DB, QueryOrderObj)
			var ResultOrderObj Order
			err = append(err, GetOrder(DB, QueryOrderObj, &ResultOrderObj)...)
			if len(err) > 0 {
				return ResultOrderObj, err[len(err)-1]
			}
			return ResultOrderObj, nil
		},
	}

	MutationFields["CreatePerson"] = &graphql.Field{
		Type: personType,
		Args: graphql.FieldConfigArgument{
			"Address": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"City": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"FirstName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"ID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"LastName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryPersonObj := Person{}
			if val, ok := p.Args["Address"]; ok {
				QueryPersonObj.Address = val.(string)
			}
			if val, ok := p.Args["City"]; ok {
				QueryPersonObj.City = val.(string)
			}
			if val, ok := p.Args["FirstName"]; ok {
				QueryPersonObj.FirstName = val.(string)
			}
			if val, ok := p.Args["ID"]; ok {
				QueryPersonObj.ID = int32(val.(int))
			}
			if val, ok := p.Args["LastName"]; ok {
				QueryPersonObj.LastName = val.(string)
			}

			err := CreatePerson(DB, QueryPersonObj)
			var ResultPersonObj Person
			err = append(err, GetPerson(DB, QueryPersonObj, &ResultPersonObj)...)
			if len(err) > 0 {
				return ResultPersonObj, err[len(err)-1]
			}
			return ResultPersonObj, nil
		},
	}

	MutationFields["DeleteAntiOrder"] = &graphql.Field{
		Type: antiOrderType,
		Args: graphql.FieldConfigArgument{
			"AntiOrderID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"AntiOrderNumber": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"PersonID": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryAntiOrderObj := AntiOrder{}
			if val, ok := p.Args["AntiOrderID"]; ok {
				QueryAntiOrderObj.AntiOrderID = int32(val.(int))
			}
			if val, ok := p.Args["AntiOrderNumber"]; ok {
				QueryAntiOrderObj.AntiOrderNumber = int32(val.(int))
			}
			if val, ok := p.Args["PersonID"]; ok {
				QueryAntiOrderObj.PersonID = int32(val.(int))
			}

			err := DeleteAntiOrder(DB, QueryAntiOrderObj)
			var ResultAntiOrderObj AntiOrder
			err = append(err, GetAntiOrder(DB, QueryAntiOrderObj, &ResultAntiOrderObj)...)
			if len(err) > 0 {
				return ResultAntiOrderObj, err[len(err)-1]
			}
			return ResultAntiOrderObj, nil
		},
	}

	MutationFields["DeleteOrder"] = &graphql.Field{
		Type: orderType,
		Args: graphql.FieldConfigArgument{
			"OrderID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"OrderNumber": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"PersonID": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryOrderObj := Order{}
			if val, ok := p.Args["OrderID"]; ok {
				QueryOrderObj.OrderID = int32(val.(int))
			}
			if val, ok := p.Args["OrderNumber"]; ok {
				QueryOrderObj.OrderNumber = int32(val.(int))
			}
			if val, ok := p.Args["PersonID"]; ok {
				QueryOrderObj.PersonID = int32(val.(int))
			}

			err := DeleteOrder(DB, QueryOrderObj)
			var ResultOrderObj Order
			err = append(err, GetOrder(DB, QueryOrderObj, &ResultOrderObj)...)
			if len(err) > 0 {
				return ResultOrderObj, err[len(err)-1]
			}
			return ResultOrderObj, nil
		},
	}

	MutationFields["DeletePerson"] = &graphql.Field{
		Type: personType,
		Args: graphql.FieldConfigArgument{
			"Address": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"City": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"FirstName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"ID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"LastName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			QueryPersonObj := Person{}
			if val, ok := p.Args["Address"]; ok {
				QueryPersonObj.Address = val.(string)
			}
			if val, ok := p.Args["City"]; ok {
				QueryPersonObj.City = val.(string)
			}
			if val, ok := p.Args["FirstName"]; ok {
				QueryPersonObj.FirstName = val.(string)
			}
			if val, ok := p.Args["ID"]; ok {
				QueryPersonObj.ID = int32(val.(int))
			}
			if val, ok := p.Args["LastName"]; ok {
				QueryPersonObj.LastName = val.(string)
			}

			err := DeletePerson(DB, QueryPersonObj)
			var ResultPersonObj Person
			err = append(err, GetPerson(DB, QueryPersonObj, &ResultPersonObj)...)
			if len(err) > 0 {
				return ResultPersonObj, err[len(err)-1]
			}
			return ResultPersonObj, nil
		},
	}

}
