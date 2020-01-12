package repository

import (
	"context"
	"github.com/colinfletch/goplaywithdb/model"
)

// PostRepo Interface outlining the methods to create within our Http Handler
type PostRepo interface {
	ViewAllPosts(ctx context.Context, num int32) ([]*model.Post, error)
	ViewPost(ctx context.Context, id int32) (*model.Post, error)
	CreatePost(ctx context.Context, p *model.Post) (int64, error)
	UpdatePost(ctx context.Context, p *model.Post) (*model.Post, error)
	DeletePost(ctx context.Context, id int32) (bool, error)
}
