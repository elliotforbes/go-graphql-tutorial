package model

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Body string
}

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
