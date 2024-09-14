package services

import (
	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/facades"
	"gorm.io/gorm/clause"
)

type PermissionService struct{}

func (*PermissionService) GetAllPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	if err := facades.DB.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (*PermissionService) GetPermissionByID(id string) (models.Permission, error) {
	var permission models.Permission
	if err := facades.DB.First(&permission, id).Error; err != nil {
		return permission, err
	}
	return permission, nil
}

// PutPermission: Method to handle both create and update
func (*PermissionService) PutPermission(permission models.Permission) (models.Permission, error) {
	if err := facades.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // Resolve conflict based on ID
		DoUpdates: clause.AssignmentColumns([]string{"permission"}),
	}).Create(&permission).Error; err != nil {
		return permission, err
	}
	return permission, nil
}

func (*PermissionService) DeletePermission(id string) error {
	var permission models.Permission
	if err := facades.DB.First(&permission, id).Error; err != nil {
		return err
	}
	return facades.DB.Delete(&permission).Error
}
