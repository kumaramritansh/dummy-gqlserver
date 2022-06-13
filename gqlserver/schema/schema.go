package schema

import (
	"gitHub.com/apigee/apigee-gqlserver/gqlserver/resolver"
	"github.com/graph-gophers/graphql-go"
)

var schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	type Query {
		dummy: Dummy
		user(id: Int!): User
		users: Users
	}

	type Mutation {
		addUser(name: String!, email: String!, profession: String!): User
	}

	type Dummy {
		message: String
		status: String
	}

	type User {
		id: Int
		name: String
		email: String
		profession: String
	}

	type Users {
		data: [User]
	}
`

func GetSchema() *graphql.Schema {
	gqlSchema := graphql.MustParseSchema(schema, &resolver.Resolver{})
	return gqlSchema
}
