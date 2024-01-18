package faq

type ListDto struct {
	UUID  string           `json:"uuid" bson:"_id,omitempty"`
	Meta  map[Locale]*Meta `json:"meta" bson:"meta" validate:"required,dive"`
	Order *int             `json:"order" bson:"order" validate:"required,min=0,max=100"`
}

type AdminListDto struct {
	*Entity
}

type AdminDetailDto struct {
	*Entity
}

func (e *Entity) ToList() ListDto {
	return ListDto{
		UUID:  e.UUID,
		Meta:  e.Meta,
		Order: e.Order,
	}
}

func (e *Entity) ToAdminList() AdminListDto {
	return AdminListDto{
		Entity: e,
	}
}

func (e *Entity) ToAdminDetail() AdminDetailDto {
	return AdminDetailDto{
		Entity: e,
	}
}
