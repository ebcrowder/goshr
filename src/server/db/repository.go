package db

import (
	"context"

	"github.com/ebcrowder/goshr/schema"
)

const keyRepository = "Repository"

type Repository interface {
	Close()
	Insert(todo *schema.File) (int, error)
	Delete(id int) error
	GetFiles() ([]schema.File, error)
}

func SetRepository(ctx context.Context, repository Repository) context.Context {
	return context.WithValue(ctx, keyRepository, repository)
}

func Close(ctx context.Context) {
	getRepository(ctx).Close()
}

func Insert(ctx context.Context, todo *schema.File) (int, error) {
	return getRepository(ctx).Insert(todo)
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
