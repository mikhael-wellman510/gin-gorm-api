package controller

import "github.com/gin-gonic/gin"

// Interface & Harus di implementasikan
type BookController interface {
	FindAllBook(ctx *gin.Context)
	FindBookFun(ctx *gin.Context)
}

// Atributes
type bookController struct{}

func NewBookController() BookController {

	return &bookController{}
}

func (b *bookController) FindAllBook(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "List of books",
	})
}

func (b *bookController) FindBookFun(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Book Fun",
	})
}
