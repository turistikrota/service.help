package app

import "github.com/turistikrota/service.help/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct{}

type Queries struct {
	FaqFilter query.FaqFilterHandler
}
