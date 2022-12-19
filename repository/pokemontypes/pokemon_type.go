package pokemontyperepository

import (
	"context"
	"database/sql"

	"github.com/winartodev/go-pokedex/entity"
)

type PokemonTypeRepository struct {
	PokemonTypeDB *sql.DB
}

type PokemonTypeRepositoryItf interface {
	CreatePokemonTypeDB(ctx context.Context, data entity.PokemonType) (err error)
	GetPokemonTypeByPokemonIDDB(ctx context.Context, pokemonID int64) (result []entity.PokemonType, err error)
	UpdatePokemonTypeDB(ctx context.Context, id int64, data entity.PokemonType) (err error)
	DeletePokemonTypeByPokemonIDDB(ctx context.Context, pokemonID int64) (err error)
	DeletePokemonTypeByIDDB(ctx context.Context, id int64) (err error)
}

func NewPokemonTypeRepository(db *sql.DB) PokemonTypeRepositoryItf {
	return &PokemonTypeRepository{
		PokemonTypeDB: db,
	}
}

func CreateTablePokemonTypes(db *sql.DB) (err error) {
	_, err = db.Exec(CreateTablePokemonTypesQuery)
	if err != nil {
		return err
	}

	return err
}

func (pt *PokemonTypeRepository) CreatePokemonTypeDB(ctx context.Context, data entity.PokemonType) (err error) {
	_, err = pt.PokemonTypeDB.ExecContext(ctx, InsertPokemonTypeQuery, &data.PokemonID, &data.TypeID)
	if err != nil {
		return err
	}

	return err
}

func (pt *PokemonTypeRepository) GetPokemonTypeByPokemonIDDB(ctx context.Context, pokemonID int64) (result []entity.PokemonType, err error) {
	rows, err := pt.PokemonTypeDB.QueryContext(ctx, GetPokemonTypesByPokemonIDQuery, pokemonID)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var row entity.PokemonType

		err = rows.Scan(&row.ID, &row.PokemonID, &row.TypeID, &row.Name)
		if err != nil {
			return result, err
		}

		result = append(result, row)
	}

	return result, err
}

func (pt *PokemonTypeRepository) UpdatePokemonTypeDB(ctx context.Context, id int64, data entity.PokemonType) (err error) {
	_, err = pt.PokemonTypeDB.ExecContext(ctx, UpdatePokemonTokenQuery, &data.PokemonID, &data.TypeID, id)
	if err != nil {
		return err
	}

	return err
}

func (pt *PokemonTypeRepository) DeletePokemonTypeByPokemonIDDB(ctx context.Context, pokemonID int64) (err error) {
	_, err = pt.PokemonTypeDB.ExecContext(ctx, DeletePokemonTypeByPokemonIDQuery, pokemonID)
	if err != nil {
		return err
	}

	return err
}

func (pt *PokemonTypeRepository) DeletePokemonTypeByIDDB(ctx context.Context, id int64) (err error) {
	_, err = pt.PokemonTypeDB.ExecContext(ctx, DeletePokemonTypeByIDQuery, id)
	if err != nil {
		return err
	}

	return err
}
