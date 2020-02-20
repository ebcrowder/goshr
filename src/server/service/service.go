package service

import (
	"context"

	"github.com/ebcrowder/goshr/db"
	"github.com/ebcrowder/goshr/schema"
)

func Insert(ctx context.Context, file *schema.File) (string, error) {
	return db.Insert(ctx, file)
}

func Delete(ctx context.Context, id string) error {
	return db.Delete(ctx, id)
}

func GetFiles(ctx context.Context, id string) ([]interface{}, error) {
	return db.GetFiles(ctx, id)
}
