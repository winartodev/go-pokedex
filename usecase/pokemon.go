package usecase

import (
	"context"
	"database/sql"

	"github.com/winartodev/go-pokedex/entity"
	pokemonrepository "github.com/winartodev/go-pokedex/repository/pokemon"
	pokemontyperepository "github.com/winartodev/go-pokedex/repository/pokemontypes"
)

type PokemonUsecase struct {
	PokemonRepository     pokemonrepository.PokemonRepositoryItf
	PokemonTypeRepository pokemontyperepository.PokemonTypeRepositoryItf
}

type PokemonUsecaseItf interface {
	GetAllPokemon(ctx context.Context) (results []entity.PokemonList, err error)
	GetAllPokemonByFilter(ctx context.Context, filter map[string]string) (results []entity.PokemonList, err error)
	CatchPokemon(ctx context.Context, id int64) (err error)
	CreatePokemon(ctx context.Context, data entity.Pokemon) (pokemonID int64, err error)
	GetPokemonByID(ctx context.Context, id int64) (result *entity.PokemonDetail, err error)
	UpdatePokemon(ctx context.Context, id int64, data entity.Pokemon) (result *entity.PokemonDetail, err error)
	DeletePokemon(ctx context.Context, id int64) (err error)
}

const (
	CATCH   = 1
	DELETED = 0
)

func NewPokemonUsecase(pokemonUsecase PokemonUsecase) PokemonUsecaseItf {
	return &PokemonUsecase{
		PokemonRepository:     pokemonUsecase.PokemonRepository,
		PokemonTypeRepository: pokemonUsecase.PokemonTypeRepository,
	}
}

func (pu *PokemonUsecase) GetAllPokemon(ctx context.Context) (results []entity.PokemonList, err error) {
	res, err := pu.PokemonRepository.GetAllPokemonDB(ctx)
	if err != nil {
		return results, err
	}

	return pu.buildResponsePokemonList(ctx, res)
}

func (pu *PokemonUsecase) GetAllPokemonByFilter(ctx context.Context, filter map[string]string) (results []entity.PokemonList, err error) {
	res, err := pu.PokemonRepository.GetAllPokemonByFilterDB(ctx, filter)
	if err != nil {
		return results, err
	}

	return pu.buildResponsePokemonList(ctx, res)
}

func (pu *PokemonUsecase) CreatePokemon(ctx context.Context, data entity.Pokemon) (pokemonID int64, err error) {
	pokemon, err := pu.buildPokemonFromRequest(data)
	if err != nil {
		return pokemonID, err
	}

	pokemonID, err = pu.PokemonRepository.CreatePokemonDB(ctx, pokemon)
	if err != nil {
		return pokemonID, err
	}

	for _, typeID := range data.Types {
		err = pu.PokemonTypeRepository.CreatePokemonTypeDB(ctx, entity.PokemonType{PokemonID: pokemonID, TypeID: typeID})
		if err != nil {
			return pokemonID, err
		}
	}

	return pokemonID, err
}

func (pu *PokemonUsecase) GetPokemonByID(ctx context.Context, id int64) (result *entity.PokemonDetail, err error) {
	pokemon, err := pu.PokemonRepository.GetPokemonByIDDB(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return result, err
	}

	return pu.buildResponsePokemonDetail(ctx, pokemon)
}

func (pu *PokemonUsecase) UpdatePokemon(ctx context.Context, id int64, data entity.Pokemon) (result *entity.PokemonDetail, err error) {
	pokemonData, err := pu.buildPokemonFromRequest(data)
	if err != nil {
		return result, err
	}

	err = pu.PokemonRepository.UpdatePokemonDB(ctx, id, pokemonData)
	if err != nil {
		return result, err
	}

	pokemonType, err := pu.PokemonTypeRepository.GetPokemonTypeByPokemonIDDB(ctx, id)
	if err != nil {
		return result, err
	}

	// will update and insert new data if length request is greather than or equal length data pokemon type that obtained from database.
	if len(data.Types) >= len(pokemonType) {
		for i := 0; i < len(data.Types); i++ {
			if len(pokemonType) > i {
				if pokemonType[i].TypeID != data.Types[i] {
					err = pu.PokemonTypeRepository.UpdatePokemonTypeDB(ctx, pokemonType[i].ID, entity.PokemonType{PokemonID: id, TypeID: data.Types[i]})
					if err != nil {
						return result, err
					}
				}
			} else {
				err = pu.PokemonTypeRepository.CreatePokemonTypeDB(ctx, entity.PokemonType{PokemonID: id, TypeID: data.Types[i]})
				if err != nil {
					return result, err
				}
			}
		}
	}

	// will soft delete when length request less than length data pokemon type that obtained from database.
	if len(data.Types) < len(pokemonType) {
		for i := 0; i < len(pokemonType); i++ {
			if len(data.Types) > i {
				err = pu.PokemonTypeRepository.UpdatePokemonTypeDB(ctx, pokemonType[i].ID, entity.PokemonType{PokemonID: id, TypeID: data.Types[i]})
				if err != nil {
					return result, err
				}
			} else {
				err = pu.PokemonTypeRepository.UpdatePokemonTypeDB(ctx, pokemonType[i].ID, entity.PokemonType{PokemonID: id, TypeID: DELETED})
				if err != nil {
					return result, err
				}
			}
		}
	}

	pokemon, err := pu.PokemonRepository.GetPokemonByIDDB(ctx, id)
	if err != nil {
		return result, err
	}

	return pu.buildResponsePokemonDetail(ctx, pokemon)
}

func (pu *PokemonUsecase) DeletePokemon(ctx context.Context, id int64) (err error) {
	err = pu.PokemonRepository.DeletePokemonByIDDB(ctx, id)
	if err != nil {
		return err
	}

	err = pu.PokemonTypeRepository.DeletePokemonTypeByPokemonIDDB(ctx, id)
	if err != nil {
		return err
	}

	return err
}

func (pu *PokemonUsecase) CatchPokemon(ctx context.Context, id int64) (err error) {
	pokemon, err := pu.PokemonRepository.GetPokemonByIDDB(ctx, id)
	if err != nil {
		return err
	}

	pokemon.Catched = CATCH

	err = pu.PokemonRepository.UpdatePokemonDB(ctx, id, pokemon)
	if err != nil {
		return err
	}

	return err
}
