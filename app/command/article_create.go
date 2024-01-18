package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
)

type ArticleMetaReq struct {
	TR *article.Meta `json:"tr" validate:"required,dive"`
	EN *article.Meta `json:"en" validate:"required,dive"`
}

type ArticleCreateCmd struct {
	CategoryID string          `json:"categoryId" validate:"required,object_id"`
	Meta       *ArticleMetaReq `json:"meta" validate:"required,dive"`
}

type ArticleCreateRes struct{}

type ArticleCreateHandler cqrs.HandlerFunc[ArticleCreateCmd, *ArticleCreateRes]

func NewArticleCreateHandler(factory article.Factory, repo article.Repo) ArticleCreateHandler {
	return func(ctx context.Context, acc ArticleCreateCmd) (*ArticleCreateRes, *i18np.Error) {
		err := repo.Create(ctx, factory.New(article.NewConfig{
			TrMeta:     acc.Meta.TR,
			EnMeta:     acc.Meta.EN,
			CategoryID: acc.CategoryID,
		}))
		if err != nil {
			return nil, err
		}
		return &ArticleCreateRes{}, nil
	}
}
