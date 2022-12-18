package usecase

import (
	"context"
	"encoding/json"

	"github.com/winartodev/go-pokedex/entity"
)

type metadata struct {
	ImageURL    string       `json:"image_url"`
	Description string       `json:"description"`
	Weight      float64      `json:"weight"`
	Height      float64      `json:"height"`
	Stats       entity.Stats `json:"stats"`
}

func (pu *PokemonUsecase) buildResponsePokemonList(ctx context.Context, pokemons []entity.PokemonDB) (result []entity.PokemonList, err error) {
	for _, pokemon := range pokemons {
		pokemonTypes, err := pu.PokemonTypeRepository.GetPokemonTypeByPokemonIDDB(ctx, pokemon.ID)
		if err != nil {
			return result, err
		}

		var types []string
		for i := range pokemonTypes {
			if pokemonTypes[i].TypeID > 0 {
				types = append(types, pokemonTypes[i].Name)
			}
		}

		var metadata metadata
		err = json.Unmarshal([]byte(pokemon.Metadata), &metadata)
		if err != nil {
			return result, err
		}

		result = append(result, entity.PokemonList{
			ID:       pokemon.ID,
			Name:     pokemon.Name,
			Species:  pokemon.Species,
			Types:    types,
			Catched:  pokemon.Catched,
			ImageURL: metadata.ImageURL,
		})
	}

	return result, err
}

func (pu *PokemonUsecase) buildResponsePokemonDetail(ctx context.Context, data entity.PokemonDB) (result *entity.PokemonDetail, err error) {
	pokemonTypes, err := pu.PokemonTypeRepository.GetPokemonTypeByPokemonIDDB(ctx, data.ID)
	if err != nil {
		return result, err
	}

	var types []string
	for i := range pokemonTypes {
		if pokemonTypes[i].TypeID > 0 {
			types = append(types, pokemonTypes[i].Name)
		}
	}

	var metadata metadata
	err = json.Unmarshal([]byte(data.Metadata), &metadata)
	if err != nil {
		return result, err
	}

	return &entity.PokemonDetail{
		ID:          data.ID,
		Name:        data.Name,
		Species:     data.Species,
		Types:       types,
		Catched:     data.Catched,
		ImageURL:    metadata.ImageURL,
		Description: metadata.Description,
		Weight:      metadata.Weight,
		Height:      metadata.Height,
		Stats:       metadata.Stats,
	}, err
}

// buildPokemonFromRequest is function to build from body request
func (pu *PokemonUsecase) buildPokemonFromRequest(data entity.Pokemon) (result entity.PokemonDB, err error) {
	metadata, err := json.Marshal(&metadata{
		ImageURL:    data.ImageURL,
		Description: data.Description,
		Weight:      data.Weight,
		Height:      data.Height,
		Stats:       data.Stats,
	})

	if err != nil {
		return result, err
	}

	return entity.PokemonDB{
		ID:       data.ID,
		Name:     data.Name,
		Species:  data.Species,
		Catched:  data.Catched,
		Metadata: string(metadata),
	}, nil
}
