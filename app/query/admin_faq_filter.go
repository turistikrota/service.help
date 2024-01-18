package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/faq"
)

type AdminFaqFilterQuery struct {
	*faq.FilterEntity
}

type AdminFaqFilterRes struct {
	List []faq.AdminListDto
}

type AdminFaqFilterHandler cqrs.HandlerFunc[AdminFaqFilterQuery, *AdminFaqFilterRes]

func NewAdminFaqFilterHandler(repo faq.Repo) AdminFaqFilterHandler {
	return func(ctx context.Context, q AdminFaqFilterQuery) (*AdminFaqFilterRes, *i18np.Error) {
		list, err := repo.Filter(ctx, *q.FilterEntity)
		if err != nil {
			return nil, err
		}
		res := make([]faq.AdminListDto, len(list))
		for i, v := range list {
			res[i] = v.ToAdminList()
		}
		return &AdminFaqFilterRes{List: res}, nil
	}
}
