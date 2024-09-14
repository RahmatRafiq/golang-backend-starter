package controllers

import (
	"net/http"

	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/app/services"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// @Summary List all categories
// @Description Mengambil daftar semua kategori
// @Tags categories
// @Security Bearer
// @Produce json
// @Success 200 {array} models.Category
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /categories [get]
func (c *CategoryController) List(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// @Summary Get a category by ID
// @Description Mengambil satu kategori berdasarkan ID
// @Tags categories
// @Security Bearer
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]string "Category not found"
// @Router /categories/{id} [get]
func (c *CategoryController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	category, err := c.service.GetCategoryByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// @Summary Create or edit a category
// @Description Membuat atau mengedit kategori
// @Tags categories
// @Security Bearer
// @Produce json
// @Param category body models.Category true "Category Data"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]string "Bad request"
// @Router /categories [put]
func (c *CategoryController) Put(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := c.service.PutCategory(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory)
}

// @Summary Delete a category by ID
// @Description Menghapus kategori berdasarkan ID
// @Tags categories
// @Security Bearer
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} map[string]string "Category deleted"
// @Failure 404 {object} map[string]string "Category not found"
// @Router /categories/{id} [delete]
func (c *CategoryController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteCategory(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
