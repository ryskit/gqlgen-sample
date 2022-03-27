package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/ryskit/gqlgen-sample/graph/generated"
	"github.com/ryskit/gqlgen-sample/graph/model"
	"github.com/ryskit/gqlgen-sample/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	u4, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	id := u4.String()
	tD := &models.Todo{
		ID:     id,
		Text:   input.Text,
		UserID: input.UserID,
	}
	err = tD.Insert(ctx, r.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	todo := &model.Todo{
		Text:   input.Text,
		ID:     id,
		UserID: &input.UserID,
	}
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	u4, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	uD := &models.User{
		ID:   u4.String(),
		Name: input.Name,
	}
	err = uD.Insert(ctx, r.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	u := &model.User{
		ID:   u4.String(),
		Name: input.Name,
	}
	return u, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := models.Users().All(ctx, r.DB)
	if err != nil {
		return nil, err
	}
	var us []*model.User
	for _, u := range users {
		uM := model.User{
			ID:   u.ID,
			Name: u.Name,
		}
		us = append(us, &uM)
	}
	return us, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{
		ID:   *obj.UserID,
		Name: "user " + *obj.UserID,
	}, nil
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
