package usecase

import (
	"context"
	"errors"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
	"restapi-altera/src/model/response"
)

func (u *usecase) GetUserByID(ctx context.Context, cookie string) (*model.User, error) {

	claims, err := u.getJWTClaims(cookie)
	if err != nil {
		return nil, err
	}

	user, err := u.repo.GetUserByID(ctx, claims.Issuer)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil

}

func (u *usecase) GetUserAllShoes(ctx context.Context, cookie string) (*response.UserShoes, error) {

	claims, err := u.getJWTClaims(cookie)
	if err != nil {
		return nil, err
	}

	data, err := u.repo.GetUserAllShoes(ctx, claims.Issuer)
	if err != nil {
		return nil, err
	}

	return data, nil

}
func (u *usecase) GetUserShoesByNoId(ctx context.Context, noId int, cookie string) (*response.UserShoes, error) {

	claims, err := u.getJWTClaims(cookie)
	if err != nil {
		return nil, err
	}

	data, err := u.repo.GetUserShoesByNoId(ctx, noId, claims.Issuer)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (u *usecase) AddUserShoes(ctx context.Context, cookie string, shoes *request.ReqUserShoes) error {

	claims, err := u.getJWTClaims(cookie)
	if err != nil {
		return err
	}

	err = u.repo.AddShoes(ctx, shoes, claims.Issuer)
	if err != nil {
		return err
	}

	return nil

}

func (u *usecase) UpdateShoesByNoId(ctx context.Context, noId int, cookie string, shoes *request.UpdateUserShoes) error {
	claims, err := u.getJWTClaims(cookie)
	if err != nil {
		return err
	}

	err = u.repo.UpdateShoesByNoId(ctx, noId, claims.Issuer, shoes)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) DeleteShoesByNoId(ctx context.Context, noId int, cookie string) error {
	claims, err := u.getJWTClaims(cookie)
	if err != nil {
		return err
	}

	err = u.repo.DeleteShoesByNoId(ctx, claims.Issuer, noId)
	if err != nil {
		return err
	}

	return nil
}
