package main

import (
	"context"
	"fmt"
	"gitHub.com/apigee/apigee-gqlserver/gqlserver/schema"
	"github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"

	pb "gitHub.com/apigee/apigee-gqlserver/greeting"
)

// printFeature gets the feature for the given point.
func sayHello(client pb.GreetingServiceClient, message *pb.Message) {
	log.Printf("Say Hello with message (%v)", message.Body)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	greeting, err := client.SayHello(ctx, message)
	if err != nil {
		log.Fatalf("client.SayHello failed: %v", err)
	}
	log.Println(greeting)
}

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetingServiceClient(conn)

	sayHello(client, &pb.Message{Body: "World!"})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status: OK")
	})

	//gqlSchema := graphql.MustParseSchema(, &resolver.Resolver{})
	//http.Handle("/graphql", &relay.Handler{Schema: schema.GetSchema()})
	http.Handle("/graphql", &relay.Handler{Schema: schema.GetGqlSchema()})
	fmt.Println("server running at port 8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
