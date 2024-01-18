package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/article"
	"github.com/turistikrota/service.help/domains/category"
)

type FilterItem struct {
	Category category.ListDto
	Articles []article.ListDto
}

type ArticleFilterQuery struct {
	*article.FilterEntity
}

type ArticleFilterRes struct {
	List []FilterItem
}

type ArticleFilterHandler cqrs.HandlerFunc[ArticleFilterQuery, *ArticleFilterRes]

func NewArticleFilterHandler(articleRepo article.Repo, categoryRepo category.Repo) ArticleFilterHandler {
	return func(ctx context.Context, q ArticleFilterQuery) (*ArticleFilterRes, *i18np.Error) {
		articles, err := articleRepo.Filter(ctx, *q.FilterEntity)
		if err != nil {
			return nil, err
		}
		categories, err := categoryRepo.List(ctx)
		if err != nil {
			return nil, err
		}
		res := make([]FilterItem, len(categories))
		for i, v := range categories {
			res[i].Category = v.ToList()
			for _, v2 := range articles {
				if v2.CategoryID == v.UUID {
					res[i].Articles = append(res[i].Articles, v2.ToList())
				}
			}
		}
		return &ArticleFilterRes{List: res}, nil
	}
}
