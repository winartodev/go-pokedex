package pokemontyperepository

import (
	"context"
	"database/sql"
	"errors"
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

func TestNewPokemonTypeRepository(t *testing.T) {
	db, _ := NewMock()
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want PokemonTypeRepositoryItf
	}{
		{
			name: "success",
			args: args{
				db: db,
			},
			want: &PokemonTypeRepository{
				PokemonTypeDB: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPokemonTypeRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPokemonTypeRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonTypeRepository_CreatePokemonTypeDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := InsertPokemonTypeQuery
	pokemonType := entity.PokemonType{
		PokemonID: 1,
		TypeID:    2,
	}

	type fields struct {
		PokemonTypeDB *sql.DB
	}
	type args struct {
		ctx  context.Context
		data entity.PokemonType
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
				PokemonTypeDB: db,
			},
			args: args{
				ctx:  ctx,
				data: pokemonType,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(pokemonType.PokemonID, pokemonType.TypeID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonTypeDB: db,
			},
			args: args{
				ctx:  ctx,
				data: pokemonType,
			},
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(pokemonType.PokemonID, pokemonType.TypeID).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pt := &PokemonTypeRepository{
				PokemonTypeDB: tt.fields.PokemonTypeDB,
			}
			if err := pt.CreatePokemonTypeDB(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("PokemonTypeRepository.CreatePokemonTypeDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPokemonTypeRepository_GetPokemonTypeByPokemonIDDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := GetPokemonTypesByPokemonIDQuery
	id := 1
	pokemonType := []entity.PokemonType{
		{
			ID:        1,
			PokemonID: 1,
			TypeID:    2,
			Name:      "FIRE",
		},
	}

	type fields struct {
		PokemonTypeDB *sql.DB
	}
	type args struct {
		ctx       context.Context
		pokemonID int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []entity.PokemonType
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			fields: fields{
				PokemonTypeDB: db,
			},
			args: args{
				ctx:       ctx,
				pokemonID: 1,
			},
			wantResult: pokemonType,
			wantErr:    false,
			mock: func() {
				dbmock.ExpectQuery(query).WithArgs(id).WillReturnRows(
					sqlmock.NewRows([]string{"id", "pokemon_id", "types_id", "types.name"}).
						AddRow(pokemonType[0].ID, pokemonType[0].PokemonID, pokemonType[0].TypeID, pokemonType[0].Name))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonTypeDB: db,
			},
			args: args{
				ctx:       ctx,
				pokemonID: 1,
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				dbmock.ExpectQuery(query).WithArgs(id).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pt := &PokemonTypeRepository{
				PokemonTypeDB: tt.fields.PokemonTypeDB,
			}
			gotResult, err := pt.GetPokemonTypeByPokemonIDDB(tt.args.ctx, tt.args.pokemonID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonTypeRepository.GetPokemonTypeByPokemonIDDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PokemonTypeRepository.GetPokemonTypeByPokemonIDDB() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPokemonTypeRepository_UpdatePokemonTypeDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	id := 1
	query := UpdatePokemonTokenQuery
	pokemon := entity.PokemonType{
		PokemonID: 123,
		TypeID:    123,
	}

	type fields struct {
		PokemonTypeDB *sql.DB
	}
	type args struct {
		ctx  context.Context
		id   int64
		data entity.PokemonType
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
				PokemonTypeDB: db,
			},
			args: args{
				ctx:  ctx,
				id:   1,
				data: pokemon,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(pokemon.PokemonID, pokemon.TypeID, id).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonTypeDB: db,
			},
			args: args{
				ctx:  ctx,
				id:   1,
				data: pokemon,
			},
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pt := &PokemonTypeRepository{
				PokemonTypeDB: tt.fields.PokemonTypeDB,
			}
			if err := pt.UpdatePokemonTypeDB(tt.args.ctx, tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("PokemonTypeRepository.UpdatePokemonTypeDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPokemonTypeRepository_DeletePokemonTypeByPokemonIDDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := DeletePokemonTypeByPokemonIDQuery
	pokemonID := 1

	type fields struct {
		PokemonTypeDB *sql.DB
	}
	type args struct {
		ctx       context.Context
		pokemonID int64
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
				PokemonTypeDB: db,
			},
			args: args{
				ctx:       ctx,
				pokemonID: 1,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(query).WithArgs(pokemonID).WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: "failed",
			fields: fields{
				PokemonTypeDB: db,
			},
			args: args{
				ctx:       ctx,
				pokemonID: 1,
			},
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(query).WithArgs(pokemonID).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			pt := &PokemonTypeRepository{
				PokemonTypeDB: tt.fields.PokemonTypeDB,
			}
			if err := pt.DeletePokemonTypeByPokemonIDDB(tt.args.ctx, tt.args.pokemonID); (err != nil) != tt.wantErr {
				t.Errorf("PokemonTypeRepository.DeletePokemonTypeByPokemonIDDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPokemonTypeRepository_DeletePokemonTypeByIDDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := DeletePokemonTypeByIDQuery
	id := 1

	type fields struct {
		PokemonTypeDB *sql.DB
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
				PokemonTypeDB: db,
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
				PokemonTypeDB: db,
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
			pt := &PokemonTypeRepository{
				PokemonTypeDB: tt.fields.PokemonTypeDB,
			}
			if err := pt.DeletePokemonTypeByIDDB(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PokemonTypeRepository.DeletePokemonTypeByIDDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
