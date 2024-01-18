package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/faq"
)

type AdminFaqGetQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type AdminFaqGetRes struct {
	Detail faq.AdminDetailDto
}

type AdminFaqGetHandler cqrs.HandlerFunc[AdminFaqGetQuery, *AdminFaqGetRes]

func NewAdminFaqGetHandler(factory faq.Factory, repo faq.Repo) AdminFaqGetHandler {
	return func(ctx context.Context, q AdminFaqGetQuery) (*AdminFaqGetRes, *i18np.Error) {
		res, notFound, err := repo.GetByID(ctx, q.UUID)
		if err != nil {
			return nil, err
		}
		if notFound {
			return nil, factory.Errors.NotFound()
		}
		return &AdminFaqGetRes{Detail: res.ToAdminDetail()}, nil
	}
}
