package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.help/app/query"
	"github.com/turistikrota/service.help/domains/faq"
)

func (h srv) FaqFilter(ctx *fiber.Ctx) error {
	filter := faq.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.FaqFilterQuery{
		FilterEntity: &filter,
	}
	res, err := h.app.Queries.FaqFilter(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}
