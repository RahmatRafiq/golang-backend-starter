package routes

import (
	"github.com/RahmatRafiq/golang_backend_starter/app/controllers"
	"github.com/RahmatRafiq/golang_backend_starter/app/middleware"
	"github.com/RahmatRafiq/golang_backend_starter/app/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	// Apply middleware logging untuk semua route
	route.Use(middleware.LoggerMiddleware())

	// Public route: Hello World
	controller := controllers.Controller{}
	route.GET("/", controller.HelloWorld)

	// Public route: Login and Logout (no auth required)
	authService := services.AuthService{}
	authController := controllers.NewAuthController(authService)
	authRoutes := route.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", middleware.AuthMiddleware(), authController.Logout)
	}

	// Routes untuk stores (protected by AuthMiddleware)
	storeService := services.StoreService{}
	storeController := controllers.NewStoreController(storeService)
	storeRoutes := route.Group("/stores", middleware.AuthMiddleware()) // Protect store routes
	{
		storeRoutes.GET("/", storeController.List)         // List stores
		storeRoutes.GET("/:id", storeController.Get)       // Show/Edit store (GET by ID)
		storeRoutes.PUT("/", storeController.Put)          // Create/Update store
		storeRoutes.DELETE("/:id", storeController.Delete) // Delete store by ID
	}

	// Routes untuk categories (protected by AuthMiddleware)
	categoryService := services.CategoryService{}
	categoryController := controllers.NewCategoryController(categoryService)
	categoryRoutes := route.Group("/categories", middleware.AuthMiddleware()) // Protect category routes
	{
		categoryRoutes.GET("/", categoryController.List)         // List categories
		categoryRoutes.GET("/:id", categoryController.Get)       // Show/Edit category (GET by ID)
		categoryRoutes.PUT("/", categoryController.Put)          // Create/Update category
		categoryRoutes.DELETE("/:id", categoryController.Delete) // Delete category by ID
	}

	// Routes untuk products (protected by AuthMiddleware)
	productService := services.ProductService{}
	productController := controllers.NewProductController(productService)
	productRoutes := route.Group("/products", middleware.AuthMiddleware()) // Protect product routes
	{
		productRoutes.GET("/", productController.List)         // List all products
		productRoutes.GET("/:id", productController.Get)       // Show/Edit product by ID
		productRoutes.PUT("/", productController.Put)          // Create/Update product
		productRoutes.DELETE("/:id", productController.Delete) // Delete product by ID
	}

	// Routes untuk users (protected by AuthMiddleware)
	userService := services.UserService{}
	userController := controllers.NewUserController(userService)
	userRoutes := route.Group("/users", middleware.AuthMiddleware()) // Protect user routes
	{
		userRoutes.GET("/", userController.List)
		userRoutes.GET("/:id", userController.Get)
		userRoutes.PUT("/", userController.Put)
		userRoutes.DELETE("/:id", userController.Delete)
		userRoutes.POST("/:id/roles", userController.AssignRoles)
		userRoutes.GET("/:id/roles", userController.GetRoles)
	}

	// Routes untuk roles (protected by AuthMiddleware)
	roleService := services.RoleService{}
	roleController := controllers.NewRoleController(roleService)
	roleRoutes := route.Group("/roles", middleware.AuthMiddleware()) // Protect role routes
	{
		roleRoutes.GET("/", roleController.List)                              // List roles
		roleRoutes.GET("/:id", roleController.Get)                            // Get role by ID
		roleRoutes.PUT("/", roleController.Put)                               // Create/Update role
		roleRoutes.DELETE("/:id", roleController.Delete)                      // Delete role by ID
		roleRoutes.POST("/:id/permissions", roleController.AssignPermissions) // Assign permissions to role
		roleRoutes.GET("/:id/permissions", roleController.GetPermissions)     // Get permissions for role
	}

	// Routes untuk permissions (protected by AuthMiddleware)
	permissionService := services.PermissionService{}
	permissionController := controllers.NewPermissionController(permissionService)
	permissionRoutes := route.Group("/permissions", middleware.AuthMiddleware()) // Protect permission routes
	{
		permissionRoutes.GET("/", permissionController.List)         // List all permissions
		permissionRoutes.GET("/:id", permissionController.Get)       // Get permission by ID
		permissionRoutes.PUT("/", permissionController.Put)          // Create/Update permission
		permissionRoutes.DELETE("/:id", permissionController.Delete) // Delete permission by ID
	}
}
