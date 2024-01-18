package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/faq"
)

type FaqFilterQuery struct {
	*faq.FilterEntity
}

type FaqFilterRes struct {
	List []faq.ListDto
}

type FaqFilterHandler cqrs.HandlerFunc[FaqFilterQuery, *FaqFilterRes]

func NewFaqFilterHandler(repo faq.Repo) FaqFilterHandler {
	return func(ctx context.Context, q FaqFilterQuery) (*FaqFilterRes, *i18np.Error) {
		list, err := repo.Filter(ctx, *q.FilterEntity)
		if err != nil {
			return nil, err
		}
		res := make([]faq.ListDto, len(list))
		for i, v := range list {
			res[i] = v.ToList()
		}
		return &FaqFilterRes{List: res}, nil
	}
}
