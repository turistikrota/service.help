package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/category"
)

type CategoryUpdateCmd struct {
	UUID string           `json:"-"`
	Meta *CategoryMetaReq `json:"meta" validate:"required,dive"`
}

type CategoryUpdateRes struct{}

type CategoryUpdateHandler cqrs.HandlerFunc[CategoryUpdateCmd, *CategoryUpdateRes]

func NewCategoryUpdateHandler(factory category.Factory, repo category.Repo) CategoryUpdateHandler {
	return func(ctx context.Context, fcc CategoryUpdateCmd) (*CategoryUpdateRes, *i18np.Error) {
		e := factory.New(category.NewConfig{
			TrMeta: fcc.Meta.TR,
			EnMeta: fcc.Meta.EN,
		})
		e.UUID = fcc.UUID
		err := repo.Update(ctx, e)
		if err != nil {
			return nil, err
		}
		return &CategoryUpdateRes{}, nil
	}
}
