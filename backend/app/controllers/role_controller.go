package controllers

import (
	"net/http"

	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/app/services"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service services.RoleService
}

func NewRoleController(service services.RoleService) *RoleController {
	return &RoleController{service: service}
}

// @Summary List all roles
// @Description Mengambil daftar semua role
// @Tags roles
// @Security Bearer
// @Produce json
// @Success 200 {array} models.Role
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /roles [get]
func (c *RoleController) List(ctx *gin.Context) {
	roles, err := c.service.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, roles)
}

// @Summary Get a role by ID
// @Description Mengambil satu role berdasarkan ID
// @Tags roles
// @Security Bearer
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} models.Role
// @Failure 404 {object} map[string]string "Role not found"
// @Router /roles/{id} [get]
func (c *RoleController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	role, err := c.service.GetRoleByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	ctx.JSON(http.StatusOK, role)
}

// @Summary Create or edit a role
// @Description Membuat atau mengedit role
// @Tags roles
// @Security Bearer
// @Produce json
// @Param role body models.Role true "Role Data"
// @Success 200 {object} models.Role
// @Router /roles [put]
func (c *RoleController) Put(ctx *gin.Context) {
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create or update role
	upsertedRole, err := c.service.PutRole(role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If role is newly created, return 201 Created, otherwise 200 OK
	if upsertedRole.ID == 0 {
		ctx.JSON(http.StatusCreated, upsertedRole)
	} else {
		ctx.JSON(http.StatusOK, upsertedRole)
	}
}

// @Summary Delete a role by ID
// @Description Menghapus role berdasarkan ID
// @Tags roles
// @Security Bearer
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} map[string]string "Role deleted"
// @Failure 404 {object} map[string]string "Role not found"
// @Router /roles/{id} [delete]
func (c *RoleController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteRole(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}

// Struct to wrap the permissions array
type AssignPermissionsRequest struct {
	Permissions []uint `json:"permissions"`
}

// @Summary Assign permissions to a role
// @Description Menetapkan permission ke role
// @Tags roles
// @Security Bearer
// @Produce json
// @Param id path string true "Role ID"
// @Param body body AssignPermissionsRequest true "Permission IDs"
// @Success 200 {object} map[string]string "Permissions assigned to role"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /roles/{id}/permissions [post]
func (c *RoleController) AssignPermissions(ctx *gin.Context) {
	var req AssignPermissionsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleId := ctx.Param("id")
	err := c.service.AssignPermissionsToRole(roleId, req.Permissions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Permissions assigned to role"})
}

// @Summary Get permissions assigned to a role
// @Description Mengambil permission yang ditetapkan ke role
// @Tags roles
// @Security Bearer
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {array} models.Permission
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /roles/{id}/permissions [get]
func (c *RoleController) GetPermissions(ctx *gin.Context) {
	roleId := ctx.Param("id")
	permissions, err := c.service.GetPermissionsByRoleId(roleId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, permissions)
}
