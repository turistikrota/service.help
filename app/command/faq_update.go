package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/faq"
)

type FaqUpdateCmd struct {
	UUID string      `json:"-"`
	Meta *FaqMetaReq `json:"meta" validate:"required,dive"`
}

type FaqUpdateRes struct{}

type FaqUpdateHandler cqrs.HandlerFunc[FaqUpdateCmd, *FaqUpdateRes]

func NewFaqUpdateHandler(factory faq.Factory, repo faq.Repo) FaqUpdateHandler {
	return func(ctx context.Context, fcc FaqUpdateCmd) (*FaqUpdateRes, *i18np.Error) {
		e := factory.New(faq.NewConfig{
			TrMeta: fcc.Meta.TR,
			EnMeta: fcc.Meta.EN,
		})
		e.UUID = fcc.UUID
		err := repo.Update(ctx, e)
		if err != nil {
			return nil, err
		}
		return &FaqUpdateRes{}, nil
	}
}
