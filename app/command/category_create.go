package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/category"
)

type CategoryMetaReq struct {
	TR *category.Meta `json:"tr" validate:"required,dive"`
	EN *category.Meta `json:"en" validate:"required,dive"`
}

type CategoryCreateCmd struct {
	Meta *CategoryMetaReq `json:"meta" validate:"required,dive"`
}

type CategoryCreateRes struct {
	UUID string `json:"uuid"`
}

type CategoryCreateHandler cqrs.HandlerFunc[CategoryCreateCmd, *CategoryCreateRes]

func NewCategoryCreateHandler(factory category.Factory, repo category.Repo) CategoryCreateHandler {
	return func(ctx context.Context, acc CategoryCreateCmd) (*CategoryCreateRes, *i18np.Error) {
		res, err := repo.Create(ctx, factory.New(category.NewConfig{
			TrMeta: acc.Meta.TR,
			EnMeta: acc.Meta.EN,
		}))
		if err != nil {
			return nil, err
		}
		return &CategoryCreateRes{
			UUID: res.UUID,
		}, nil
	}
}
