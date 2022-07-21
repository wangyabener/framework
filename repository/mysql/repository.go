package mysql

import (
	"context"

	"gorm.io/gorm"
)

type ID interface{}

type GormModel[E any] interface {
	ToEntity() E
	FromEntity(entity E) interface{}
}

type Repository[M GormModel[E], E any] struct {
	db *gorm.DB
}

func NewRepostitory[M GormModel[E], E any](db *gorm.DB) *Repository[M, E] {
	return &Repository[M, E]{
		db: db,
	}
}

func (r *Repository[M, E]) Insert(ctx context.Context, entity *E) error {
	var start M
	model := start.FromEntity(*entity).(M)

	err := r.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		return err
	}

	*entity = model.ToEntity()
	return nil
}

func (r *Repository[M, E]) FindById(ctx context.Context, id ID) (E, error) {
	var model M

	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return *new(E), err
	}

	return model.ToEntity(), nil
}

func (r *Repository[M, E]) Find(ctx context.Context) ([]E, error) {
	var models []M
	// err := r.db.WithContext(ctx).Where().Find(&models).Error
	err := r.db.WithContext(ctx).Find(&models).Error
	if err != nil {
		return nil, err
	}

	result := make([]E, 0, len(models))
	for _, row := range models {
		result = append(result, row.ToEntity())
	}

	return result, nil
}
