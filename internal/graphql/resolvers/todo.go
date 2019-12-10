package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) AddTodo(ctx context.Context, input *gqlmodels.NewTodo) (string, error) {
	// Implement your login logic here
	return "MyFakeToken", nil
}


// Queries
func (r *queryResolver) Todos(ctx context.Context, userId string (*[]models.Todo, error) {
	panic("not implemented")
}
