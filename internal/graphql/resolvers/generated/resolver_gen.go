package resolvers

import (
	"context"

	"github.com/friday182/gql-api-server/internal/graphql/generated"
	"github.com/jinzhu/gorm"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	Db *gorm.DB
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Signup(ctx context.Context, email string, password string) (*generated.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) Signin(ctx context.Context, email string, token string) (*generated.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddTodo(ctx context.Context, input generated.NewTodo) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*generated.Todo, error) {
	panic("not implemented")
}
