package controller

import (
	"GinGonicGorm/dto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ProductController interface {
		CreateProduct(ctx *gin.Context)
		FindAllProduct(ctx *gin.Context)
		FindProductById(ctx *gin.Context)
		UpdatedProduct(ctx *gin.Context)
	}

	productController struct {
	}
)

func NewProductController() ProductController {

	return &productController{}
}

func (pc *productController) CreateProduct(ctx *gin.Context) {

	reqBody := dto.ProductRequest{}

	/*gunakan return untuk menyelesaikan program*/
	if err := ctx.ShouldBind(&reqBody); err != nil {
		/*
			ini akan masuk error [tipe data beda , ada field yg tidak di isi (required)]
		*/
		log.Println("Error : ", err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Hasil reqBody : ", reqBody)

}
func (pc *productController) FindAllProduct(ctx *gin.Context) {

}

func (pc *productController) FindProductById(ctx *gin.Context) {

}

func (pc *productController) UpdatedProduct(ctx *gin.Context) {

}
