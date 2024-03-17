package main

import (
	"golang-intern/assignment-2/db"
	"golang-intern/assignment-2/router"
	"log"

	_ "golang-intern/assignment-2/docs"

	_ "entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

//	@title			Golang Intern Assignment 2
//	@version		1.0
//	@description	This is a Assignment 2.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		petstore.swagger.io
//	@BasePath	/v2

func main() {
	db.ConnectDb()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	router.SetRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
