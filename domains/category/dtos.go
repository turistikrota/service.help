package category

type ListDto struct {
	UUID  string           `json:"uuid" bson:"_id,omitempty"`
	Meta  map[Locale]*Meta `json:"meta" bson:"meta"`
	Order *int             `json:"order" bson:"order"`
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
