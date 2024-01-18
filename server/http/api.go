package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.help/app/command"
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
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) FaqCreate(ctx *fiber.Ctx) error {
	cmd := command.FaqCreateCmd{}
	h.parseBody(ctx, &cmd)
	res, err := h.app.Commands.FaqCreate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) FaqUpdate(ctx *fiber.Ctx) error {
	detail := command.DetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.FaqUpdateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.FaqUpdate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}
