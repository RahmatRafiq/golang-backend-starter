package services

import (
	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/facades"
	"gorm.io/gorm/clause"
)

type ProductService struct{}

func (*ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := facades.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (*ProductService) GetProductByID(id string) (models.Product, error) {
	var product models.Product
	if err := facades.DB.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

// PutProduct: Method to handle both create and update
func (*ProductService) PutProduct(newProduct models.Product) (models.Product, error) {
	if err := facades.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // Conflict resolution based on ID
		DoUpdates: clause.AssignmentColumns([]string{"reference", "store_id", "category_id", "name", "description", "price", "margin", "stock", "sold", "images", "updated_at"}),
	}).Create(&newProduct).Error; err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (*ProductService) DeleteProduct(id string) error {
	var product models.Product
	if err := facades.DB.First(&product, id).Error; err != nil {
		return err
	}
	return facades.DB.Delete(&product).Error
}
