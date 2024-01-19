package article

type ListDto struct {
	UUID       string           `json:"uuid" bson:"_id,omitempty"`
	CategoryID string           `json:"categoryId" bson:"category_id"`
	Meta       map[Locale]*Meta `json:"meta" bson:"meta"`
	Order      *int             `json:"order" bson:"order"`
}

type DetailDto struct {
	UUID       string           `json:"uuid" bson:"_id,omitempty"`
	CategoryID string           `json:"categoryId" bson:"category_id"`
	Meta       map[Locale]*Meta `json:"meta" bson:"meta"`
}

type AdminListDto struct {
	*Entity
}

type AdminDetailDto struct {
	*Entity
}

func (e *Entity) ToList() ListDto {
	return ListDto{
		UUID:       e.UUID,
		CategoryID: e.CategoryID,
		Meta:       e.Meta,
		Order:      e.Order,
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

func (e *Entity) ToDetail() DetailDto {
	return DetailDto{
		UUID:       e.UUID,
		CategoryID: e.CategoryID,
		Meta:       e.Meta,
	}
}
