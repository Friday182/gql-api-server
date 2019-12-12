package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/friday182/gql-api-server/internal/graphql/generated"
	// "github.com/friday182/gql-api-server/internal/orm/models"
)

// Query
func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*generated.Todo, error) {
	var tmp []*generated.Todo
	t1 := generated.Todo{
		ID: "10",
		Text: "this is a test",
		Done: true,
	}
	tmp = append(tmp, &t1)
	return tmp, nil
}

// Mutation
func (r *mutationResolver) AddTodo(ctx context.Context, input generated.NewTodo) (bool, error) {
	panic("not implemented")
}