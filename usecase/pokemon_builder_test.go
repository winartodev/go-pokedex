package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/winartodev/go-pokedex/entity"
	pokemonrepositorymock "github.com/winartodev/go-pokedex/repository/pokemon/mocks"
	pokemontyperepositorymock "github.com/winartodev/go-pokedex/repository/pokemontypes/mocks"
)

type mockBuildPokemonProvider struct {
	PokemonRepository     *pokemonrepositorymock.PokemonRepositoryItf
	PokemonTypeRepository *pokemontyperepositorymock.PokemonTypeRepositoryItf
}

func buildPokemonProvider() mockBuildPokemonProvider {
	return mockBuildPokemonProvider{
		PokemonRepository:     new(pokemonrepositorymock.PokemonRepositoryItf),
		PokemonTypeRepository: new(pokemontyperepositorymock.PokemonTypeRepositoryItf),
	}
}

func newBuildPokemonProviderMock(repo PokemonUsecase) PokemonUsecase {
	return PokemonUsecase{
		PokemonRepository:     repo.PokemonRepository,
		PokemonTypeRepository: repo.PokemonTypeRepository,
	}
}

func TestPokemonUsecase_buildResponsePokemonList(t *testing.T) {
	ctx := context.Background()
	prov := buildPokemonProvider()

	pokemons := []entity.PokemonList{
		{
			ID:       1,
			Name:     "Bulbasour",
			Species:  "Seed Pokémon",
			Types:    nil,
			Catched:  1,
			ImageURL: "",
		},
	}
	type args struct {
		ctx      context.Context
		pokemons []entity.PokemonDB
	}
	tests := []struct {
		name       string
		args       args
		wantResult []entity.PokemonList
		wantErr    bool
		mock       func()
	}{
		{
			name: "fail GetPokemonTypeByPokemonIDDB",
			args: args{
				ctx:      ctx,
				pokemons: []entity.PokemonDB{{ID: 1, Name: "Bulbasour", Species: "Seed Pokémon", Catched: 1, Metadata: `{}`}},
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				prov.PokemonTypeRepository.Mock.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).Return([]entity.PokemonType{{ID: 1, Name: "Fire"}}, errors.New("error")).Times(1)
			},
		},
		{
			name: "marhsall error",
			args: args{
				ctx:      ctx,
				pokemons: []entity.PokemonDB{{ID: 1, Name: "Bulbasour", Species: "Seed Pokémon", Catched: 1, Metadata: ""}},
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				prov.PokemonTypeRepository.Mock.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).Return([]entity.PokemonType{{ID: 1, Name: "Fire"}}, nil).Times(1)
			},
		},
		{
			name: "success",
			args: args{
				ctx:      ctx,
				pokemons: []entity.PokemonDB{{ID: 1, Name: "Bulbasour", Species: "Seed Pokémon", Catched: 1, Metadata: `{}`}},
			},
			wantResult: pokemons,
			wantErr:    false,
			mock: func() {
				prov.PokemonTypeRepository.Mock.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).Return([]entity.PokemonType{{ID: 1, Name: "Fire"}}, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := newBuildPokemonProviderMock(PokemonUsecase{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			})

			gotResult, err := pu.buildResponsePokemonList(tt.args.ctx, tt.args.pokemons)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.buildResponsePokemonList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PokemonUsecase.buildResponsePokemonList() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPokemonUsecase_buildResponsePokemonDetail(t *testing.T) {
	ctx := context.Background()
	prov := buildPokemonProvider()
	pokemons := &entity.PokemonDetail{
		ID:       1,
		Name:     "Bulbasour",
		Species:  "Seed Pokémon",
		Types:    nil,
		Catched:  1,
		ImageURL: "",
	}

	type args struct {
		ctx  context.Context
		data entity.PokemonDB
	}
	tests := []struct {
		name       string
		args       args
		wantResult *entity.PokemonDetail
		wantErr    bool
		mock       func()
	}{
		{
			name: "fail GetPokemonTypeByPokemonIDDB",
			args: args{
				ctx:  ctx,
				data: entity.PokemonDB{ID: 1, Name: "Bulbasour", Species: "Seed Pokémon", Catched: 1, Metadata: `{}`},
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				prov.PokemonTypeRepository.Mock.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).Return([]entity.PokemonType{{ID: 1, Name: "Fire"}}, errors.New("error")).Times(1)
			},
		},
		{
			name: "marhsall error",
			args: args{
				ctx:  ctx,
				data: entity.PokemonDB{ID: 1, Name: "Bulbasour", Species: "Seed Pokémon", Catched: 1, Metadata: ""},
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				prov.PokemonTypeRepository.Mock.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).Return([]entity.PokemonType{{ID: 1, Name: "Fire"}}, errors.New("error")).Times(1)
			},
		},

		{
			name: "success",
			args: args{
				ctx:  ctx,
				data: entity.PokemonDB{ID: 1, Name: "Bulbasour", Species: "Seed Pokémon", Catched: 1, Metadata: `{}`},
			},
			wantResult: pokemons,
			wantErr:    false,
			mock: func() {
				prov.PokemonTypeRepository.Mock.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).Return([]entity.PokemonType{{ID: 1, Name: "Fire"}}, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := newBuildPokemonProviderMock(PokemonUsecase{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			})

			gotResult, err := pu.buildResponsePokemonDetail(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.buildResponsePokemonDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PokemonUsecase.buildResponsePokemonDetail() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPokemonUsecase_buildPokemonFromRequest(t *testing.T) {
	prov := buildPokemonProvider()

	type args struct {
		data entity.Pokemon
	}
	tests := []struct {
		name       string
		args       args
		wantResult entity.PokemonDB
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			args: args{
				data: entity.Pokemon{
					ID: 1,
				},
			},
			wantResult: entity.PokemonDB{
				ID:       1,
				Metadata: `{"image_url":"","description":"","weight":0,"height":0,"stats":{"hp":0,"attack":0,"def":0,"speed":0}}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pu := newBuildPokemonProviderMock(PokemonUsecase{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			})

			gotResult, err := pu.buildPokemonFromRequest(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.buildPokemonFromRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PokemonUsecase.buildPokemonFromRequest() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
