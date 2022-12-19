package pokemonrepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/winartodev/go-pokedex/entity"
)

type PokemonRepository struct {
	PokemonDB *sql.DB
}

type PokemonRepositoryItf interface {
	GetAllPokemonDB(ctx context.Context) (results []entity.PokemonDB, err error)
	GetAllPokemonByFilterDB(ctx context.Context, filter map[string]string) (results []entity.PokemonDB, err error)
	CreatePokemonDB(ctx context.Context, data entity.PokemonDB) (id int64, err error)
	GetPokemonByIDDB(ctx context.Context, id int64) (result entity.PokemonDB, err error)
	UpdatePokemonDB(ctx context.Context, id int64, data entity.PokemonDB) (err error)
	DeletePokemonByIDDB(ctx context.Context, id int64) (err error)
}

func NewPokemonRepository(db *sql.DB) PokemonRepositoryItf {
	return &PokemonRepository{
		PokemonDB: db,
	}
}

func (pr *PokemonRepository) GetAllPokemonDB(ctx context.Context) (results []entity.PokemonDB, err error) {
	rows, err := pr.PokemonDB.QueryContext(ctx, fmt.Sprintf("%s %s", GetPokemonQuery, `GROUP BY pokemons.id`))
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var row entity.PokemonDB

		err := rows.Scan(&row.ID, &row.Name, &row.Species, &row.Catched, &row.Metadata)
		if err != nil {
			return results, err
		}

		results = append(results, row)
	}

	return results, err
}

func (pr *PokemonRepository) CreatePokemonDB(ctx context.Context, data entity.PokemonDB) (id int64, err error) {
	row, err := pr.PokemonDB.ExecContext(ctx, InsertPokemonQuery, &data.Name, &data.Species, &data.Catched, &data.Metadata)
	if err != nil {
		return id, err
	}

	id, err = row.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, err
}

func (pr *PokemonRepository) GetPokemonByIDDB(ctx context.Context, id int64) (result entity.PokemonDB, err error) {
	err = pr.PokemonDB.QueryRowContext(ctx, fmt.Sprintf(`%s %s`, GetPokemonQuery, `WHERE pokemons.id = ?`), id).Scan(&result.ID, &result.Name, &result.Species, &result.Catched, &result.Metadata)
	if err != nil {
		return result, err
	}

	return result, err
}

func (pr *PokemonRepository) UpdatePokemonDB(ctx context.Context, id int64, data entity.PokemonDB) (err error) {
	_, err = pr.PokemonDB.ExecContext(ctx, UpdatePokemonQuery, &data.Name, &data.Species, &data.Catched, &data.Metadata, id)
	if err != nil {
		return err
	}

	return err
}

func (pr *PokemonRepository) DeletePokemonByIDDB(ctx context.Context, id int64) (err error) {
	_, err = pr.PokemonDB.ExecContext(ctx, DeletePokemonQuery, id)
	if err != nil {
		return err
	}

	return err
}

func (pr *PokemonRepository) GetAllPokemonByFilterDB(ctx context.Context, filter map[string]string) (pokemons []entity.PokemonDB, err error) {
	query := GetPokemonQuery

	query += fmt.Sprint(`WHERE pokemons.name LIKE '%`, filter["name"], `%'`)

	if filter["options"] != "" {
		query += fmt.Sprintf(`AND pokemons.catched = %v `, filter["options"])
	}

	if filter["type"] != "" {
		query += fmt.Sprintf(`AND pokemon_types.types_id IN (%v) `, filter["type"])
	}

	query += `GROUP BY pokemons.id `

	if filter["sort_by"] != "" && filter["order_by"] != "" {
		query += fmt.Sprintf("ORDER BY %v %v", filter["sort_by"], filter["order_by"])
	}

	rows, err := pr.PokemonDB.QueryContext(ctx, query)
	if err != nil {
		return pokemons, err
	}

	for rows.Next() {
		var row entity.PokemonDB

		err := rows.Scan(&row.ID, &row.Name, &row.Species, &row.Catched, &row.Metadata)
		if err != nil {
			return pokemons, err
		}

		pokemons = append(pokemons, row)
	}

	return pokemons, err
}
