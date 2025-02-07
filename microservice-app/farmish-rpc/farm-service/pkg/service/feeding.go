package service

import (
	"farmish/model"
	"farmish/pkg/repository"
	"fmt"
	"log/slog"
)

type Feeding interface {
	FeedAnimal(model.CreateFeedingDTO) error
}

type FeedingService struct {
	repository          *repository.FeedingRepository
	warehouseRepository *repository.WarehouseRepository
}

func NewFeedingService(repository *repository.FeedingRepository, warehouseRepository *repository.WarehouseRepository) *FeedingService {
	return &FeedingService{
		repository,
		warehouseRepository,
	}
}

func (f *FeedingService) FeedAnimal(feed model.CreateFeedingDTO) error {
	stock, err := f.warehouseRepository.FindStockByName(feed.Type)
	if err != nil {
		slog.Error(fmt.Sprintf("feed-animal: %s", err.Error()))
		return fmt.Errorf("the feed %s is not registered", feed.Type)
	}

	if stock.Quantity < feed.Quantity {
		slog.Error(fmt.Sprintf("feed-animal: %s", "not enough quantity"))
		return fmt.Errorf("no enough quantity of feed %s", feed.Type)
	}

	return nil
}
