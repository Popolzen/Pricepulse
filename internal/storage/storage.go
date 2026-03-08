package storage

import (
	"context"

	"github.com/popolzen/pricepulse/internal/model"
)

type Storage interface {
	Add(ctx context.Context, item model.Item) (model.Item, error)
	Get(ctx context.Context, item model.Item) (model.Item, error)
}
