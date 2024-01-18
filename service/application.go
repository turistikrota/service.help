package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.help/app"
	"github.com/turistikrota/service.help/app/command"
	"github.com/turistikrota/service.help/app/query"
	"github.com/turistikrota/service.help/config"
	"github.com/turistikrota/service.help/domains/faq"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Validator   *validation.Validator
	MongoDB     *mongo.DB
	CacheSrv    cache.Service
}

func NewApplication(cnf Config) app.Application {

	faqFactory := faq.NewFactory()
	faqRepo := faq.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Faq.Collection), faqFactory)

	return app.Application{
		Commands: app.Commands{
			FaqCreate: command.NewFaqCreateHandler(faqFactory, faqRepo),
		},
		Queries: app.Queries{
			FaqFilter: query.NewFaqFilterHandler(faqRepo),
		},
	}
}
