package controller

import (
	"GinGonicGorm/dto"
	"GinGonicGorm/service"
	"GinGonicGorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	CategoryController interface {
		CreateCategory(ctx *gin.Context)
	}

	categoryController struct {
		categoryService service.CategoryService
	}
)

func NewCategoryService(categoryService service.CategoryService) CategoryController {

	return &categoryController{
		categoryService: categoryService,
	}
}

func (cc *categoryController) CreateCategory(ctx *gin.Context) {

	reqBody := dto.CategoryRequest{}

	if err := ctx.ShouldBind(&reqBody); err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := cc.categoryService.SaveCategory(reqBody)

	if err != nil {

		fail := utils.BuildResponseFailed(err.Error())
		ctx.JSON(http.StatusBadRequest, fail)
		return
	}

	result := utils.BuildResponseSucces("SUCCES", res)
	ctx.JSON(http.StatusOK, result)
}
