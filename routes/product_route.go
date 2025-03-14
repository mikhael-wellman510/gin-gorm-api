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
	}
}
