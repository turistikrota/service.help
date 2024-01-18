package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/category"
)

type CategoryDeactivateCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type CategoryDeactivateRes struct{}

type CategoryDeactivateHandler cqrs.HandlerFunc[CategoryDeactivateCmd, *CategoryDeactivateRes]

func NewCategoryDeactivateHandler(repo category.Repo) CategoryDeactivateHandler {
	return func(ctx context.Context, fac CategoryDeactivateCmd) (*CategoryDeactivateRes, *i18np.Error) {
		err := repo.Deactivate(ctx, fac.UUID)
		if err != nil {
			return nil, err
		}
		return &CategoryDeactivateRes{}, nil
	}
}
