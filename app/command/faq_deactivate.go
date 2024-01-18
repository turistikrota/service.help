package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/faq"
)

type FaqDeactivateCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type FaqDeactivateRes struct{}

type FaqDeactivateHandler cqrs.HandlerFunc[FaqDeactivateCmd, *FaqDeactivateRes]

func NewFaqDeactivateHandler(repo faq.Repo) FaqDeactivateHandler {
	return func(ctx context.Context, fac FaqDeactivateCmd) (*FaqDeactivateRes, *i18np.Error) {
		err := repo.Deactivate(ctx, fac.UUID)
		if err != nil {
			return nil, err
		}
		return &FaqDeactivateRes{}, nil
	}
}
