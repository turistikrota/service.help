package app

import (
	"github.com/turistikrota/service.help/app/command"
	"github.com/turistikrota/service.help/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	FaqCreate     command.FaqCreateHandler
	FaqUpdate     command.FaqUpdateHandler
	FaqActivate   command.FaqActivateHandler
	FaqDeactivate command.FaqDeactivateHandler
	FaqReOrder    command.FaqReOrderHandler

	ArticleCreate     command.ArticleCreateHandler
	ArticleUpdate     command.ArticleUpdateHandler
	ArticleActivate   command.ArticleActivateHandler
	ArticleDeactivate command.ArticleDeactivateHandler
	ArticleReOrder    command.ArticleReOrderHandler

	CategoryCreate     command.CategoryCreateHandler
	CategoryUpdate     command.CategoryUpdateHandler
	CategoryActivate   command.CategoryActivateHandler
	CategoryDeactivate command.CategoryDeactivateHandler
	CategoryReOrder    command.CategoryReOrderHandler
}

type Queries struct {
	FaqFilter     query.FaqFilterHandler
	ArticleFilter query.ArticleFilterHandler

	AdminArticleFilter query.AdminArticleFilterHandler
	AdminArticleGet    query.AdminArticleGetHandler
	AdminCategoryList  query.AdminCategoryListHandler
	AdminCategoryGet   query.AdminCategoryGetHandler
	AdminFaqFilter     query.AdminFaqFilterHandler
	AdminFaqGet        query.AdminFaqGetHandler
}
