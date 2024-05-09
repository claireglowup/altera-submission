package repository

import (
	"context"
	"fmt"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
)

func (r *repository) Register(ctx context.Context, user *model.User) error {

	err := r.config.Database().WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}

	fmt.Println(user)
	return nil

}

func (r *repository) Login(ctx context.Context, user *request.UserLogin) (*model.User, error) {

	var data *model.User

	err := r.config.Database().WithContext(ctx).Where("email = ?", user.Email).First(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}
