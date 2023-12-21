package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)

var oddsListFromAPI []odds
var apiKey string

var queryAPIOddType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single upcoming odd by id and sport_key from API
			   http://localhost:8080/odds?query={oddfromapi(id:%220dfadf729863f66fd9b4f1d25e44be7d%22,sport_key:%22icehockey_sweden_hockey_league%22){id,sport_key,sport_title,commence_time,home_team,away_team,bookmakers{key,title,last_update,markets{key,last_update,outcomes{name,price}}}}}
				id and sport_key should be valid and this event cannot be expired
				{id,sport_key,sport_title,commence_time,home_team,away_team,bookmakers{key,title,last_update,markets{key,last_update,outcomes{name,price}}}
				are all optional
			*/
			"oddfromapi": &graphql.Field{
				Type:        oddsType,
				Description: "Get upcoming odd by id from API",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"sport_key": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if !ok {
						return nil, nil
					}
					sportkey, ok := p.Args["sport_key"].(string)
					if !ok {
						return nil, nil
					}
					url := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%v/events/%v/odds?apiKey=%v&regions=us&oddsFormat=american", sportkey, id, apiKey)
					result := FetchOdd(url)
					return result, nil
				},
			},
			/* Get (read) odds list
			   http://localhost:8080/odds?query={odds{id,sport_key,sport_title,commence_time,home_team,away_team}}
				{id,sport_key,sport_title,commence_time,home_team,away_team} are all optional
			*/
			"odds": &graphql.Field{
				Type:        graphql.NewList(oddsType),
				Description: "Get odds list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					url := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/upcoming/odds/?apiKey=%v&regions=us&markets=h2h&bookmakers=draftkings", apiKey)
					result := FetchOdds(url)
					return result, nil
				},
			},
		},
	})

var schemaForAPI, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryAPIOddType,
	})

func executeAPIQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schemaForAPI,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func FetchDataFromAPI() {
	apiKey = GetApiKey()

	http.HandleFunc("/odds", func(writer http.ResponseWriter, request *http.Request) {
		result := executeAPIQuery(request.URL.Query().Get("query"), schemaForAPI)
		json.NewEncoder(writer).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
