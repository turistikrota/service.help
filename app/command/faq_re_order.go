package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/faq"
)

type FaqReOrderCmd struct {
	UUID  string `json:"-"`
	Order *int   `json:"order" validate:"required"`
}

type FaqReOrderRes struct{}

type FaqReOrderHandler cqrs.HandlerFunc[FaqReOrderCmd, *FaqReOrderRes]

func NewFaqReOrderHandler(repo faq.Repo) FaqReOrderHandler {
	return func(ctx context.Context, fcc FaqReOrderCmd) (*FaqReOrderRes, *i18np.Error) {
		err := repo.ReOrder(ctx, fcc.UUID, *fcc.Order)
		if err != nil {
			return nil, err
		}
		return &FaqReOrderRes{}, nil
	}
}
