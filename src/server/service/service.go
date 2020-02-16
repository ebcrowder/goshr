package service

import (
	"context"

	"github.com/ebcrowder/goshr/db"
	"github.com/ebcrowder/goshr/schema"
)

func Close(ctx context.Context) {
	db.Close(ctx)
}

func Insert(ctx context.Context, file *schema.File) (int, error) {
	return db.Insert(ctx, file)
}

func Delete(ctx context.Context, id int) error {
	return db.Delete(ctx, id)
}

func GetFiles(ctx context.Context) ([]schema.File, error) {
	return db.GetFiles(ctx)
}