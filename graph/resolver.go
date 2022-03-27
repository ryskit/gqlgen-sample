package graph

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/ryskit/gqlgen-sample/graph/model"
	"github.com/ryskit/gqlgen-sample/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	DB    *sql.DB
}

func (r *mutationResolver) createUser(ctx context.Context, input model.NewUser) (*model.User, error) {
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

func (r *mutationResolver) createTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
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

func (r *queryResolver) queryUsers(ctx context.Context) ([]*model.User, error) {
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

func (r *todoResolver) relationTodoUser(ctx context.Context, obj *model.Todo) (*model.User, error) {
	u, err := models.FindUser(ctx, r.DB, *obj.UserID)
	if err != nil {
		return nil, err
	}
	uM := &model.User{
		ID:   u.ID,
		Name: u.Name,
	}
	return uM, nil
}

func (r *queryResolver) queryTodos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := models.Todos().All(ctx, r.DB)
	if err != nil {
		return nil, err
	}
	var ts []*model.Todo
	for _, t := range todos {
		tM := &model.Todo{
			ID:     t.ID,
			Text:   t.Text,
			Done:   t.Done,
			UserID: &t.UserID,
		}
		ts = append(ts, tM)
	}
	return ts, nil
}
