package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
)

type ArticleUpdateCmd struct {
	UUID       string          `json:"-"`
	Meta       *ArticleMetaReq `json:"meta" validate:"required,dive"`
	CategoryID string          `json:"categoryId" validate:"required,object_id"`
}

type ArticleUpdateRes struct{}

type ArticleUpdateHandler cqrs.HandlerFunc[ArticleUpdateCmd, *ArticleUpdateRes]

func NewArticleUpdateHandler(factory article.Factory, repo article.Repo) ArticleUpdateHandler {
	return func(ctx context.Context, fcc ArticleUpdateCmd) (*ArticleUpdateRes, *i18np.Error) {
		e := factory.New(article.NewConfig{
			TrMeta:     fcc.Meta.TR,
			EnMeta:     fcc.Meta.EN,
			CategoryID: fcc.CategoryID,
		})
		e.UUID = fcc.UUID
		err := repo.Update(ctx, e)
		if err != nil {
			return nil, err
		}
		return &ArticleUpdateRes{}, nil
	}
}
