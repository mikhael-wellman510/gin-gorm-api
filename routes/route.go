package routes

import (
	"GinGonicGorm/controller"

	"GinGonicGorm/middleware"

	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, productController controller.ProductController, authController controller.AuthController, categoryController controller.CategoryController, photoController controller.PhotoController, testingController controller.TestingController) {

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

	// --- Photo ---
	photo := api.Group("/photo")
	{
		photo.POST("/upload", photoController.UploadPhoto)
		photo.GET("/photos/:filename", photoController.GetPhoto)
		// photo.GET("/photo")
	}

	test := api.Group("/testing")
	{
		test.POST("/ai", testingController.Testing)
	}

	// Yg kiri -> Path uRL yg akan di akses di browser
	// Yg kanan -> lokasi file server
	route.Static("/uploads", "./uploads")

}
