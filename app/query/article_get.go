package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
)

type ArticleGetQuery struct {
	Locale string `params:"-"`
	Slug   string `params:"slug" validate:"required,slug"`
}

type ArticleGetRes struct {
	Detail article.DetailDto
}

type ArticleGetHandler cqrs.HandlerFunc[ArticleGetQuery, *ArticleGetRes]

func NewArticleGetHandler(repo article.Repo) ArticleGetHandler {
	return func(ctx context.Context, q ArticleGetQuery) (*ArticleGetRes, *i18np.Error) {
		res, notFound, err := repo.GetBySlug(ctx, q.Locale, q.Slug)
		if err != nil {
			return nil, err
		}
		if notFound {
			return nil, nil
		}
		return &ArticleGetRes{Detail: res.ToDetail()}, nil
	}
}
