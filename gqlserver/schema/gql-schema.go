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
	}

	type Mutation {
		addTodo(title: String!, body: String!, status: String!): Todo
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
`

func GetGqlSchema() *graphql.Schema {
	gqlSchema := graphql.MustParseSchema(gqlSchema, &resolver.Resolver{})
	return gqlSchema
}
