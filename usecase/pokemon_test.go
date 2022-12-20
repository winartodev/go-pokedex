package usecase

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/winartodev/go-pokedex/entity"
	pokemonrepository "github.com/winartodev/go-pokedex/repository/pokemon"
	pokemonrepositorymock "github.com/winartodev/go-pokedex/repository/pokemon/mocks"
	pokemontyperepository "github.com/winartodev/go-pokedex/repository/pokemontypes"
	pokemontyperepositorymock "github.com/winartodev/go-pokedex/repository/pokemontypes/mocks"
)

type mockPokemonProvider struct {
	PokemonRepository     *pokemonrepositorymock.PokemonRepositoryItf
	PokemonTypeRepository *pokemontyperepositorymock.PokemonTypeRepositoryItf
}

func pokemonProvider() mockPokemonProvider {
	return mockPokemonProvider{
		PokemonRepository:     new(pokemonrepositorymock.PokemonRepositoryItf),
		PokemonTypeRepository: new(pokemontyperepositorymock.PokemonTypeRepositoryItf),
	}
}

func TestNewPokemonUsecase(t *testing.T) {
	pokemonUsecase := PokemonUsecase{
		PokemonRepository:     new(pokemonrepositorymock.PokemonRepositoryItf),
		PokemonTypeRepository: new(pokemontyperepositorymock.PokemonTypeRepositoryItf),
	}

	type args struct {
		pokemonUsecase PokemonUsecase
	}
	tests := []struct {
		name string
		args args
		want PokemonUsecaseItf
	}{
		{
			name: "success",
			args: args{
				pokemonUsecase: pokemonUsecase,
			},
			want: &pokemonUsecase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPokemonUsecase(tt.args.pokemonUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPokemonUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonUsecase_GetAllPokemon(t *testing.T) {
	ctx := context.Background()
	prov := pokemonProvider()

	type fields struct {
		PokemonRepository     pokemonrepository.PokemonRepositoryItf
		PokemonTypeRepository pokemontyperepository.PokemonTypeRepositoryItf
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResults []entity.PokemonList
		wantErr     bool
		mock        func()
	}{
		{
			name: "success",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
			},
			wantResults: []entity.PokemonList{{ID: 1}},
			wantErr:     false,
			mock: func() {
				prov.PokemonRepository.On("GetAllPokemonDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonDB{{ID: 1, Metadata: "{}"}}, nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1}}, nil).Times(1)
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
			},
			wantResults: nil,
			wantErr:     true,
			mock: func() {
				prov.PokemonRepository.On("GetAllPokemonDB", mock.Anything, mock.Anything).
					Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := &PokemonUsecase{
				PokemonRepository:     tt.fields.PokemonRepository,
				PokemonTypeRepository: tt.fields.PokemonTypeRepository,
			}

			gotResults, err := pu.GetAllPokemon(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.GetAllPokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("PokemonUsecase.GetAllPokemon() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestPokemonUsecase_GetAllPokemonByFilter(t *testing.T) {
	ctx := context.Background()
	prov := pokemonProvider()

	type fields struct {
		PokemonRepository     pokemonrepository.PokemonRepositoryItf
		PokemonTypeRepository pokemontyperepository.PokemonTypeRepositoryItf
	}
	type args struct {
		ctx    context.Context
		filter map[string]string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResults []entity.PokemonList
		wantErr     bool
		mock        func()
	}{
		{
			name: "success",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				filter: map[string]string{
					"name": "bulbasour",
				},
			},
			wantResults: []entity.PokemonList{{ID: 1}},
			wantErr:     false,
			mock: func() {
				prov.PokemonRepository.On("GetAllPokemonByFilterDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonDB{{ID: 1, Metadata: "{}"}}, nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1}}, nil).Times(1)
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				filter: map[string]string{
					"name": "bulbasour",
				},
			},
			wantResults: nil,
			wantErr:     true,
			mock: func() {
				prov.PokemonRepository.On("GetAllPokemonByFilterDB", mock.Anything, mock.Anything).
					Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := &PokemonUsecase{
				PokemonRepository:     tt.fields.PokemonRepository,
				PokemonTypeRepository: tt.fields.PokemonTypeRepository,
			}
			gotResults, err := pu.GetAllPokemonByFilter(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.GetAllPokemonByFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("PokemonUsecase.GetAllPokemonByFilter() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestPokemonUsecase_CreatePokemon(t *testing.T) {
	ctx := context.Background()
	prov := pokemonProvider()
	data := entity.Pokemon{
		Name:     "Bulbasour",
		Species:  "Pokemon",
		Types:    []int64{1, 2, 3},
		Catched:  0,
		ImageURL: "https://image.com/image/1",
	}

	type fields struct {
		PokemonRepository     pokemonrepository.PokemonRepositoryItf
		PokemonTypeRepository pokemontyperepository.PokemonTypeRepositoryItf
	}
	type args struct {
		ctx  context.Context
		data entity.Pokemon
	}
	tests := []struct {
		name          string
		args          args
		fields        fields
		wantPokemonID int64
		wantErr       bool
		mock          func()
	}{
		{
			name: "success",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx:  ctx,
				data: data,
			},
			wantPokemonID: 1,
			wantErr:       false,
			mock: func() {
				prov.PokemonRepository.On("CreatePokemonDB", mock.Anything, mock.Anything).
					Return(int64(1), nil).Times(1)

				prov.PokemonTypeRepository.On("CreatePokemonTypeDB", mock.Anything, mock.Anything).
					Return(nil).Times(3)
			},
		},
		{
			name: "failed create pokemonDB",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx:  ctx,
				data: data,
			},
			wantPokemonID: 0,
			wantErr:       true,
			mock: func() {
				prov.PokemonRepository.On("CreatePokemonDB", mock.Anything, mock.Anything).
					Return(int64(0), errors.New("errors")).Times(1)
			},
		},
		{
			name: "failed create pokemonDB",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx:  ctx,
				data: data,
			},
			wantPokemonID: 1,
			wantErr:       true,
			mock: func() {
				prov.PokemonRepository.On("CreatePokemonDB", mock.Anything, mock.Anything).
					Return(int64(1), nil).Times(1)

				prov.PokemonTypeRepository.On("CreatePokemonTypeDB", mock.Anything, mock.Anything).
					Return(errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := &PokemonUsecase{
				PokemonRepository:     tt.fields.PokemonRepository,
				PokemonTypeRepository: tt.fields.PokemonTypeRepository,
			}

			gotPokemonID, err := pu.CreatePokemon(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.CreatePokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPokemonID != tt.wantPokemonID {
				t.Errorf("PokemonUsecase.CreatePokemon() = %v, want %v", gotPokemonID, tt.wantPokemonID)
			}
		})
	}
}

func TestPokemonUsecase_GetPokemonByID(t *testing.T) {
	ctx := context.Background()
	prov := pokemonProvider()

	type fields struct {
		PokemonRepository     pokemonrepository.PokemonRepositoryItf
		PokemonTypeRepository pokemontyperepository.PokemonTypeRepositoryItf
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult *entity.PokemonDetail
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResult: &entity.PokemonDetail{ID: 1, Name: "bulbasour", Species: "pokemon", Catched: 0},
			wantErr:    false,
			mock: func() {
				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{ID: 1, Name: "bulbasour", Species: "pokemon", Catched: 0, Metadata: "{}"}, nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, Name: "FIRE"}}, nil).Times(1)
			},
		},
		{
			name: "failed no rows",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResult: nil,
			wantErr:    false,
			mock: func() {
				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{}, sql.ErrNoRows).Times(1)
			},
		},
		{
			name: "failed GetPokemonByIDDB",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{}, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := &PokemonUsecase{
				PokemonRepository:     tt.fields.PokemonRepository,
				PokemonTypeRepository: tt.fields.PokemonTypeRepository,
			}

			gotResult, err := pu.GetPokemonByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.GetPokemonByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PokemonUsecase.GetPokemonByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPokemonUsecase_UpdatePokemon(t *testing.T) {
	ctx := context.Background()
	prov := pokemonProvider()

	type fields struct {
		PokemonRepository     pokemonrepository.PokemonRepositoryItf
		PokemonTypeRepository pokemontyperepository.PokemonTypeRepositoryItf
	}
	type args struct {
		ctx  context.Context
		id   int64
		data entity.Pokemon
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantResult *entity.PokemonDetail
		wantErr    bool
		mock       func()
	}{
		{
			name: "success update pokemon",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
				data: entity.Pokemon{
					Name:    "Bulbasour",
					Species: "pokemon",
					Types:   []int64{1},
				},
			},
			wantResult: &entity.PokemonDetail{
				ID:      1,
				Name:    "Bulbasour",
				Species: "pokemon",
				Types:   []string{"FIRE"},
			},
			wantErr: false,
			mock: func() {
				prov.PokemonRepository.On("UpdatePokemonDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, PokemonID: 1, TypeID: 1, Name: "FIRE"}}, nil).Times(1)

				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{ID: 1, Name: "Bulbasour", Species: "pokemon", Catched: 0, Metadata: "{}"}, nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, PokemonID: 1, TypeID: 1, Name: "FIRE"}}, nil).Times(1)
			},
		},
		// case when length request is greather than or equal length data pokemon type that obtained from database.
		{
			name: "success update pokemon type",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
				data: entity.Pokemon{
					Name:    "Bulbasour",
					Species: "pokemon",
					Types:   []int64{2},
				},
			},
			wantResult: &entity.PokemonDetail{
				ID:      1,
				Name:    "Bulbasour",
				Species: "pokemon",
				Types:   []string{"WATER"},
			},
			wantErr: false,
			mock: func() {
				prov.PokemonRepository.On("UpdatePokemonDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, PokemonID: 1, TypeID: 1, Name: "FIRE"}}, nil).Times(1)

				prov.PokemonTypeRepository.On("UpdatePokemonTypeDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{ID: 1, Name: "Bulbasour", Species: "pokemon", Catched: 0, Metadata: "{}"}, nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, PokemonID: 1, TypeID: 2, Name: "WATER"}}, nil).Times(1)
			},
		},
		{
			name: "success update and create new pokemon type",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
				data: entity.Pokemon{
					Name:    "Bulbasour",
					Species: "pokemon",
					Types:   []int64{2, 3},
				},
			},
			wantResult: &entity.PokemonDetail{
				ID:      1,
				Name:    "Bulbasour",
				Species: "pokemon",
				Types:   []string{"WATER", "ICE"},
			},
			wantErr: false,
			mock: func() {
				prov.PokemonRepository.On("UpdatePokemonDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, PokemonID: 1, TypeID: 1, Name: "FIRE"}}, nil).Times(1)

				prov.PokemonTypeRepository.On("UpdatePokemonTypeDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonTypeRepository.On("CreatePokemonTypeDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{ID: 1, Name: "Bulbasour", Species: "pokemon", Catched: 0, Metadata: "{}"}, nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, PokemonID: 1, TypeID: 2, Name: "WATER"}, {ID: 1, PokemonID: 1, TypeID: 3, Name: "ICE"}}, nil).Times(1)
			},
		},
		// case when length request less than length data pokemon type that obtained from database.
		{
			name: "success update and create new pokemon type",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
				data: entity.Pokemon{
					Name:    "Bulbasour",
					Species: "pokemon",
					Types:   []int64{3},
				},
			},
			wantResult: &entity.PokemonDetail{
				ID:      1,
				Name:    "Bulbasour",
				Species: "pokemon",
				Types:   []string{"ICE"},
			},
			wantErr: false,
			mock: func() {
				prov.PokemonRepository.On("UpdatePokemonDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, PokemonID: 1, TypeID: 1, Name: "FIRE"}, {ID: 1, PokemonID: 1, TypeID: 2, Name: "WATER"}}, nil).Times(1)

				prov.PokemonTypeRepository.On("UpdatePokemonTypeDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(2)

				prov.PokemonTypeRepository.On("CreatePokemonTypeDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{ID: 1, Name: "Bulbasour", Species: "pokemon", Catched: 0, Metadata: "{}"}, nil).Times(1)

				prov.PokemonTypeRepository.On("GetPokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return([]entity.PokemonType{{ID: 1, PokemonID: 1, TypeID: 3, Name: "ICE"}}, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := &PokemonUsecase{
				PokemonRepository:     tt.fields.PokemonRepository,
				PokemonTypeRepository: tt.fields.PokemonTypeRepository,
			}

			gotResult, err := pu.UpdatePokemon(tt.args.ctx, tt.args.id, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.UpdatePokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PokemonUsecase.UpdatePokemon() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPokemonUsecase_DeletePokemon(t *testing.T) {
	ctx := context.Background()
	prov := pokemonProvider()

	type fields struct {
		PokemonRepository     pokemonrepository.PokemonRepositoryItf
		PokemonTypeRepository pokemontyperepository.PokemonTypeRepositoryItf
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success delete pokemon",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: false,
			mock: func() {
				prov.PokemonRepository.On("DeletePokemonByIDDB", mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonTypeRepository.On("DeletePokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return(nil).Times(1)
			},
		},
		{
			name: "failed delete pokemon",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				prov.PokemonRepository.On("DeletePokemonByIDDB", mock.Anything, mock.Anything).
					Return(errors.New("error")).Times(1)
			},
		},
		{
			name: "failed delete type pokemon",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				prov.PokemonRepository.On("DeletePokemonByIDDB", mock.Anything, mock.Anything).
					Return(nil).Times(1)

				prov.PokemonTypeRepository.On("DeletePokemonTypeByPokemonIDDB", mock.Anything, mock.Anything).
					Return(errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := &PokemonUsecase{
				PokemonRepository:     tt.fields.PokemonRepository,
				PokemonTypeRepository: tt.fields.PokemonTypeRepository,
			}

			if err := pu.DeletePokemon(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.DeletePokemon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPokemonUsecase_CatchPokemon(t *testing.T) {
	ctx := context.Background()
	prov := pokemonProvider()

	type fields struct {
		PokemonRepository     pokemonrepository.PokemonRepositoryItf
		PokemonTypeRepository pokemontyperepository.PokemonTypeRepositoryItf
	}
	type args struct {
		ctx context.Context
		id  int64
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
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: false,
			mock: func() {
				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{ID: 1, Name: "Bulbasour", Catched: 0}, nil).Times(1)

				prov.PokemonRepository.On("UpdatePokemonDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)
			},
		},
		{
			name: "failed get pokemon by id",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{}, errors.New("error")).Times(1)
			},
		},
		{
			name: "failed update pokemon",
			fields: fields{
				PokemonRepository:     prov.PokemonRepository,
				PokemonTypeRepository: prov.PokemonTypeRepository,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				prov.PokemonRepository.On("GetPokemonByIDDB", mock.Anything, mock.Anything).
					Return(entity.PokemonDB{ID: 1, Name: "Bulbasour", Catched: 0}, nil).Times(1)

				prov.PokemonRepository.On("UpdatePokemonDB", mock.Anything, mock.Anything, mock.Anything).
					Return(errors.New("errors")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pu := &PokemonUsecase{
				PokemonRepository:     tt.fields.PokemonRepository,
				PokemonTypeRepository: tt.fields.PokemonTypeRepository,
			}

			if err := pu.CatchPokemon(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PokemonUsecase.CatchPokemon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
