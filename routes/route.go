package routes

import (
	"GinGonicGorm/controller"

	"github.com/gin-gonic/gin"
)

func Product(route *gin.Engine, productController controller.ProductController, authController controller.AuthController) {

	api := route.Group("/api/v1")

	product := api.Group("/product")
	{
		product.GET("/findAll", productController.FindAllProduct)
		product.POST("/createProduct", productController.CreateProduct)
		product.GET("/findProductById/:id", productController.FindProductById)
		product.GET("/findAllProduct", productController.FindAllProduct)
		product.PUT("/updatedProduct", productController.UpdatedProduct)
		product.GET("/paggingAndSearch", productController.PagginationProductWithFilter)
	}

	user := api.Group("/user")
	{
		user.POST("/register", authController.Register)
		user.POST("/signIn", authController.Login)
	}
}
