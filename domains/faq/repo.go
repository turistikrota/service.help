package faq

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo interface {
	Create(ctx context.Context, entity *Entity) *i18np.Error
	Update(ctx context.Context, entity *Entity) *i18np.Error
	GetByID(ctx context.Context, id string) (*Entity, bool, *i18np.Error)
	Activate(ctx context.Context, id string) *i18np.Error
	Deactivate(ctx context.Context, id string) *i18np.Error
	ReOrder(ctx context.Context, id string, order int) *i18np.Error
	Filter(ctx context.Context, filter FilterEntity) ([]*Entity, *i18np.Error)
	FilterAdmin(ctx context.Context, filter FilterEntity) ([]*Entity, *i18np.Error)
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func NewRepo(collection *mongo.Collection, factory Factory) Repo {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}

func (r *repo) Create(ctx context.Context, e *Entity) *i18np.Error {
	_, err := r.collection.InsertOne(ctx, e)
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) Update(ctx context.Context, e *Entity) *i18np.Error {
	id, err := mongo2.TransformId(e.UUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.Meta:      e.Meta,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) GetByID(ctx context.Context, id string) (*Entity, bool, *i18np.Error) {
	uuid, _err := mongo2.TransformId(id)
	if _err != nil {
		return nil, false, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: uuid,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, false, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, false, nil
	}
	return *e, true, nil
}

func (r *repo) Activate(ctx context.Context, id string) *i18np.Error {
	uuid, err := mongo2.TransformId(id)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: uuid,
		fields.IsActive: bson.M{
			"$ne": true,
		},
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsActive:  true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Deactivate(ctx context.Context, id string) *i18np.Error {
	uuid, err := mongo2.TransformId(id)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:     uuid,
		fields.IsActive: true,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsActive:  false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) ReOrder(ctx context.Context, id string, order int) *i18np.Error {
	uuid, err := mongo2.TransformId(id)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: uuid,
	}
	update := bson.M{
		"$set": bson.M{
			fields.Order:     order,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Filter(ctx context.Context, filter FilterEntity) ([]*Entity, *i18np.Error) {
	filters := r.filterToBson(filter, bson.M{
		fields.IsActive: true,
	})
	return r.helper.GetListFilter(ctx, filters, r.filterOptions())
}

func (r *repo) FilterAdmin(ctx context.Context, filter FilterEntity) ([]*Entity, *i18np.Error) {
	filters := r.filterToBson(filter)
	return r.helper.GetListFilter(ctx, filters, r.filterAdminOptions())
}

func (r *repo) filterOptions() *options.FindOptions {
	return options.Find().SetProjection(bson.M{
		fields.IsActive:  0,
		fields.UpdatedAt: 0,
		fields.CreatedAt: 0,
	}).SetSort(bson.M{
		fields.Order: 1,
	})
}

func (r *repo) filterAdminOptions() *options.FindOptions {
	return options.Find().SetSort(bson.M{
		fields.Order: 1,
	})
}
