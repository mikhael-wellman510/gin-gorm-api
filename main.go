package main

import (
	"GinGonicGorm/config"
	"GinGonicGorm/controller"
	_ "GinGonicGorm/docs"

	// Import Swagger docs
	"GinGonicGorm/migrations"
	"GinGonicGorm/repository"
	"GinGonicGorm/routes"
	"GinGonicGorm/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	/*Config DB Mysql*/
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	/* ini untuk mengisi migrations */
	err := migrations.Migrate(db)
	if err != nil {
		log.Fatal("Migration Failed !!!")
	}
	// ==== Routes ====
	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Swagger setup

	// Tipe Data interface semua
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository(db)
	var categoryService service.CategoryService = service.NewCategoryService(categoryRepository)
	var categoryController controller.CategoryController = controller.NewCategoryService(categoryService)
	// Register product Module
	var productRepository repository.ProductRepository = repository.NewProductRepository(db)
	var productService service.ProductService = service.NewProductService(productRepository, categoryService)
	var productController controller.ProductController = controller.NewProductController(productService)
	// var bookController controller.BookController = controller.NewBookController()

	// routes.Book(app, bookController)
	var userRepository repository.UserRepository = repository.NewUserRepository(db)
	var authService service.AuthService = service.NewAuthService(userRepository)
	var authController controller.AuthController = controller.NewAuthController(authService)

	var photoController controller.PhotoController = controller.NewPhotoController()
	routes.Router(app, productController, authController, categoryController, photoController)

	// Port and Running
	port := os.Getenv("SERVER_PORT")
	app.Run(":" + port)

}
