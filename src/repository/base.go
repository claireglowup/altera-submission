package repository

import (
	"context"
	"restapi-altera/config"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
	"restapi-altera/src/model/response"
)

type Repository interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *request.UserLogin) (*model.User, error)

	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserAllShoes(ctx context.Context, id string) (*response.UserShoes, error)
	GetUserShoesByNoId(ctx context.Context, noId int, id string) (*response.UserShoes, error)
	AddShoes(ctx context.Context, shoes *request.ReqUserShoes, id string) error
	UpdateShoesByNoId(ctx context.Context, noId int, id string, payload *request.UpdateUserShoes) error
	DeleteShoesByNoId(ctx context.Context, id string, noId int) error
}

type repository struct {
	config      config.Config
	keyAllShoes string
	keyById     string
	keyUser     string
}

func NewRepository(config config.Config) Repository {
	return &repository{
		config: config,
	}
}
