package repository

import (
	"context"
	"github.com/colinfletch/goplaywithdb/model"
)

// PostRepo Interface outlining the methods to create within our Http Handler
type PostRepo interface {
	Fetch(ctx context.Context, num int64) ([]*model.Post, error)
	GetByID(ctx context.Context, id int64) (*model.Post, error)
	Create(ctx context.Context, p *model.Post) (int64, error)
	Update(ctx context.Context, p *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
