package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/winartodev/go-pokedex/entity"
	userrepository "github.com/winartodev/go-pokedex/repository/user"
	userrepositorymocks "github.com/winartodev/go-pokedex/repository/user/mocks"
)

type mockUserProvider struct {
	UserRepository *userrepositorymocks.UserRepositoryItf
}

func userProvider() mockUserProvider {
	return mockUserProvider{
		UserRepository: new(userrepositorymocks.UserRepositoryItf),
	}
}

func TestNewUserUsecase(t *testing.T) {
	userUsecase := UserUsecase{
		UserRepository: new(userrepositorymocks.UserRepositoryItf),
	}
	type args struct {
		userUsecase UserUsecase
	}
	tests := []struct {
		name string
		args args
		want UserUsecase
	}{
		{
			name: "success",
			args: args{
				userUsecase: userUsecase,
			},
			want: userUsecase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.userUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_Register(t *testing.T) {
	ctx := context.Background()
	prov := userProvider()

	type fields struct {
		UserRepository userrepository.UserRepositoryItf
	}
	type args struct {
		ctx      context.Context
		username string
		email    string
		password string
		role     int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			fields: fields{
				UserRepository: prov.UserRepository,
			},
			args: args{
				ctx:      ctx,
				username: "budi",
				email:    "budi@mail.com",
				password: "123",
				role:     1,
			},
			wantId:  1,
			wantErr: false,
			mock: func() {
				prov.UserRepository.On("GetUserByUsername", mock.Anything, mock.Anything).
					Return(entity.User{ID: 1, Username: "winarto"}, nil).Times(1)

				prov.UserRepository.On("CreateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(int64(1), nil).Times(1)
			},
		},
		{
			name: "failed username already taken",
			fields: fields{
				UserRepository: prov.UserRepository,
			},
			args: args{
				ctx:      ctx,
				username: "winarto",
				email:    "winarto@mail.com",
				password: "123",
				role:     1,
			},
			wantId:  0,
			wantErr: true,
			mock: func() {
				prov.UserRepository.On("GetUserByUsername", mock.Anything, mock.Anything).
					Return(entity.User{ID: 1, Username: "winarto"}, nil).Times(1)
			},
		},
		{
			name: "failed user data",
			fields: fields{
				UserRepository: prov.UserRepository,
			},
			args: args{
				ctx:      ctx,
				username: "winarto",
				email:    "winarto@mail.com",
				password: "123",
				role:     1,
			},
			wantId:  0,
			wantErr: true,
			mock: func() {
				prov.UserRepository.On("GetUserByUsername", mock.Anything, mock.Anything).
					Return(entity.User{}, errors.New("error")).Times(1)
			},
		},
		{
			name: "failed create user",
			fields: fields{
				UserRepository: prov.UserRepository,
			},
			args: args{
				ctx:      ctx,
				username: "budi",
				email:    "budi@mail.com",
				password: "123",
				role:     1,
			},
			wantId:  0,
			wantErr: true,
			mock: func() {
				prov.UserRepository.On("GetUserByUsername", mock.Anything, mock.Anything).
					Return(entity.User{ID: 1, Username: "winarto"}, nil).Times(1)

				prov.UserRepository.On("CreateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(int64(0), errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			uu := &UserUsecase{
				UserRepository: tt.fields.UserRepository,
			}
			gotId, err := uu.Register(tt.args.ctx, tt.args.username, tt.args.email, tt.args.password, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("UserUsecase.Register() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestUserUsecase_Login(t *testing.T) {
	ctx := context.Background()
	prov := userProvider()

	type fields struct {
		UserRepository userrepository.UserRepositoryItf
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			fields: fields{
				UserRepository: prov.UserRepository,
			},
			args: args{
				ctx:      ctx,
				username: "winarto",
				password: "123",
			},
			wantErr: false,
			mock: func() {
				prov.UserRepository.On("GetUserByUsername", mock.Anything, mock.Anything).
					Return(entity.User{ID: 1, Username: "winarto", Email: "winarto@mail.com", Password: "$2a$12$EuMhNWuTVUF9G8tYSgH5BuL.8JYvrCRiKEx3flcemaIDa7INrei96", Role: 1}, nil).Times(1)
			},
		},
		{
			name: "failed get user data",
			fields: fields{
				UserRepository: prov.UserRepository,
			},
			args: args{
				ctx:      ctx,
				username: "winarto",
				password: "123",
			},
			wantErr: true,
			mock: func() {
				prov.UserRepository.On("GetUserByUsername", mock.Anything, mock.Anything).
					Return(entity.User{}, errors.New("error")).Times(1)
			},
		},
		{
			name: "failed password not valid",
			fields: fields{
				UserRepository: prov.UserRepository,
			},
			args: args{
				ctx:      ctx,
				username: "winarto",
				password: "123333",
			},
			wantErr: true,
			mock: func() {
				prov.UserRepository.On("GetUserByUsername", mock.Anything, mock.Anything).
					Return(entity.User{ID: 1, Username: "winarto", Email: "winarto@mail.com", Password: "$2a$12$EuMhNWuTVUF9G8tYSgH5BuL.8JYvrCRiKEx3flcemaIDa7INrei96", Role: 1}, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			uu := &UserUsecase{
				UserRepository: tt.fields.UserRepository,
			}
			_, err := uu.Login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
