package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/Divan009/DumGo/graph/jwt"

	"github.com/Divan009/DumGo/graph/generated"
	"github.com/Divan009/DumGo/graph/logic"
	"github.com/Divan009/DumGo/graph/model"
	"github.com/Divan009/DumGo/graph/postgres"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	link, err := logic.AddLink(input.Title, input.Address)
	return link, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := postgres.User{
		Username: input.Username,
		Password: input.Password,
	}
	// user.Username = input.Username
	// user.Password = input.Password
	msg, err := postgres.Create(user)
	log.Print(msg)
	token, err := jwt.GenerateToken(input.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	links, err := logic.GetLinks()

	if err != nil {
		fmt.Println("Wrong mistake")
	}
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
