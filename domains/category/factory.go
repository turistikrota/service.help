package category

import "time"

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
	EnMeta *Meta
	TrMeta *Meta
}

func (f Factory) New(cnf NewConfig) *Entity {
	order := 0
	meta := map[Locale]*Meta{
		LocaleEN: cnf.EnMeta,
		LocaleTR: cnf.TrMeta,
	}
	t := time.Now()
	return &Entity{
		Meta:      meta,
		Order:     &order,
		IsActive:  true,
		CreatedAt: t,
		UpdatedAt: t,
	}
}
