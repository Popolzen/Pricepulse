package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/popolzen/pricepulse/internal/model"
)

type MemoryRepo struct {
	items  map[int64]model.Item
	mu     sync.Mutex
	nextID int64
}

func (r *MemoryRepo) Add(ctx context.Context, item model.Item) (model.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nextID++
	item.ID = r.nextID
	r.items[item.ID] = item
	return item, nil
}

func (r *MemoryRepo) GetByID(ctx context.Context, id int64) (model.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, ok := r.items[id]
	if !ok {
		return model.Item{}, fmt.Errorf("товар %d не найден", id)
	}
	return item, nil
}
