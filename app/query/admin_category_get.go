package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/category"
)

type AdminCategoryGetQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type AdminCategoryGetRes struct {
	Detail category.AdminDetailDto
}

type AdminCategoryGetHandler cqrs.HandlerFunc[AdminCategoryGetQuery, *AdminCategoryGetRes]

func NewAdminCategoryGetHandler(factory category.Factory, repo category.Repo) AdminCategoryGetHandler {
	return func(ctx context.Context, q AdminCategoryGetQuery) (*AdminCategoryGetRes, *i18np.Error) {
		res, notFound, err := repo.GetByID(ctx, q.UUID)
		if err != nil {
			return nil, err
		}
		if notFound {
			return nil, nil
		}
		return &AdminCategoryGetRes{Detail: res.ToAdminDetail()}, nil
	}
}
