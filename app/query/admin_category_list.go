package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.help/domains/category"
)

type AdminCategoryListQuery struct {
}

type AdminCategoryListRes struct {
	List []category.AdminListDto
}

type AdminCategoryListHandler cqrs.HandlerFunc[AdminCategoryListQuery, *AdminCategoryListRes]

func NewAdminCategoryListHandler(repo category.Repo) AdminCategoryListHandler {
	return func(ctx context.Context, q AdminCategoryListQuery) (*AdminCategoryListRes, *i18np.Error) {
		list, err := repo.List(ctx)
		if err != nil {
			return nil, err
		}
		res := make([]category.AdminListDto, len(list))
		for i, v := range list {
			res[i] = v.ToAdminList()
		}
		return &AdminCategoryListRes{List: res}, nil
	}
}
