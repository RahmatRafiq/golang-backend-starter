package controllers

import (
	"net/http"

	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/app/services"
	"github.com/gin-gonic/gin"
)

type StoreController struct {
	service services.StoreService
}

func NewStoreController(service services.StoreService) *StoreController {
	return &StoreController{service: service}
}

// @Summary List all stores
// @Description Mengambil daftar semua toko
// @Tags stores
// @Security Bearer
// @Produce json
// @Success 200 {array} models.Store
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /stores [get]
func (c *StoreController) List(ctx *gin.Context) {
	stores, err := c.service.GetAllStores()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, stores)
}

// @Summary Get a store by ID
// @Description Mengambil satu toko berdasarkan ID
// @Tags stores
// @Security Bearer
// @Produce json
// @Param id path string true "Store ID"
// @Success 200 {object} models.Store
// @Failure 404 {object} map[string]string "Store not found"
// @Router /stores/{id} [get]
func (c *StoreController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	store, err := c.service.GetStoreByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}
	ctx.JSON(http.StatusOK, store)
}

// @Summary Create or edit a store
// @Description Membuat atau mengedit toko
// @Tags stores
// @Security Bearer
// @Produce json
// @Param store body models.Store true "Store Data"
// @Success 200 {object} models.Store
// @Router /stores [post]
func (c *StoreController) Put(ctx *gin.Context) {
	var store models.Store
	if err := ctx.ShouldBindJSON(&store); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gunakan PutStore yang menangani create/update
	updatedStore, err := c.service.PutStore(store)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedStore)
}

// @Summary Delete a store by ID
// @Description Menghapus toko berdasarkan ID
// @Tags stores
// @Security Bearer
// @Produce json
// @Param id path string true "Store ID"
// @Success 200 {object} map[string]string "Store deleted"
// @Failure 404 {object} map[string]string "Store not found"
// @Router /stores/{id} [delete]
func (c *StoreController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteStore(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Store deleted"})
}
