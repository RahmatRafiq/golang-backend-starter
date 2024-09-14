package controllers

import (
	"net/http"

	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/app/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// @Summary List all products
// @Description Mengambil daftar semua produk
// @Tags products
// @Security Bearer
// @Produce json
// @Success 200 {array} models.Product
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /products [get]
func (c *ProductController) List(ctx *gin.Context) {
	products, err := c.service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// @Summary Get a product by ID
// @Description Mengambil satu produk berdasarkan ID
// @Tags products
// @Security Bearer
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string "Product not found"
// @Router /products/{id} [get]
func (c *ProductController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := c.service.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// @Summary Create or edit a product
// @Description Membuat atau mengedit produk
// @Tags products
// @Security Bearer
// @Produce json
// @Param product body models.Product true "Product Data"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string "Product not found"
// @Router /products [put]
func (c *ProductController) Put(ctx *gin.Context) {
	var product models.Product

	// Bind request data to the product model
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create or update product
	upsertedProduct, err := c.service.PutProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If product is newly created, return 201 Created, otherwise 200 OK
	if upsertedProduct.ID == 0 {
		ctx.JSON(http.StatusCreated, upsertedProduct)
	} else {
		ctx.JSON(http.StatusOK, upsertedProduct)
	}
}

// @Summary Delete a product by ID
// @Description Menghapus produk berdasarkan ID
// @Tags products
// @Security Bearer
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} map[string]string "Product deleted"
// @Failure 404 {object} map[string]string "Product not found"
// @Router /products/{id} [delete]
func (c *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteProduct(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
