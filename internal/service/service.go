package service

import (
	"context"

	"github.com/popolzen/pricepulse/internal/model"
	"github.com/popolzen/pricepulse/internal/storage"
)

type ItemService struct {
	repo storage.Storage // интерфейс, не конкретная реализация!
}

func NewItemService(repo storage.Storage) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) AddItem(ctx context.Context, url string, targetPrice int) (model.Item, error) {
	// TODO
	//  определить магазин по URL
	// scarapperom получить текущую цену
	// сохранить в репозиторий
	// оповестить при соблюжении условий
	item := model.Item{
		URL:         url,
		TargetPrice: targetPrice,
	}
	return s.repo.Add(ctx, item)
}

// func (s *ItemService) GetItems(ctx context.Context) ([]model.Item, error) {
// 	return s.repo.GetAll(ctx)
// }

// func (s *ItemService) DeleteItem(ctx context.Context, id int64) error {
// 	return s.repo.Delete(ctx, id)
// }
