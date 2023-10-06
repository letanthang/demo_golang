package mysql

import (
	"app/db_gorm_advance/model"
	"app/db_gorm_advance/repo"
	"context"

	"gorm.io/gorm"
)

type pageRepository struct {
	db *gorm.DB
}

func NewPageRepository(db *gorm.DB) repo.PageRepository {
	return &pageRepository{
		db: db,
	}
}

func (r *pageRepository) GetOne(ctx context.Context, id int) (model.Page, error) {
	var page model.Page
	err := r.db.WithContext(ctx).Where("ID = ?", id).First(&page).Error

	return page, err
}

func (r *pageRepository) GetAll(ctx context.Context) ([]model.Page, error) {
	var pages []model.Page
	err := r.db.WithContext(ctx).Find(&pages).Error
	return pages, err
}

func (r *pageRepository) Create(ctx context.Context, pages []model.Page) error {
	err := r.db.WithContext(ctx).Create(pages).Error
	return err
}
