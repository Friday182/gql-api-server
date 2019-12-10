package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) Signin(ctx context.Context, input *gqlmodels.UserInput) (string, error) {
	// Implement your login logic here
	return "MyFakeToken", nil
}

func (r *mutationResolver) Signup(ctx context.Context, input *gqlmodels.UserInput) (string, error) {
	// Create new user
	// email := c.get("email")
	if r.db.find(&email).RecordNotFound() {
		// New user, create account
	} else {
		// exist user, go to signin page
	}
	return "MyFakeToken", nil
}

// Queries
func (r *queryResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	panic("not implemented")
}
