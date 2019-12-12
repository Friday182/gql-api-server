package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/friday182/gql-api-server/internal/graphql/generated"
	// "github.com/friday182/gql-api-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) Signup(ctx context.Context, email string, password string) (*generated.User, error) {
	// Create new user
	// email := c.get("email")
	if r.Db.Find(&email).RecordNotFound() {
		// New user, create account
	} else {
		// exist user, go to signin page
	}
	var tmp *generated.User
	return tmp, nil
}
func (r *mutationResolver) Signin(ctx context.Context, email string, token string) (*generated.User, error) {
	panic("not implemented")
}

