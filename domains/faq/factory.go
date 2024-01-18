package faq

import (
	"time"

	"github.com/ssibrahimbas/slug"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewConfig struct {
	TrMeta *Meta
	EnMeta *Meta
}

func (f Factory) New(cnf NewConfig) *Entity {
	order := 0
	t := time.Now()
	cnf.TrMeta.Slug = slug.New(cnf.TrMeta.Title, slug.TR)
	cnf.EnMeta.Slug = slug.New(cnf.EnMeta.Title, slug.EN)
	meta := map[Locale]*Meta{
		LocaleTR: cnf.TrMeta,
		LocaleEN: cnf.EnMeta,
	}
	return &Entity{
		Meta:      meta,
		Order:     &order,
		IsActive:  true,
		CreatedAt: t,
		UpdatedAt: t,
	}
}
