package model

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name      string
	Tutorials []Tutorial
}

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

func SetupAuthorMutations() graphql.Fields {
	authorMutationType := graphql.Fields{
		"create": &graphql.Field{
			Type:        authorType,
			Description: "Create a new author",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				author := Author{Name: params.Args["name"].(string)}
				db, _ := gorm.Open("sqlite3", "authors.db")
				db.Save(&author)
				return author, nil
			},
		},
	}
	return authorMutationType
}
