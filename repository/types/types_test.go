package typesrepository

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

func TestNewTypeRepository(t *testing.T) {
	db, _ := NewMock()
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want TypeRepositoryItf
	}{
		{
			name: "success",
			args: args{
				db: db,
			},
			want: &TypeRepository{
				TypeDB: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTypeRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTypeRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeRepository_CreateTypeDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := InsertTypeQuery
	typeData := entity.Type{
		Name: "FIRE",
	}

	type fields struct {
		TypeDB *sql.DB
	}
	type args struct {
		ctx  context.Context
		data entity.Type
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
				TypeDB: db,
			},
			args: args{
				ctx:  ctx,
				data: typeData,
			},
			wantId:  1,
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(typeData.Name).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "failed",
			fields: fields{
				TypeDB: db,
			},
			args: args{
				ctx:  ctx,
				data: typeData,
			},
			wantId:  0,
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(typeData.Name).
					WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TypeRepository{
				TypeDB: tt.fields.TypeDB,
			}
			gotId, err := tr.CreateTypeDB(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeRepository.CreateTypeDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("TypeRepository.CreateTypeDB() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestTypeRepository_GetAllTypeDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	query := GetTypesQuery
	typeData := []entity.Type{
		{Name: "FIRE"},
	}

	type fields struct {
		TypeDB *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResults []entity.Type
		wantErr     bool
		mock        func()
	}{
		{
			name: "success",
			fields: fields{
				TypeDB: db,
			},
			args: args{
				ctx: ctx,
			},
			wantResults: typeData,
			wantErr:     false,
			mock: func() {
				dbmock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
					AddRow(typeData[0].ID, typeData[0].Name))
			},
		},
		{
			name: "failed",
			fields: fields{
				TypeDB: db,
			},
			args: args{
				ctx: ctx,
			},
			wantResults: nil,
			wantErr:     true,
			mock: func() {
				dbmock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TypeRepository{
				TypeDB: tt.fields.TypeDB,
			}
			gotResults, err := tr.GetAllTypeDB(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeRepository.GetAllTypeDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("TypeRepository.GetAllTypeDB() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestTypeRepository_GeTypeByIDDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	id := 1
	query := fmt.Sprintf(`%s %s`, GetTypesQuery, `WHERE id = ?`)
	typeData := entity.Type{
		Name: "FIRE",
	}

	type fields struct {
		TypeDB *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult entity.Type
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			fields: fields{
				TypeDB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResult: typeData,
			wantErr:    false,
			mock: func() {
				dbmock.ExpectQuery(query).WithArgs(id).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name"}).AddRow(typeData.ID, typeData.Name),
				)
			},
		},
		{
			name: "failed",
			fields: fields{
				TypeDB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResult: entity.Type{},
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
			tr := &TypeRepository{
				TypeDB: tt.fields.TypeDB,
			}
			gotResult, err := tr.GeTypeByIDDB(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeRepository.GeTypeByIDDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TypeRepository.GeTypeByIDDB() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTypeRepository_UpdateTypeDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()
	id := 1
	query := UpdateTypeQuery
	typeData := entity.Type{
		Name: "FIRE",
	}

	type fields struct {
		TypeDB *sql.DB
	}
	type args struct {
		ctx  context.Context
		id   int64
		data entity.Type
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
				TypeDB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
				data: entity.Type{
					Name: "FIRE",
				},
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(typeData.Name, id).WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: "failed",
			fields: fields{
				TypeDB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
				data: entity.Type{
					Name: "FIRE",
				},
			},
			wantErr: true,
			mock: func() {
				dbmock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(typeData.Name, id).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TypeRepository{
				TypeDB: tt.fields.TypeDB,
			}
			if err := tr.UpdateTypeDB(tt.args.ctx, tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("TypeRepository.UpdateTypeDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
