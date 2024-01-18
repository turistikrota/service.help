package command
import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
)

type ArticleActivateCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type ArticleActivateRes struct{}

type ArticleActivateHandler cqrs.HandlerFunc[ArticleActivateCmd, *ArticleActivateRes]

func NewArticleActivateHandler(repo article.Repo) ArticleActivateHandler {
	return func(ctx context.Context, fac ArticleActivateCmd) (*ArticleActivateRes, *i18np.Error) {
		err := repo.Activate(ctx, fac.UUID)
		if err != nil {
			return nil, err
		}
		return &ArticleActivateRes{}, nil
	}
}
