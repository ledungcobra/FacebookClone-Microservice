package dao

import (
	"gorm.io/gorm"
)

type CommonDao[T any] struct {
	db *gorm.DB
}

func (c CommonDao[T]) RawQuery(query string, args ...any) error {
	return c.db.Exec(query, args).Error
}

func (c CommonDao[T]) Save(object *T) error {
	return c.db.Save(object).Error
}

func (c CommonDao[T]) Delete(object any) error {
	return c.db.Delete(object).Error
}

func (c CommonDao[T]) Find(query any, args ...any) (*T, error) {
	var dest T
	if result := c.db.Where(query, args).First(&dest); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
	}
	return &dest, nil
}

func NewCommonDao[T any](db *gorm.DB) *CommonDao[T] {
	return &CommonDao[T]{db: db}
}
