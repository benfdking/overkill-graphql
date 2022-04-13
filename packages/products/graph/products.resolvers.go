package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bendfking/overkill-graphql/packages/products/graph/generated"
	"github.com/bendfking/overkill-graphql/packages/products/graph/model"
)

func (r *queryResolver) ListProducts(ctx context.Context) ([]*model.Product, error) {
	return []*model.Product{{
		ID:    "product id 1",
		Name:  "Test 1",
		Price: "100",
	}}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
