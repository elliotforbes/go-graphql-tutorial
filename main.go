package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/elliotforbes/go-graphql-tutorial/pkg/model"
	"github.com/graphql-go/graphql"
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

var aggregateSchema = graphql.Fields{
	"tutorial": model.SingleTutorialSchema(),
	"list":     model.ListTutorialSchema(),
}

var aggregateMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": model.CreateTutorialMutation(),
	},
})

func main() {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: aggregateSchema}
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphql.NewObject(rootQuery),
			Mutation: aggregateMutations,
		},
	)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	query := `
		mutation {
			create(id: 3, title: "Sweet") {
				id
				title
			}
		}
	`
	result := executeQuery(query, schema)
	rJSON, _ := json.Marshal(result)
	fmt.Printf("%s \n", rJSON)

	query = `
		{
			list {
				id
				title
			}
		}
	`
	result = executeQuery(query, schema)
	rJSON, _ = json.Marshal(result)
	fmt.Printf("%s \n", rJSON)
}
