package category

type fieldsType struct {
	UUID      string
	Meta      string
	Order     string
	IsActive  string
	CreatedAt string
	UpdatedAt string
}

type metaFieldsType struct {
	Title string
}

var fields = fieldsType{
	UUID:      "_id",
	Meta:      "meta",
	Order:     "order",
	IsActive:  "is_active",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var metaFields = metaFieldsType{
	Title: "title",
}

func metaField(locale string, field string) string {
	return fields.Meta + "." + locale + "." + field
}
