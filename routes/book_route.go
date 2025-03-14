package routes

import (
	"GinGonicGorm/controller"

	"github.com/gin-gonic/gin"
)

func Book(route *gin.Engine, bookController controller.BookController) {

	routes := route.Group("/api/v1")
	{
		routes.GET("/book1", bookController.FindAllBook)
		routes.GET("/book2", bookController.FindBookFun)
	}

}
