package typesrepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/winartodev/go-pokedex/entity"
)

type TypeRepository struct {
	TypeDB *sql.DB
}

type TypeRepositoryItf interface {
	CreateTypeDB(ctx context.Context, data entity.Type) (id int64, err error)
	GetAllTypeDB(ctx context.Context) (results []entity.Type, err error)
	GeTypeByIDDB(ctx context.Context, id int64) (result entity.Type, err error)
	UpdateTypeDB(ctx context.Context, id int64, data entity.Type) (err error)
}

func NewTypeRepository(db *sql.DB) TypeRepositoryItf {
	return &TypeRepository{
		TypeDB: db,
	}
}

func (tr *TypeRepository) CreateTypeDB(ctx context.Context, data entity.Type) (id int64, err error) {
	row, err := tr.TypeDB.ExecContext(ctx, InsertTypeQuery, &data.Name)
	if err != nil {
		return id, err
	}

	id, err = row.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, err
}

func (tr *TypeRepository) GetAllTypeDB(ctx context.Context) (results []entity.Type, err error) {
	rows, err := tr.TypeDB.QueryContext(ctx, GetTypesQuery)
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var row entity.Type

		err := rows.Scan(&row.ID, &row.Name)
		if err != nil {
			return results, err
		}

		results = append(results, row)
	}

	return results, err
}

func (tr *TypeRepository) GeTypeByIDDB(ctx context.Context, id int64) (result entity.Type, err error) {
	err = tr.TypeDB.QueryRowContext(ctx, fmt.Sprintf(`%s %s`, GetTypesQuery, `WHERE id = ?`), id).Scan(&result.ID, &result.Name)
	if err != nil {
		return result, err
	}

	return result, err
}

func (tr *TypeRepository) UpdateTypeDB(ctx context.Context, id int64, data entity.Type) (err error) {
	_, err = tr.TypeDB.ExecContext(ctx, UpdateTypeQuery, data.Name, id)
	if err != nil {
		return err
	}

	return err
}
