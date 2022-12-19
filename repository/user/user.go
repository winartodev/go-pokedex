package userrepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/winartodev/go-pokedex/entity"
)

type UserRepository struct {
	DB *sql.DB
}

type UserRepositoryItf interface {
	CreateUser(ctx context.Context, username string, email string, password string, role int64) (id int64, err error)
	GetUserByUsername(ctx context.Context, username string) (result entity.User, err error)
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func CreateTableUsers(db *sql.DB) (err error) {
	_, err = db.Exec(CreateTableUsersQuery)
	if err != nil {
		return err
	}

	return err
}

func (ur *UserRepository) CreateUser(ctx context.Context, username string, email string, password string, role int64) (id int64, err error) {
	row, err := ur.DB.ExecContext(ctx, InsertUserQuery, username, email, password, role)
	if err != nil {
		return id, err
	}

	id, err = row.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, err
}

func (ur *UserRepository) GetUserByUsername(ctx context.Context, username string) (result entity.User, err error) {
	err = ur.DB.QueryRowContext(ctx, fmt.Sprintf(`%v %v`, GetUserQuery, `WHERE username = ?`), username).Scan(&result.ID, &result.Username, &result.Email, &result.Password, &result.Role)
	if err != nil {
		return result, err
	}

	return result, err
}
