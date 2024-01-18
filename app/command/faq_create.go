package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/faq"
)

type FaqMetaReq struct {
	TR *faq.Meta `json:"tr" validate:"required,dive"`
	EN *faq.Meta `json:"en" validate:"required,dive"`
}

type FaqCreateCmd struct {
	Meta *FaqMetaReq `json:"meta" validate:"required,dive"`
}

type FaqCreateRes struct{}

type FaqCreateHandler cqrs.HandlerFunc[FaqCreateCmd, *FaqCreateRes]

func NewFaqCreateHandler(factory faq.Factory, repo faq.Repo) FaqCreateHandler {
	return func(ctx context.Context, fcc FaqCreateCmd) (*FaqCreateRes, *i18np.Error) {
		err := repo.Create(ctx, factory.New(faq.NewConfig{
			TrMeta: fcc.Meta.TR,
			EnMeta: fcc.Meta.EN,
		}))
		if err != nil {
			return nil, err
		}
		return &FaqCreateRes{}, nil
	}
}
