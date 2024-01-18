package article

import "time"

type Entity struct {
	UUID       string           `json:"uuid" bson:"_id,omitempty"`
	CategoryID string           `json:"categoryId" bson:"category_id" validate:"required"`
	Meta       map[Locale]*Meta `json:"meta" bson:"meta" validate:"required,dive"`
	Order      *int             `json:"order" bson:"order" validate:"required,min=0,max=100"`
	IsActive   bool             `json:"isActive" bson:"is_active" validate:"required,boolean"`
	CreatedAt  time.Time        `json:"createdAt" bson:"created_at" validate:"required"`
	UpdatedAt  time.Time        `json:"updatedAt" bson:"updated_at" validate:"required"`
}

type Meta struct {
	Title    string `json:"title" bson:"title" validate:"required,max=100"`
	Slug     string `json:"slug" bson:"slug" validate:"omitempty"`
	Keywords string `json:"keywords" bson:"keywords" validate:"required,max=100"`
}

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleTR Locale = "tr"
)

func (l Locale) String() string {
	return string(l)
}
