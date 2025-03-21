package routes

import (
	"GinGonicGorm/controller"
	"GinGonicGorm/middleware"

	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, productController controller.ProductController, authController controller.AuthController, categoryController controller.CategoryController) {

	api := route.Group("/api/v1")

	// --- Product ---
	product := api.Group("/product")
	// Tambahka Middleware
	product.Use(middleware.AuthMiddleware())
	{
		product.GET("/findAll", productController.FindAllProduct)
		product.POST("/createProduct", productController.CreateProduct)
		product.GET("/findProductById/:id", productController.FindProductById)
		product.GET("/findAllProduct", productController.FindAllProduct)
		product.PUT("/updatedProduct", productController.UpdatedProduct)
		product.GET("/paggingAndSearch", productController.PagginationProductWithFilter)
	}
	category := api.Group("/category")
	category.Use(middleware.AuthMiddleware())
	{
		category.POST("/createCategory", categoryController.CreateCategory)
	}
	// --- User ---
	user := api.Group("/user")

	{
		user.POST("/register", authController.Register)
		user.POST("/signIn", authController.Login)
	}

}
