package article

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

type NewConfig struct{}

func (f Factory) New(cnf NewConfig) *Entity {
	return &Entity{}
}
