package http

import (
	"net/http"

	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.help/app/command"
	"github.com/turistikrota/service.help/app/query"
	"github.com/turistikrota/service.help/domains/article"
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

func (h srv) FaqAdminFilter(ctx *fiber.Ctx) error {
	filter := faq.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.AdminFaqFilterQuery{
		FilterEntity: &filter,
	}
	res, err := h.app.Queries.AdminFaqFilter(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) FaqAdminGet(ctx *fiber.Ctx) error {
	query := query.AdminFaqGetQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.AdminFaqGet(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	if res == nil {
		return result.ErrorDetail(Messages.Error.NotFound, map[string]interface{}{}, http.StatusNotFound)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
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

func (h srv) FaqReOrder(ctx *fiber.Ctx) error {
	detail := command.DetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.FaqReOrderCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.FaqReOrder(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) FaqActivate(ctx *fiber.Ctx) error {
	cmd := command.FaqActivateCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.FaqActivate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) FaqDeactivate(ctx *fiber.Ctx) error {
	cmd := command.FaqDeactivateCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.FaqDeactivate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ArticleFilter(ctx *fiber.Ctx) error {
	filter := article.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.ArticleFilterQuery{
		FilterEntity: &filter,
	}
	res, err := h.app.Queries.ArticleFilter(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) ArticleAdminFilter(ctx *fiber.Ctx) error {
	filter := article.FilterEntity{}
	h.parseQuery(ctx, &filter)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	filter.Locale = l
	query := query.AdminArticleFilterQuery{
		FilterEntity: &filter,
	}
	res, err := h.app.Queries.AdminArticleFilter(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) ArticleAdminGet(ctx *fiber.Ctx) error {
	query := query.AdminArticleGetQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.AdminArticleGet(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	if res == nil {
		return result.ErrorDetail(Messages.Error.NotFound, map[string]interface{}{}, http.StatusNotFound)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) ArticleCreate(ctx *fiber.Ctx) error {
	cmd := command.ArticleCreateCmd{}
	h.parseBody(ctx, &cmd)
	res, err := h.app.Commands.ArticleCreate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ArticleUpdate(ctx *fiber.Ctx) error {
	detail := command.DetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.ArticleUpdateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.ArticleUpdate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ArticleReOrder(ctx *fiber.Ctx) error {
	detail := command.DetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.ArticleReOrderCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.ArticleReOrder(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ArticleActivate(ctx *fiber.Ctx) error {
	cmd := command.ArticleActivateCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.ArticleActivate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ArticleDeactivate(ctx *fiber.Ctx) error {
	cmd := command.ArticleDeactivateCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.ArticleDeactivate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) CategoryAdminList(ctx *fiber.Ctx) error {
	res, err := h.app.Queries.AdminCategoryList(ctx.UserContext(), query.AdminCategoryListQuery{})
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) CategoryAdminGet(ctx *fiber.Ctx) error {
	query := query.AdminCategoryGetQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.AdminCategoryGet(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	if res == nil {
		return result.ErrorDetail(Messages.Error.NotFound, map[string]interface{}{}, http.StatusNotFound)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) CategoryCreate(ctx *fiber.Ctx) error {
	cmd := command.CategoryCreateCmd{}
	h.parseBody(ctx, &cmd)
	res, err := h.app.Commands.CategoryCreate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) CategoryUpdate(ctx *fiber.Ctx) error {
	detail := command.DetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.CategoryUpdateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.CategoryUpdate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) CategoryReOrder(ctx *fiber.Ctx) error {
	detail := command.DetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.CategoryReOrderCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.CategoryReOrder(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) CategoryActivate(ctx *fiber.Ctx) error {
	cmd := command.CategoryActivateCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.CategoryActivate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) CategoryDeactivate(ctx *fiber.Ctx) error {
	cmd := command.CategoryDeactivateCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.CategoryDeactivate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}
