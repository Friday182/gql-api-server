package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/friday182/gql-api-server/internal/auth"
	gql "github.com/friday182/gql-api-server/internal/graphql/generated"
	res "github.com/friday182/gql-api-server/internal/graphql/resolvers"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(db *gorm.DB) gin.HandlerFunc {

	// Verify token
	// h := auth.Middleware(
		cfg := gql.Config{
			Resolvers: &res.Resolver{
				Db: db,
			},
		}
	h := handler.GraphQL(gql.NewExecutableSchema(cfg))
	// )

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
