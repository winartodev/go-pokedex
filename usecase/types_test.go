package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/winartodev/go-pokedex/entity"
	typesrepository "github.com/winartodev/go-pokedex/repository/types"
	typesrepositorymock "github.com/winartodev/go-pokedex/repository/types/mocks"
)

type mockTypeProvider struct {
	TypesRepository *typesrepositorymock.TypeRepositoryItf
}

func typeProvider() mockTypeProvider {
	return mockTypeProvider{
		TypesRepository: new(typesrepositorymock.TypeRepositoryItf),
	}
}

func TestNewTypeUsecase(t *testing.T) {
	typeRepository := TypeUsecase{
		TypesRepository: new(typesrepositorymock.TypeRepositoryItf),
	}

	type args struct {
		typeUsecase TypeUsecase
	}
	tests := []struct {
		name string
		args args
		want TypeUsecaseItf
	}{
		{
			name: "success",
			args: args{
				typeUsecase: typeRepository,
			},
			want: &typeRepository,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTypeUsecase(tt.args.typeUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTypeUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeUsecase_CreateType(t *testing.T) {
	ctx := context.Background()
	prov := typeProvider()

	type fields struct {
		TypesRepository typesrepository.TypeRepositoryItf
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
				TypesRepository: prov.TypesRepository,
			},
			args: args{
				ctx: ctx,
				data: entity.Type{
					Name: "FIRE",
				},
			},
			wantId:  1,
			wantErr: false,
			mock: func() {
				prov.TypesRepository.On("CreateTypeDB", mock.Anything, mock.Anything).
					Return(int64(1), nil).Times(1)
			},
		},
		{
			name: "failed",
			fields: fields{
				TypesRepository: prov.TypesRepository,
			},
			args: args{
				ctx: ctx,
				data: entity.Type{
					Name: "FIRE",
				},
			},
			wantId:  0,
			wantErr: true,
			mock: func() {
				prov.TypesRepository.On("CreateTypeDB", mock.Anything, mock.Anything).
					Return(int64(0), errors.New("errors")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TypeUsecase{
				TypesRepository: tt.fields.TypesRepository,
			}

			gotId, err := tr.CreateType(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeUsecase.CreateType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("TypeUsecase.CreateType() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestTypeUsecase_GetAllType(t *testing.T) {
	ctx := context.Background()
	prov := typeProvider()

	type fields struct {
		TypesRepository typesrepository.TypeRepositoryItf
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
				TypesRepository: prov.TypesRepository,
			},
			args: args{
				ctx: ctx,
			},
			wantResults: []entity.Type{{ID: 1, Name: "FIRE"}},
			wantErr:     false,
			mock: func() {
				prov.TypesRepository.On("GetAllTypeDB", mock.Anything, mock.Anything).
					Return([]entity.Type{{ID: 1, Name: "FIRE"}}, nil).Times(1)
			},
		},
		{
			name: "failed",
			fields: fields{
				TypesRepository: prov.TypesRepository,
			},
			args: args{
				ctx: ctx,
			},
			wantResults: nil,
			wantErr:     true,
			mock: func() {
				prov.TypesRepository.On("GetAllTypeDB", mock.Anything, mock.Anything).
					Return(nil, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TypeUsecase{
				TypesRepository: tt.fields.TypesRepository,
			}
			gotResults, err := tr.GetAllType(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeUsecase.GetAllType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("TypeUsecase.GetAllType() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestTypeUsecase_GeTypeByID(t *testing.T) {
	ctx := context.Background()
	prov := typeProvider()

	type fields struct {
		TypesRepository typesrepository.TypeRepositoryItf
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
				TypesRepository: prov.TypesRepository,
			},
			args: args{
				ctx: ctx,
			},
			wantResult: entity.Type{ID: 1, Name: "FIRE"},
			wantErr:    false,
			mock: func() {
				prov.TypesRepository.On("GeTypeByIDDB", mock.Anything, mock.Anything).
					Return(entity.Type{ID: 1, Name: "FIRE"}, nil).Times(1)
			},
		},
		{
			name: "failed",
			fields: fields{
				TypesRepository: prov.TypesRepository,
			},
			args: args{
				ctx: ctx,
			},
			wantResult: entity.Type{},
			wantErr:    true,
			mock: func() {
				prov.TypesRepository.On("GeTypeByIDDB", mock.Anything, mock.Anything).
					Return(entity.Type{}, errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TypeUsecase{
				TypesRepository: tt.fields.TypesRepository,
			}
			gotResult, err := tr.GeTypeByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeUsecase.GeTypeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TypeUsecase.GeTypeByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTypeUsecase_UpdateType(t *testing.T) {
	ctx := context.Background()
	prov := typeProvider()

	type fields struct {
		TypesRepository typesrepository.TypeRepositoryItf
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
				TypesRepository: prov.TypesRepository,
			},
			args: args{
				ctx:  ctx,
				id:   1,
				data: entity.Type{ID: 1, Name: "FIRE"},
			},
			wantErr: false,
			mock: func() {
				prov.TypesRepository.On("UpdateTypeDB", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Times(1)
			},
		},
		{
			name: "failed",
			fields: fields{
				TypesRepository: prov.TypesRepository,
			},
			args: args{
				ctx:  ctx,
				id:   1,
				data: entity.Type{ID: 1, Name: "FIRE"},
			},
			wantErr: true,
			mock: func() {
				prov.TypesRepository.On("UpdateTypeDB", mock.Anything, mock.Anything, mock.Anything).
					Return(errors.New("error")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TypeUsecase{
				TypesRepository: tt.fields.TypesRepository,
			}
			if err := tr.UpdateType(tt.args.ctx, tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("TypeUsecase.UpdateType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
