package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/category"
)

type CategoryActivateCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type CategoryActivateRes struct{}

type CategoryActivateHandler cqrs.HandlerFunc[CategoryActivateCmd, *CategoryActivateRes]

func NewCategoryActivateHandler(repo category.Repo) CategoryActivateHandler {
	return func(ctx context.Context, fac CategoryActivateCmd) (*CategoryActivateRes, *i18np.Error) {
		err := repo.Activate(ctx, fac.UUID)
		if err != nil {
			return nil, err
		}
		return &CategoryActivateRes{}, nil
	}
}
