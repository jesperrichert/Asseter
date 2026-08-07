package repository

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T, id any) error {
	return db.Model(new(T)).
		Where("id = ?", id).
		Updates(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) Get(db *gorm.DB, entity *T, id any) error {
	return db.Where("id = ?", id).Take(entity).Error
}

func (Repository[T]) List(db *gorm.DB, entitis *[]T) error {
	return db.Find(entitis).Error
}

//GetPaged - Get all entries Paginated
