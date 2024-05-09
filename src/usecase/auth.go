package usecase

import (
	"context"
	"errors"
	"os"
	"restapi-altera/src/model"
	"restapi-altera/src/model/request"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Register(ctx context.Context, user request.UserRegister) error {

	uuId := uuid.New()
	uuId2 := uuid.New()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return err
	}

	newUser := new(model.User)
	newUser.Id = uuId.String()
	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.RackId = uuId2.String()
	newUser.Password = string(passwordHash)

	err = u.repo.Register(ctx, newUser)
	if err != nil {
		return err
	}

	return nil

}

func (u *usecase) Login(ctx context.Context, user request.UserLogin) (string, error) {

	data, err := u.repo.Login(ctx, &user)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    data.Id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil

}
