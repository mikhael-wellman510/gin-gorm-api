package routes

import (
	"GinGonicGorm/controller"

	"github.com/gin-gonic/gin"
)

func Product(route *gin.Engine, productController controller.ProductController) {

	routes := route.Group("/api/v1/product")
	{
		routes.GET("/findAll", productController.FindAllProduct)
		routes.POST("/createProduct", productController.CreateProduct)
		routes.GET("/findProductById/:id", productController.FindProductById)
		routes.GET("/findAllProduct", productController.FindAllProduct)
		routes.PUT("/updatedProduct", productController.UpdatedProduct)
		routes.GET("/paggingAndSearch", productController.PagginationProductWithFilter)
	}
}
