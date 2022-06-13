package schema

import (
	"gitHub.com/apigee/apigee-gqlserver/gqlserver/resolver"
	"github.com/graph-gophers/graphql-go"
)

var gqlSchema = `
	schema {
		query: Query
		mutation: Mutation
	}

	type Query {
		todo(id: Int!): Todo
		todos: Todos
		events: Events
	}

	type Mutation {
		addTodo(title: String!, body: String!, status: String!): Todo
		createEvent(Id: String!, Title: String!, Description: String!): Event
	}

	type Todo {
		id: Int
		title: String
		body: String
		status: String
	}

	type Todos {
		data: [Todo]
	}

	type Event {
		Id: String
		Title: String
		Description: String
	}

	type Events {
		data: [Event]
	}
`

func GetGqlSchema() *graphql.Schema {
	gqlSchema := graphql.MustParseSchema(gqlSchema, &resolver.Resolver{})
	return gqlSchema
}
