package controller

import (
	"GinGonicGorm/dto"
	"GinGonicGorm/service"
	"GinGonicGorm/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	ProductController interface {
		CreateProduct(ctx *gin.Context)
		FindAllProduct(ctx *gin.Context)
		FindProductById(ctx *gin.Context)
		UpdatedProduct(ctx *gin.Context)
		PagginationProductWithFilter(ctx *gin.Context)
	}

	productController struct {
		productService service.ProductService
	}
)

func NewProductController(productService service.ProductService) ProductController {

	return &productController{
		productService: productService,
	}
}

// @Summary Create Product
// @Description Menambahkan produk baru ke database
// @Tags Products
// @Accept json
// @Produce json
// @Param request body dto.ProductRequest true "Product Data"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/product/createProduct [post]
func (pc *productController) CreateProduct(ctx *gin.Context) {

	reqBody := dto.ProductRequest{}

	/*gunakan return untuk menyelesaikan program*/
	if err := ctx.ShouldBind(&reqBody); err != nil {
		/*
			ini akan masuk error [tipe data beda , ada field yg tidak di isi (required)]
		*/

		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := pc.productService.CreateProductService(ctx, reqBody)

	if err != nil {

		fail := utils.BuildResponseFailed(err.Error())

		ctx.JSON(http.StatusBadRequest, fail)
		return
	}

	res := utils.BuildResponseSucces("Succes", result)
	ctx.JSON(http.StatusOK, res)
}

// @Summary Get All Products
// @Description Fetch all available products
// @Tags Products
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/product/findAll [get]
func (pc *productController) FindAllProduct(ctx *gin.Context) {

	res, err := pc.productService.FindAllProduct(ctx)

	if err != nil {
		fail := utils.BuildResponseFailed(err.Error())
		ctx.JSON(http.StatusBadRequest, fail)
		return
	}

	result := utils.BuildResponseSucces("Succes Find All", res)
	ctx.JSON(http.StatusOK, result)
}

func (pc *productController) FindProductById(ctx *gin.Context) {
	userId := ctx.Param("id")

	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Id Tidak boleh kosong")
		return
	}

	res, err := pc.productService.FindProductById(ctx, userId)

	if err != nil {
		fail := utils.BuildResponseFailed(err.Error())
		ctx.JSON(http.StatusNotFound, fail)
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSucces("Found", res))
}

func (pc *productController) UpdatedProduct(ctx *gin.Context) {
	var req dto.ProductRequest

	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return

	}

	res, err := pc.productService.UpdateProduct(ctx, req)

	if err != nil {
		res := utils.BuildResponseFailed(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := utils.BuildResponseSucces("Success Updated Data", res)

	ctx.JSON(http.StatusOK, result)
}

func (pc *productController) PagginationProductWithFilter(ctx *gin.Context) {
	namePrefix := ctx.Query("productName")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	req := dto.PagginationRequest{
		Filter: namePrefix,
		Limit:  limit,
		Offset: offset,
	}

	res, err := pc.productService.PagginationProductWithFilterService(ctx, req)

	if err != nil {
		res := utils.BuildResponseFailed(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result := utils.BuildResponseSucces("Success Get Data", res)

	ctx.JSON(http.StatusOK, result)
}
