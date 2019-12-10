package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	// "github.com/friday182/gql-api-server/internal/auth"
	gql "github.com/friday182/gql-api-server/internal/graphql/generated"
	res "github.com/friday182/gql-api-server/internal/graphql/resolvers"
	"github.com/friday182/gql-api-server/internal/orm"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(db *gorm.DB) gin.HandlerFunc {

	// Verify token
	// h := auth.Middleware(
	handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &res.Resolver{db: db,}})),
	// )

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
