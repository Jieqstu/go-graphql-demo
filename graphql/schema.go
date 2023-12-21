package main

import (
	"github.com/graphql-go/graphql"
	"time"
)

type outcome struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type market struct {
	Key        string    `json:"key"`
	LastUpdate time.Time `json:"last_update"`
	OutComes   []outcome `json:"outcomes"`
}

type bookmaker struct {
	Key        string    `json:"key"`
	Title      string    `json:"title"`
	LastUpdate time.Time `json:"last_update"`
	Markets    []market  `json:"markets"`
}

type odds struct {
	ID           string      `json:"id"`
	SportKey     string      `json:"sport_key"`
	SportTitle   string      `json:"sport_title"`
	CommenceTime time.Time   `json:"commence_time"`
	HomeTeam     string      `json:"home_team"`
	AwayTeam     string      `json:"away_team"`
	Bookmakers   []bookmaker `json:"bookmakers"`
}

var outcomeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Outcome",
		Fields: graphql.Fields{
			"name":  &graphql.Field{Type: graphql.String},
			"price": &graphql.Field{Type: graphql.Float},
		},
	},
)

var marketType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Market",
		Fields: graphql.Fields{
			"key":         &graphql.Field{Type: graphql.String},
			"last_update": &graphql.Field{Type: graphql.DateTime},
			"outcomes":    &graphql.Field{Type: graphql.NewList(outcomeType)},
		},
	})

var bookmakerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Bookmaker",
		Fields: graphql.Fields{
			"key":         &graphql.Field{Type: graphql.String},
			"title":       &graphql.Field{Type: graphql.String},
			"last_update": &graphql.Field{Type: graphql.DateTime},
			"markets":     &graphql.Field{Type: graphql.NewList(marketType)},
		},
	},
)

var oddsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Odds",
		Fields: graphql.Fields{
			"id":            &graphql.Field{Type: graphql.String},
			"sport_key":     &graphql.Field{Type: graphql.String},
			"sport_title":   &graphql.Field{Type: graphql.String},
			"commence_time": &graphql.Field{Type: graphql.DateTime},
			"home_team":     &graphql.Field{Type: graphql.String},
			"away_team":     &graphql.Field{Type: graphql.String},
			"bookmakers":    &graphql.Field{Type: graphql.NewList(bookmakerType)},
		},
	},
)
