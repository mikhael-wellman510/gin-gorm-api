package service

import (
	"GinGonicGorm/dto"
	"GinGonicGorm/entity"
	"GinGonicGorm/repository"
	"context"

	"math"
)

type (
	ProductService interface {
		CreateProductService(ctx context.Context, req dto.ProductRequest) (dto.ProductResponse, error)
		FindProductById(ctx context.Context, productId string) (dto.ProductResponse, error)
		FindAllProduct(ctx context.Context) ([]dto.ProductResponse, error)
		UpdateProduct(ctx context.Context, req dto.ProductRequest) (dto.ProductResponse, error)
		PagginationProductWithFilterService(ctx context.Context, req dto.PagginationRequest) (dto.ProductResponsePaggingAndFilter, error)
	}

	productService struct {
		productRepository repository.ProductRepository
	}
)

func NewProductService(productRepository repository.ProductRepository) ProductService {

	return &productService{
		productRepository: productRepository,
	}
}

func (ps *productService) CreateProductService(ctx context.Context, req dto.ProductRequest) (dto.ProductResponse, error) {

	prod := entity.Product{
		Name:               req.Name,
		Brand:              req.Brand,
		Price:              req.Price,
		WeightProduct:      req.WeightProduct,
		StockProduct:       req.StockProduct,
		DescriptionProduct: req.DescriptionProduct,
	}

	// panggil lewat struct
	product, err := ps.productRepository.SaveProduct(ctx, nil, prod)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		Id:                 product.ID,
		Name:               product.Name,
		Price:              product.Price,
		Brand:              product.Brand,
		WeightProduct:      product.WeightProduct,
		StockProduct:       product.StockProduct,
		DescriptionProduct: product.DescriptionProduct,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
	}, nil
}

func (ps *productService) FindProductById(ctx context.Context, productId string) (dto.ProductResponse, error) {

	product, err := ps.productRepository.FindProductById(ctx, productId)

	if err != nil {

		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		Id:                 productId,
		Name:               product.Name,
		Price:              product.Price,
		Brand:              product.Brand,
		WeightProduct:      product.WeightProduct,
		StockProduct:       product.StockProduct,
		DescriptionProduct: product.DescriptionProduct,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
	}, nil
}

func (ps *productService) FindAllProduct(ctx context.Context) ([]dto.ProductResponse, error) {
	product, err := ps.productRepository.FindAllProduct(ctx)

	if err != nil {
		return []dto.ProductResponse{}, err
	}

	var arrProductBuilder []dto.ProductResponse

	for _, val := range product {

		data := dto.ProductResponse{
			Id:                 val.ID,
			Name:               val.Name,
			Price:              val.Price,
			Brand:              val.Brand,
			WeightProduct:      val.WeightProduct,
			StockProduct:       val.StockProduct,
			DescriptionProduct: val.DescriptionProduct,
			CreatedAt:          val.CreatedAt,
			UpdatedAt:          val.UpdatedAt,
		}

		arrProductBuilder = append(arrProductBuilder, data)
	}
	return arrProductBuilder, nil
}

func (ps *productService) UpdateProduct(ctx context.Context, req dto.ProductRequest) (dto.ProductResponse, error) {

	// Cek ada atau tidak dulu
	pr, err := ps.FindProductById(ctx, req.Id)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	product := entity.Product{
		Base:               entity.Base{ID: pr.Id},
		Name:               req.Name,
		Brand:              req.Brand,
		Price:              req.Price,
		WeightProduct:      req.WeightProduct,
		StockProduct:       req.StockProduct,
		DescriptionProduct: req.DescriptionProduct,
	}

	editProducts, err := ps.productRepository.UpdatedProduct(ctx, nil, product)

	if err != nil {
		return dto.ProductResponse{}, err
	}
	return dto.ProductResponse{
		Id:                 editProducts.ID,
		Name:               editProducts.Name,
		Price:              editProducts.Price,
		Brand:              editProducts.Brand,
		WeightProduct:      editProducts.WeightProduct,
		StockProduct:       editProducts.StockProduct,
		DescriptionProduct: editProducts.DescriptionProduct,
		CreatedAt:          editProducts.CreatedAt,
		UpdatedAt:          editProducts.UpdatedAt,
	}, nil
}

func (ps *productService) PagginationProductWithFilterService(ctx context.Context, req dto.PagginationRequest) (dto.ProductResponsePaggingAndFilter, error) {

	res, totalItems, err := ps.productRepository.PagginationAndSearchProduct(ctx, req.Filter, req.Limit, req.Offset)
	totalPage := int(math.Ceil(float64(totalItems) / float64(req.Limit)))

	if err != nil {
		return dto.ProductResponsePaggingAndFilter{}, err
	}

	var productRespons []dto.ProductResponse
	for _, val := range res {
		res := dto.ProductResponse{
			Id:                 val.ID,
			Name:               val.Name,
			Price:              val.Price,
			Brand:              val.Brand,
			WeightProduct:      val.WeightProduct,
			StockProduct:       val.StockProduct,
			DescriptionProduct: val.DescriptionProduct,
			CreatedAt:          val.CreatedAt,
			UpdatedAt:          val.UpdatedAt,
		}
		productRespons = append(productRespons, res)
	}
	pagging := dto.PagginationResponse{
		Page:             req.Offset,
		Size:             req.Limit,
		TotalDataCurrent: len(productRespons),
		TotalPage:        totalPage,
		TotalData:        totalItems,
	}

	result := dto.ProductResponsePaggingAndFilter{
		ProductResponse: productRespons,
		PaggingResponse: pagging,
	}
	return result, nil
}
