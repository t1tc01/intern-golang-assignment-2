package main

import (
	"golang-intern/assignment-2/db"
	"golang-intern/assignment-2/router"
	"log"

	_ "entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	db.ConnectDb()

	app := fiber.New()
	app.Use(cors.New())

	router.SetRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
