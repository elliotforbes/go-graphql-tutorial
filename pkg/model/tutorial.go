package model

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Tutorial struct {
	gorm.Model
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	},
)

func init() {
	db, err := gorm.Open("sqlite3", "tutorials.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&Tutorial{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Author{})
}

func SingleTutorialSchema() *graphql.Field {
	return &graphql.Field{
		Type:        tutorialType,
		Description: "Get Tutorial By ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var tutorial Tutorial
			db, _ := gorm.Open("sqlite3", "tutorials.db")
			db.First(&tutorial, params.Args["id"].(int))
			return tutorial, nil
		},
	}
}

func ListTutorialSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(tutorialType),
		Description: "Get Tutorial List",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var tutorials []Tutorial
			db, _ := gorm.Open("sqlite3", "tutorials.db")
			db.Find(&tutorials)
			return tutorials, nil
		},
	}
}

func CreateTutorialMutation() *graphql.Field {
	return &graphql.Field{
		Type:        tutorialType,
		Description: "Create a new Tutorial",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			tutorial := Tutorial{ID: params.Args["id"].(int), Title: params.Args["title"].(string)}
			db, _ := gorm.Open("sqlite3", "tutorials.db")
			db.Save(&tutorial)
			return tutorial, nil
		},
	}
}
