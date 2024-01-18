package faq

import "go.mongodb.org/mongo-driver/bson"

type FilterEntity struct {
	Locale string `query:"-"`
	Query  string `query:"q,omitempty" validate:"omitempty,max=100"`
}

func (r *repo) filterToBson(filter FilterEntity, defaultFilters ...bson.M) bson.M {
	list := make([]bson.M, 0)
	if len(defaultFilters) > 0 {
		list = append(list, defaultFilters...)
	}
	list = r.filterByQuery(list, filter)
	listLen := len(list)
	if listLen == 0 {
		return bson.M{}
	}
	if listLen == 1 {
		return list[0]
	}
	return bson.M{
		"$and": list,
	}
}

func (r *repo) filterByQuery(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Query != "" {
		list = append(list, bson.M{
			"$or": []bson.M{
				{
					metaField(filter.Locale, metaFields.Title): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					metaField(filter.Locale, metaFields.Description): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					metaField(filter.Locale, metaFields.Keywords): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}
