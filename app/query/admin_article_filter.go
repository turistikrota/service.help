package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
)

type AdminArticleFilterQuery struct {
	*article.FilterEntity
}

type AdminArticleFilterRes struct {
	List []article.AdminListDto
}

type AdminArticleFilterHandler cqrs.HandlerFunc[AdminArticleFilterQuery, *AdminArticleFilterRes]

func NewAdminArticleFilterHandler(repo article.Repo) AdminArticleFilterHandler {
	return func(ctx context.Context, q AdminArticleFilterQuery) (*AdminArticleFilterRes, *i18np.Error) {
		list, err := repo.Filter(ctx, *q.FilterEntity)
		if err != nil {
			return nil, err
		}
		res := make([]article.AdminListDto, len(list))
		for i, v := range list {
			res[i] = v.ToAdminList()
		}
		return &AdminArticleFilterRes{List: res}, nil
	}
}
