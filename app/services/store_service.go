package services

import (
	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/facades"
	"gorm.io/gorm/clause"
)

type StoreService struct{}

func (*StoreService) GetAllStores() ([]models.Store, error) {
	var stores []models.Store
	if err := facades.DB.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func (*StoreService) GetStoreByID(id string) (models.Store, error) {
	var store models.Store
	if err := facades.DB.First(&store, id).Error; err != nil {
		return store, err
	}
	return store, nil
}

func (*StoreService) PutStore(store models.Store) (models.Store, error) {
	if err := facades.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "phone", "address", "city", "state", "country", "zip", "updated_at"}),
	}).Create(&store).Error; err != nil {
		return store, err
	}
	return store, nil
}

func (*StoreService) DeleteStore(id string) error {
	var store models.Store
	if err := facades.DB.First(&store, id).Error; err != nil {
		return err
	}
	return facades.DB.Delete(&store).Error
}
