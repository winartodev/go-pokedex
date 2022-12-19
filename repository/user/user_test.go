package userrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/winartodev/go-pokedex/entity"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("%s", err)
	}

	return db, mock
}

func TestNewUserRepository(t *testing.T) {
	db, _ := NewMock()
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *UserRepository
	}{
		{
			name: "success",
			args: args{
				db: db,
			},
			want: &UserRepository{
				DB: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableUsers(t *testing.T) {
	db, dbmock := NewMock()
	query := CreateTableUsersQuery
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				db,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: "failed",
			args: args{
				db,
			},
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTableUsers(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("CreateTableUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_CreateUser(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := InsertUserQuery
	user := entity.User{
		Username: "ganteng",
		Email:    "ganteng@mail.com",
		Password: "ganteng banget",
		Role:     1,
	}

	type fields struct {
		DB *sql.DB
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
				DB: db,
			},
			args: args{
				ctx:      ctx,
				username: "ganteng",
				email:    "ganteng@mail.com",
				password: "ganteng banget",
				role:     1,
			},
			wantId:  1,
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(user.Username, user.Email, user.Password, user.Role).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:      ctx,
				username: "ganteng",
				email:    "ganteng@mail.com",
				password: "ganteng banget",
				role:     1,
			},
			wantId:  0,
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(user.Username, user.Email, user.Password, user.Role).
					WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepository{
				DB: tt.fields.DB,
			}
			gotId, err := ur.CreateUser(tt.args.ctx, tt.args.username, tt.args.email, tt.args.password, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("UserRepository.CreateUser() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	username := "ganteng"
	query := fmt.Sprintf(`%v %v`, GetUserQuery, `WHERE username = ?`)
	user := entity.User{
		Username: "ganteng",
		Email:    "ganteng@mail.com",
		Password: "ganteng banget",
		Role:     1,
	}

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult entity.User
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:      ctx,
				username: "ganteng",
			},
			wantResult: user,
			wantErr:    false,
			mock: func() {
				dbmock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(username).WillReturnRows(
					sqlmock.NewRows([]string{"id", "username", "email", "password", "role"}).
						AddRow(user.ID, user.Username, user.Email, user.Password, user.Role),
				)
			},
		},
		{
			name: "failed",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:      ctx,
				username: "ganteng",
			},
			wantResult: entity.User{},
			wantErr:    true,
			mock: func() {
				dbmock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(username).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepository{
				DB: tt.fields.DB,
			}
			gotResult, err := ur.GetUserByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("UserRepository.GetUserByUsername() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
