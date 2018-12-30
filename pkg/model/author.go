package model

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name      string
	Tutorials []int
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
