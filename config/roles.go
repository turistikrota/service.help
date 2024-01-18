package config

import "github.com/turistikrota/service.shared/base_roles"

type helpRoles struct {
	FaqCreate     string
	FaqUpdate     string
	FaqDeactivate string
	FaqActivate   string
	FaqReOrder    string
	FaqList       string
	FaqRead       string

	CategoryCreate     string
	CategoryUpdate     string
	CategoryDeactivate string
	CategoryActivate   string
	CategoryReOrder    string
	CategoryList       string
	CategoryRead       string

	ArticleCreate     string
	ArticleUpdate     string
	ArticleDeactivate string
	ArticleActivate   string
	ArticleReOrder    string
	ArticleList       string
	ArticleRead       string
}

type roles struct {
	base_roles.Roles
	Help helpRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Help: helpRoles{
		FaqCreate:     "help.faq.create",
		FaqUpdate:     "help.faq.update",
		FaqDeactivate: "help.faq.deactivate",
		FaqActivate:   "help.faq.activate",
		FaqReOrder:    "help.faq.reorder",
		FaqList:       "help.faq.list",
		FaqRead:       "help.faq.read",

		CategoryCreate:     "help.category.create",
		CategoryUpdate:     "help.category.update",
		CategoryDeactivate: "help.category.deactivate",
		CategoryActivate:   "help.category.activate",
		CategoryReOrder:    "help.category.reorder",
		CategoryList:       "help.category.list",
		CategoryRead:       "help.category.read",

		ArticleCreate:     "help.article.create",
		ArticleUpdate:     "help.article.update",
		ArticleDeactivate: "help.article.deactivate",
		ArticleActivate:   "help.article.activate",
		ArticleReOrder:    "help.article.reorder",
		ArticleList:       "help.article.list",
		ArticleRead:       "help.article.read",
	},
}
