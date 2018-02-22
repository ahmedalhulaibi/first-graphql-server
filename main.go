package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

var DB *gorm.DB

func main() {

	DB, _ = gorm.Open("mysql", "ahmed:password@tcp(localhost:3306)/delivery")
	defer DB.Close()

	fmt.Println("Test with Get	: curl -g 'http://localhost:8080/graphql?query={Person{ID,LastName,FirstName,Address,City,},AntiOrder{AntiOrderID,AntiOrderNumber,PersonID,},Order{OrderNumber,PersonID,OrderID,},}'")
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: Fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
