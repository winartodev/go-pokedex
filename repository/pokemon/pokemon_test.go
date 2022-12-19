package pokemonrepository

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

func TestNewPokemonRepository(t *testing.T) {
	db, _ := NewMock()
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want PokemonRepositoryItf
	}{
		{
			name: "success",
			args: args{
				db: db,
			},
			want: &PokemonRepository{
				PokemonDB: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPokemonRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPokemonRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonRepository_GetAllPokemonDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := fmt.Sprintf("%s %s", GetPokemonQuery, `GROUP BY pokemons.id`)
	pokemon := []entity.PokemonDB{
		{
			ID:       1,
			Name:     "Bulbasour",
			Species:  "ganteng",
			Catched:  0,
			Metadata: "",
		},
	}

	type fields struct {
		PokemonDB *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResults []entity.PokemonDB
		wantErr     bool
		mock        func()
	}{
		{
			name: "success",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx: ctx,
			},
			wantResults: pokemon,
			wantErr:     false,
			mock: func() {
				dbmock.ExpectQuery(query).WillReturnRows(
					dbmock.NewRows([]string{"id", "name", "species", "catched", "metadata"}).
						AddRow(pokemon[0].ID, pokemon[0].Name, pokemon[0].Species, pokemon[0].Catched, pokemon[0].Metadata))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx: ctx,
			},
			wantResults: nil,
			wantErr:     true,
			mock: func() {
				dbmock.ExpectQuery(query).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pr := &PokemonRepository{
				PokemonDB: tt.fields.PokemonDB,
			}
			gotResults, err := pr.GetAllPokemonDB(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonRepository.GetAllPokemonDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("PokemonRepository.GetAllPokemonDB() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestPokemonRepository_CreatePokemonDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := InsertPokemonQuery
	pokemon := entity.PokemonDB{
		Name:     "Bulbasour",
		Species:  "ganteng",
		Catched:  0,
		Metadata: "{}",
	}

	type fields struct {
		PokemonDB *sql.DB
	}
	type args struct {
		ctx  context.Context
		data entity.PokemonDB
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
				PokemonDB: db,
			},
			args: args{
				ctx:  ctx,
				data: pokemon,
			},
			wantId:  1,
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(pokemon.Name, pokemon.Species, pokemon.Catched, pokemon.Metadata).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx:  ctx,
				data: pokemon,
			},
			wantId:  0,
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(pokemon.Name, pokemon.Species, pokemon.Catched, pokemon.Metadata).WillReturnError(errors.New("err"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pr := &PokemonRepository{
				PokemonDB: tt.fields.PokemonDB,
			}
			gotId, err := pr.CreatePokemonDB(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonRepository.CreatePokemonDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("PokemonRepository.CreatePokemonDB() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestPokemonRepository_GetPokemonByIDDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := fmt.Sprintf(`%s %s`, GetPokemonQuery, `WHERE pokemons.id = \\?`)
	pokemon := entity.PokemonDB{
		ID:       1,
		Name:     "Bulbasour",
		Species:  "ganteng",
		Catched:  0,
		Metadata: "",
	}

	type fields struct {
		PokemonDB *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult entity.PokemonDB
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResult: pokemon,
			wantErr:    false,
			mock: func() {
				dbmock.ExpectQuery(query).WithArgs(pokemon.ID).WillReturnRows(
					dbmock.NewRows([]string{"id", "name", "species", "catched", "metadata"}).
						AddRow(pokemon.ID, pokemon.Name, pokemon.Species, pokemon.Catched, pokemon.Metadata))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResult: entity.PokemonDB{},
			wantErr:    true,
			mock: func() {
				dbmock.ExpectQuery(query).WithArgs(pokemon.ID).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pr := &PokemonRepository{
				PokemonDB: tt.fields.PokemonDB,
			}
			gotResult, err := pr.GetPokemonByIDDB(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonRepository.GetPokemonByIDDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PokemonRepository.GetPokemonByIDDB() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPokemonRepository_UpdatePokemonDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := UpdatePokemonQuery
	id := 1

	pokemon := entity.PokemonDB{
		ID:       1,
		Name:     "Bulbasour",
		Species:  "ganteng",
		Catched:  0,
		Metadata: "",
	}

	type fields struct {
		PokemonDB *sql.DB
	}
	type args struct {
		ctx  context.Context
		id   int64
		data entity.PokemonDB
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
				PokemonDB: db,
			},
			args: args{
				ctx:  ctx,
				id:   1,
				data: pokemon,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(pokemon.Name, pokemon.Species, pokemon.Catched, pokemon.Metadata, id).WillReturnResult(sqlmock.NewResult(1, 0))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx:  ctx,
				id:   1,
				data: pokemon,
			},
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(pokemon.Name, pokemon.Species, pokemon.Catched, pokemon.Metadata, id).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pr := &PokemonRepository{
				PokemonDB: tt.fields.PokemonDB,
			}
			if err := pr.UpdatePokemonDB(tt.args.ctx, tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("PokemonRepository.UpdatePokemonDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPokemonRepository_DeletePokemonByIDDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := DeletePokemonQuery
	id := 1

	type fields struct {
		PokemonDB *sql.DB
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
				PokemonDB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(id).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pr := &PokemonRepository{
				PokemonDB: tt.fields.PokemonDB,
			}
			if err := pr.DeletePokemonByIDDB(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PokemonRepository.DeletePokemonByIDDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPokemonRepository_GetAllPokemonByFilterDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := GetPokemonQuery
	pokemon := []entity.PokemonDB{
		{
			ID:       1,
			Name:     "Bulbasour",
			Species:  "ganteng",
			Catched:  0,
			Metadata: "",
		},
	}

	type fields struct {
		PokemonDB *sql.DB
	}
	type args struct {
		ctx    context.Context
		filter map[string]string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantPokemons []entity.PokemonDB
		wantErr      bool
		mock         func()
	}{
		{
			name: "success",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx: ctx,
				filter: map[string]string{
					"name":     "Bulbasour",
					"options":  "1",
					"type":     "1,2,3",
					"sort_by":  "id",
					"order_by": "desc",
				},
			},
			wantPokemons: pokemon,
			wantErr:      false,
			mock: func() {
				dbmock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(
					dbmock.NewRows([]string{"id", "name", "species", "catched", "metadata"}).
						AddRow(pokemon[0].ID, pokemon[0].Name, pokemon[0].Species, pokemon[0].Catched, pokemon[0].Metadata))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonDB: db,
			},
			args: args{
				ctx: ctx,
				filter: map[string]string{
					"name":     "Bulbasour",
					"options":  "1",
					"type":     "1,2,3",
					"sort_by":  "id",
					"order_by": "desc",
				},
			},
			wantPokemons: nil,
			wantErr:      true,
			mock: func() {
				dbmock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pr := &PokemonRepository{
				PokemonDB: tt.fields.PokemonDB,
			}
			gotPokemons, err := pr.GetAllPokemonByFilterDB(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonRepository.GetAllPokemonByFilterDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPokemons, tt.wantPokemons) {
				t.Errorf("PokemonRepository.GetAllPokemonByFilterDB() = %v, want %v", gotPokemons, tt.wantPokemons)
			}
		})
	}
}
