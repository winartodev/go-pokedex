package usecase

import (
	"context"

	"github.com/winartodev/go-pokedex/entity"
	typesrepository "github.com/winartodev/go-pokedex/repository/types"
)

type TypeUsecase struct {
	TypesRepository typesrepository.TypeRepositoryItf
}

type TypeUsecaseItf interface {
	CreateType(ctx context.Context, data entity.Type) (id int64, err error)
	GetAllType(ctx context.Context) (results []entity.Type, err error)
	GeTypeByID(ctx context.Context, id int64) (result entity.Type, err error)
	UpdateType(ctx context.Context, id int64, data entity.Type) (err error)
}

func NewTypeUsecase(typeUsecase TypeUsecase) TypeUsecaseItf {
	return &TypeUsecase{
		TypesRepository: typeUsecase.TypesRepository,
	}
}

func (tr *TypeUsecase) CreateType(ctx context.Context, data entity.Type) (id int64, err error) {
	id, err = tr.TypesRepository.CreateTypeDB(ctx, data)
	if err != nil {
		return id, err
	}

	return id, err
}

func (tr *TypeUsecase) GetAllType(ctx context.Context) (results []entity.Type, err error) {
	results, err = tr.TypesRepository.GetAllTypeDB(ctx)
	if err != nil {
		return results, err
	}

	return results, err
}

func (tr *TypeUsecase) GeTypeByID(ctx context.Context, id int64) (result entity.Type, err error) {
	result, err = tr.TypesRepository.GeTypeByIDDB(ctx, id)
	if err != nil {
		return result, err
	}

	return result, err
}

func (tr *TypeUsecase) UpdateType(ctx context.Context, id int64, data entity.Type) (err error) {
	err = tr.TypesRepository.UpdateTypeDB(ctx, id, data)
	if err != nil {
		return err
	}

	return err
}
