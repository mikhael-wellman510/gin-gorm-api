package controller

import (
	"GinGonicGorm/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	TestingController interface {
		Testing(ctx *gin.Context)
	}

	testingController struct {
	}
)

func NewTestingController() TestingController {

	return &testingController{}
}

func (tc *testingController) Testing(ctx *gin.Context) {
	reqBody := dto.TestingDto{}

	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusOK, reqBody.Scope)

}
