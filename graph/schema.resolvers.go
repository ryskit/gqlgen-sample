package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/ryskit/gqlgen-sample/graph/generated"
	"github.com/ryskit/gqlgen-sample/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return r.createTodo(ctx, input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.createUser(ctx, input)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.queryTodos(ctx)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.queryUsers(ctx)
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return r.relationTodoUser(ctx, obj)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
