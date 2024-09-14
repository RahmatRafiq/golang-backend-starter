package services

import (
	"errors"

	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/facades"
	"gorm.io/gorm/clause"
)

type RoleService struct{}

func (*RoleService) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := facades.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (*RoleService) GetRoleByID(id string) (models.Role, error) {
	var role models.Role
	if err := facades.DB.First(&role, id).Error; err != nil {
		return role, err
	}
	return role, nil
}

// PutRole: Handles both create and update operations
func (*RoleService) PutRole(role models.Role) (models.Role, error) {
	if err := facades.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // Resolves conflict based on ID
		DoUpdates: clause.AssignmentColumns([]string{"role"}),
	}).Create(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (*RoleService) DeleteRole(id string) error {
	var role models.Role
	if err := facades.DB.First(&role, id).Error; err != nil {
		return err
	}
	return facades.DB.Delete(&role).Error
}

// Assign permissions to a role
func (*RoleService) AssignPermissionsToRole(roleId string, permissions []uint) error {
	var role models.Role
	if err := facades.DB.First(&role, roleId).Error; err != nil {
		return err
	}

	// Validate permissions before assigning them
	var validPermissions []uint
	facades.DB.Table("permissions").Where("id IN ?", permissions).Pluck("id", &validPermissions)

	if len(validPermissions) != len(permissions) {
		return errors.New("one or more permission IDs are invalid")
	}

	// Clear existing permissions for the role
	facades.DB.Where("role_id = ?", role.ID).Delete(&models.RoleHasPermissions{})

	// Assign new permissions
	for _, permId := range validPermissions {
		rolePerm := models.RoleHasPermissions{
			RoleID:       role.ID,
			PermissionID: permId,
		}
		if err := facades.DB.Create(&rolePerm).Error; err != nil {
			return err
		}
	}

	return nil
}

func (*RoleService) GetPermissionsByRoleId(roleId string) ([]models.Permission, error) {
	var permissions []models.Permission
	if err := facades.DB.Table("permissions").
		Select("permissions.*").
		Joins("join role_has_permissions on permissions.id = role_has_permissions.permission_id").
		Where("role_has_permissions.role_id = ?", roleId).
		Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}
