package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/winartodev/go-pokedex/enum"
	"github.com/winartodev/go-pokedex/middleware/auth"
	userrepository "github.com/winartodev/go-pokedex/repository/user"
	"github.com/winartodev/go-pokedex/util"
)

type UserUsecase struct {
	UserRepository userrepository.UserRepositoryItf
}

type UserUsecaseItf interface {
	Register(ctx context.Context, username string, email string, password string, role int64) (id int64, err error)
	Login(ctx context.Context, username string, password string) (token string, err error)
}

func NewUserUsecase(userUsecase UserUsecase) UserUsecase {
	return UserUsecase{
		userUsecase.UserRepository,
	}
}

func (uu *UserUsecase) Register(ctx context.Context, username string, email string, password string, role int64) (id int64, err error) {
	user, err := uu.UserRepository.GetUserByUsername(ctx, username)
	if err != nil && err != sql.ErrNoRows {
		return id, err
	}

	if user.Username == username {
		err = fmt.Errorf("username %s already taken", username)
		return id, err
	}

	passwordHash, err := util.HashPassword(password)

	if err != nil {
		return id, err
	}

	id, err = uu.UserRepository.CreateUser(ctx, username, email, passwordHash, role)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (uu *UserUsecase) Login(ctx context.Context, username string, password string) (token string, err error) {
	user, err := uu.UserRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return token, err
	}

	isValid := util.CheckPasswordHash(password, user.Password)
	if !isValid {
		return token, errors.New("username or password not valid")
	}

	token, err = auth.GenerateJWT(user.Username, user.Email, enum.Role(user.Role))
	if err != nil {
		return token, err
	}

	return token, err
}
