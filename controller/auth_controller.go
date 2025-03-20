package controller

import (
	"GinGonicGorm/dto"
	"GinGonicGorm/service"
	"GinGonicGorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AuthController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
	}

	authController struct {
		authService service.AuthService
	}
)

func NewAuthController(authService service.AuthService) AuthController {

	return &authController{
		authService: authService,
	}
}

func (uc *authController) Register(ctx *gin.Context) {

	var reqBody = dto.UserRequest{}

	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := uc.authService.RegisterAccount(ctx, reqBody)

	if err != nil {
		fail := utils.BuildResponseFailed(err.Error())
		ctx.JSON(http.StatusBadRequest, fail)
		return
	}

	result := utils.BuildResponseSucces("Succes Register", res)
	ctx.JSON(http.StatusOK, result)
}

func (uc *authController) Login(ctx *gin.Context) {

	reqBody := dto.LoginRequest{}

	if err := ctx.ShouldBind(&reqBody); err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return

	}

	res, err := uc.authService.LoginAccount(ctx, reqBody)

	if err != nil {
		fail := utils.BuildResponseFailed(err.Error())
		ctx.JSON(http.StatusBadRequest, fail)
		return
	}

	result := utils.BuildResponseSucces("Succes Login", res)
	ctx.JSON(http.StatusOK, result)

}
