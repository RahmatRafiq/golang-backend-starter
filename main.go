package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RahmatRafiq/golang_backend_starter/facades"
	"github.com/RahmatRafiq/golang_backend_starter/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/RahmatRafiq/golang_backend_starter/docs"
)

func main() {
	// Memuat variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inisialisasi koneksi facades
	facades.ConnectDB()

	// Menginisialisasi Gin Router
	route := gin.Default()

	// Endpoint untuk mengecek kesehatan koneksi facades
	route.GET("/health", func(c *gin.Context) {
		sqlDB, err := facades.DB.DB() // Mengambil facades/sql *DB dari GORM *DB
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get facades connection",
				"error":   err.Error(),
			})
			return
		}

		err = sqlDB.Ping() // Menggunakan sqlDB untuk ping ke facades
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "facades connection failed",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "facades is connected",
			"facades": "supply_chain_retail", // Sesuaikan dengan nama facades Anda
		})
	})

	// Setup Swagger documentation
	docs.SwaggerInfo.Title = "Golang Backend Starter"
	docs.SwaggerInfo.Description = "This is a sample server for Golang Backend Starter With Auth."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	// Daftarkan routes
	routes.RegisterRoutes(route)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Server is running on port 8080")
	route.Run(":8080")
}
