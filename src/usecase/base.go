package usecase

import (
	"context"
	"os"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
	"restapi-altera/src/model/response"
	"restapi-altera/src/repository"

	"github.com/golang-jwt/jwt/v4"
)

type Usecase interface {
	Register(ctx context.Context, user request.UserRegister) error
	Login(ctx context.Context, user request.UserLogin) (string, error)

	GetUserByID(ctx context.Context, cookie string) (*model.User, error)
	GetUserAllShoes(ctx context.Context, cookie string) (*response.UserShoes, error)
	GetUserShoesByNoId(ctx context.Context, noId int, cookie string) (*response.UserShoes, error)
	AddUserShoes(ctx context.Context, cookie string, shoes *request.ReqUserShoes) error
	UpdateShoesByNoId(ctx context.Context, noId int, cookie string, shoes *request.UpdateUserShoes) error
	DeleteShoesByNoId(ctx context.Context, noId int, cookie string) error
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) getJWTClaims(cookie string) (*jwt.RegisteredClaims, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	return claims, nil

}
