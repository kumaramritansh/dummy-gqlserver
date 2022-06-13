package main

import (
	"fmt"
	"gitHub.com/apigee/apigee-gqlserver/gqlserver/schema"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status: OK")
	})

	//gqlSchema := graphql.MustParseSchema(, &resolver.Resolver{})
	//http.Handle("/graphql", &relay.Handler{Schema: schema.GetSchema()})
	http.Handle("/graphql", &relay.Handler{Schema: schema.GetGqlSchema()})
	fmt.Println("server running at port 8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
