package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	pokemonrepository "github.com/winartodev/go-pokedex/repository/pokemon"
	pokemontyperepository "github.com/winartodev/go-pokedex/repository/pokemontypes"
	typesrepository "github.com/winartodev/go-pokedex/repository/types"
	userrepository "github.com/winartodev/go-pokedex/repository/user"
)

// MakeMigrate is function to migrate database and returning error
func MakeMigrate(db *sql.DB) (err error) {
	err = pokemonrepository.CreateTablePokemons(db)
	if err != nil {
		return err
	}

	err = pokemontyperepository.CreateTablePokemonTypes(db)
	if err != nil {
		return err
	}

	err = typesrepository.CreateTableTypes(db)
	if err != nil {
		return err
	}

	err = userrepository.CreateTableUsers(db)
	if err != nil {
		return err
	}

	return err
}

// NewDatabase is function to make connection to database
func NewDatabase(cfg Config) (db *sql.DB, err error) {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)

	db, err = sql.Open(cfg.Database.Connection, dbConfig)
	if err != nil {
		return db, err
	}

	_, err = db.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s`, cfg.Database.Database))
	if err != nil {
		return db, err
	}
	db.Close()

	db, err = sql.Open(cfg.Database.Connection, fmt.Sprint(dbConfig, cfg.Database.Database))
	if err != nil {
		return db, err
	}

	return db, err
}
