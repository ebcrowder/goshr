package db

import (
	"context"

	"github.com/ebcrowder/goshr/schema"
)

const keyRepository = "Repository"

type Repository interface {
	Insert(file *schema.File) (string, error)
	Delete(id int) error
	GetFiles() ([]schema.File, error)
}

func SetRepository(ctx context.Context, repository Repository) context.Context {
	return context.WithValue(ctx, keyRepository, repository)
}

func Insert(ctx context.Context, file *schema.File) (string, error) {
	return getRepository(ctx).Insert(file)
}

func Delete(ctx context.Context, id int) error {
	return getRepository(ctx).Delete(id)
}

func GetFiles(ctx context.Context) ([]schema.File, error) {
	return getRepository(ctx).GetFiles()
}

func getRepository(ctx context.Context) Repository {
	return ctx.Value(keyRepository).(Repository)
}
