package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bendfking/overkill-graphql/packages/users/graph/generated"
	"github.com/bendfking/overkill-graphql/packages/users/graph/model"
)

func (r *queryResolver) ListUsers(ctx context.Context) ([]*model.User, error) {
	return []*model.User{
		{
			ID:   "id1",
			Name: "Name 1",
		},
		{
			ID:   "id2",
			Name: "Name 2",
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
