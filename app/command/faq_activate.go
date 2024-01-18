package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/faq"
)

type FaqActivateCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type FaqActivateRes struct{}

type FaqActivateHandler cqrs.HandlerFunc[FaqActivateCmd, *FaqActivateRes]

func NewFaqActivateHandler(repo faq.Repo) FaqActivateHandler {
	return func(ctx context.Context, fac FaqActivateCmd) (*FaqActivateRes, *i18np.Error) {
		err := repo.Activate(ctx, fac.UUID)
		if err != nil {
			return nil, err
		}
		return &FaqActivateRes{}, nil
	}
}
