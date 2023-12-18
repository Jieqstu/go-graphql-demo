package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)

var oddslist []odds

var queryOddType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single odd by id
			http://localhost:8080/odds?query={odd(id:%2256468ff158e85d48b0d04faa35a86461%22){id,sport_key,sport_title,commence_time,home_team,away_team,bookmakers{key,title,last_update,markets{key,last_update,outcomes{name,price}}}}}
			*/
			"odd": &graphql.Field{
				Type:        oddsType,
				Description: "Get odd by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						for _, odd := range oddslist {
							if odd.ID == id {
								return odd, nil
							}
						}
					}
					return nil, nil
				},
			},
			/* Get (read) odds list
			   http://localhost:8080/odds?query={odds{id,sport_key,sport_title,commence_time,home_team,away_team}}
			*/
			"odds": &graphql.Field{
				Type:        graphql.NewList(oddsType),
				Description: "Get odds list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return oddslist, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryOddType,
	})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func ReadList() {
	_ = importJSONFromFile("list.json", &oddslist)

	http.HandleFunc("/odds", func(writer http.ResponseWriter, request *http.Request) {
		result := executeQuery(request.URL.Query().Get("query"), schema)
		json.NewEncoder(writer).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
