package command

type DetailCmd struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}
