package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/silvercrow/go-graphql/database"
	"github.com/silvercrow/go-graphql/graph/generated"
	"github.com/silvercrow/go-graphql/graph/model"
)

func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	return db.Save(input), nil
}

func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	return db.FindByID(id), nil
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
