package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
)

type ArticleReOrderCmd struct {
	UUID  string `json:"-"`
	Order *int   `json:"order" validate:"required"`
}

type ArticleReOrderRes struct{}

type ArticleReOrderHandler cqrs.HandlerFunc[ArticleReOrderCmd, *ArticleReOrderRes]

func NewArticleReOrderHandler(repo article.Repo) ArticleReOrderHandler {
	return func(ctx context.Context, fcc ArticleReOrderCmd) (*ArticleReOrderRes, *i18np.Error) {
		err := repo.ReOrder(ctx, fcc.UUID, *fcc.Order)
		if err != nil {
			return nil, err
		}
		return &ArticleReOrderRes{}, nil
	}
}
