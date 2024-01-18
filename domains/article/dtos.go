package article

type ListDto struct{}

type AdminListDto struct{}

type AdminDetailDto struct{}

func (e *Entity) ToList() ListDto {
	return ListDto{}
}

func (e *Entity) ToAdminList() AdminListDto {
	return AdminListDto{}
}

func (e *Entity) ToAdminDetail() AdminDetailDto {
	return AdminDetailDto{}
}
