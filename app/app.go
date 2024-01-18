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
}

type Queries struct {
	FaqFilter query.FaqFilterHandler
}
