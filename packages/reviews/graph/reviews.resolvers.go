package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bendfking/overkill-graphql/packages/reviews/graph/generated"
	"github.com/bendfking/overkill-graphql/packages/reviews/graph/model"
)

func (r *queryResolver) ListReviews(ctx context.Context) ([]*model.Review, error) {
	return []*model.Review{{
		ID:      "Id 1",
		Author:  &model.User{ID: "Id 1"},
		Product: &model.Product{ID: "Id 1"},
	}}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
