package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
)

type AdminArticleGetQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type AdminArticleGetRes struct {
	Detail article.AdminDetailDto
}

type AdminArticleGetHandler cqrs.HandlerFunc[AdminArticleGetQuery, *AdminArticleGetRes]

func NewAdminArticleGetHandler(factory article.Factory, repo article.Repo) AdminArticleGetHandler {
	return func(ctx context.Context, q AdminArticleGetQuery) (*AdminArticleGetRes, *i18np.Error) {
		res, notFound, err := repo.GetByID(ctx, q.UUID)
		if err != nil {
			return nil, err
		}
		if notFound {
			return nil, factory.Errors.NotFound()
		}
		return &AdminArticleGetRes{Detail: res.ToAdminDetail()}, nil
	}
}
