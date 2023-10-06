package repo

import (
	"app/db_gorm_advance/model"

	"golang.org/x/net/context"
)

type PageRepository interface {
	GetOne(ctx context.Context, ID int) (model.Page, error)
	GetAll(ctx context.Context) ([]model.Page, error)
	Create(ctx context.Context, pages []model.Page) error
}
