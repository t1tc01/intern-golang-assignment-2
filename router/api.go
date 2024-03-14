package router

import (
	"golang-intern/assignment-2/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/fetchgeo", func(c *fiber.Ctx) error {
		err := service.HandlerData()
		if err != nil {
			return err
		}
		return c.SendString("fetched data!")
	})

	app.Get("/earthquake", func(c *fiber.Ctx) error {
		limitValue := c.Query("limit")
		offsetValue := c.Query("offset")
		filterValue := c.Query("filteredBy")

		if limitValue == "" && offsetValue == "" {
			limitValue = "10"
			offsetValue = "1"
		}

		limit, err := strconv.ParseUint(limitValue, 10, 32)

		if err != nil {
			return err
		}

		offset, err := strconv.ParseUint(offsetValue, 10, 32)

		if err != nil {
			return err
		}

		//filtered
		if filterValue != "" {
			return service.GetEarthquakeByLimitOffset(c, int(limit), int(offset))
		}

		//non-filtered
		return service.GetEarthquakeByLimitOffset(c, int(limit), int(offset))
	})

	app.Get("/earthquake/all", service.GetAllEarthquake)

	app.Get("/earthquake/:id", func(c *fiber.Ctx) error {
		//
		idValue := utils.CopyString(c.Params("id"))

		id, err := strconv.ParseUint(idValue, 10, 32)

		if err != nil {
			return err
		}

		return service.GetEarthquakeFilteredByID(c, int(id))
	})

}
