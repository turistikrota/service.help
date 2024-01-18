package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
)

type ArticleDeactivateCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type ArticleDeactivateRes struct{}

type ArticleDeactivateHandler cqrs.HandlerFunc[ArticleDeactivateCmd, *ArticleDeactivateRes]

func NewArticleDeactivateHandler(repo article.Repo) ArticleDeactivateHandler {
	return func(ctx context.Context, fac ArticleDeactivateCmd) (*ArticleDeactivateRes, *i18np.Error) {
		err := repo.Deactivate(ctx, fac.UUID)
		if err != nil {
			return nil, err
		}
		return &ArticleDeactivateRes{}, nil
	}
}
