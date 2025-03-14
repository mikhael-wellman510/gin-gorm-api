package main

import (
	"GinGonicGorm/config"
	"GinGonicGorm/controller"
	"GinGonicGorm/migrations"
	"GinGonicGorm/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	/*Config DB Mysql*/
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)
	/* ini untuk mengisi migrations */
	migrations.Migrate(db)

	// ==== Routes ====
	app := gin.Default()

	var bookController controller.BookController = controller.NewBookController()
	var productController controller.ProductController = controller.NewProductController()
	routes.Book(app, bookController)
	routes.Product(app, productController)

	// Port and Running
	port := os.Getenv("SERVER_PORT")
	app.Run(":" + port)

}
