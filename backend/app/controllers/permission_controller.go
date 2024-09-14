package controllers

import (
	"net/http"

	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/app/services"
	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	service services.PermissionService
}

func NewPermissionController(service services.PermissionService) *PermissionController {
	return &PermissionController{service: service}
}

// @Summary List all permissions
// @Description Mengambil daftar semua permission
// @Tags permissions
// @Security Bearer
// @Produce json
// @Success 200 {array} models.Permission
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /permissions [get]
func (c *PermissionController) List(ctx *gin.Context) {
	permissions, err := c.service.GetAllPermissions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, permissions)
}

// @Summary Get a permission by ID
// @Description Mengambil satu permission berdasarkan ID
// @Tags permissions
// @Security Bearer
// @Produce json
// @Param id path string true "Permission ID"
// @Success 200 {object} models.Permission
// @Failure 404 {object} map[string]string "Permission not found"
// @Router /permissions/{id} [get]
func (c *PermissionController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	permission, err := c.service.GetPermissionByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}
	ctx.JSON(http.StatusOK, permission)
}

// @Summary Create or edit a permission
// @Description Membuat atau mengedit permission
// @Tags permissions
// @Security Bearer
// @Produce json
// @Param permission body models.Permission true "Permission Data"
// @Success 200 {object} models.Permission
// @Router /permissions [put]
func (c *PermissionController) Put(ctx *gin.Context) {
	var permission models.Permission
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create or update permission
	upsertedPermission, err := c.service.PutPermission(permission)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If permission is newly created, return 201 Created, otherwise 200 OK
	if upsertedPermission.ID == 0 {
		ctx.JSON(http.StatusCreated, upsertedPermission)
	} else {
		ctx.JSON(http.StatusOK, upsertedPermission)
	}
}

// @Summary Delete a permission by ID
// @Description Menghapus permission berdasarkan ID
// @Tags permissions
// @Security Bearer
// @Produce json
// @Param id path string true "Permission ID"
// @Success 200 {object} map[string]string "Permission deleted"
// @Failure 404 {object} map[string]string "Permission not found"
// @Router /permissions/{id} [delete]
func (c *PermissionController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeletePermission(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Permission deleted"})
}
