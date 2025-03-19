package main

import (
	"GinGonicGorm/config"
	"GinGonicGorm/controller"
	"GinGonicGorm/migrations"
	"GinGonicGorm/repository"
	"GinGonicGorm/routes"
	"GinGonicGorm/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	// Register product Module
	var productRepository repository.ProductRepository = repository.NewProductRepository(db)
	var productService service.ProductService = service.NewProductService(productRepository)
	var productController controller.ProductController = controller.NewProductController(productService)
	// var bookController controller.BookController = controller.NewBookController()

	// routes.Book(app, bookController)
	routes.Product(app, productController)

	// Port and Running
	port := os.Getenv("SERVER_PORT")
	app.Run(":" + port)

}
