package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/category"
)

type CategoryReOrderCmd struct {
	UUID  string `json:"-"`
	Order *int   `json:"order" validate:"required"`
}

type CategoryReOrderRes struct{}

type CategoryReOrderHandler cqrs.HandlerFunc[CategoryReOrderCmd, *CategoryReOrderRes]

func NewCategoryReOrderHandler(repo category.Repo) CategoryReOrderHandler {
	return func(ctx context.Context, fcc CategoryReOrderCmd) (*CategoryReOrderRes, *i18np.Error) {
		err := repo.ReOrder(ctx, fcc.UUID, *fcc.Order)
		if err != nil {
			return nil, err
		}
		return &CategoryReOrderRes{}, nil
	}
}
