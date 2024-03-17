package router

import (
	"golang-intern/assignment-2/middleware"
	"golang-intern/assignment-2/service"
	"golang-intern/assignment-2/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func SetRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {

		//middleware
		reqTime := time.Now()
		requestParams := ""
		requestHeaders := util.ConvertMapToString(c.GetReqHeaders())
		requestBody := string(c.Body())

		responseTime := time.Now()
		responseStatus := c.Response().StatusCode()
		responseHeaders := util.ConvertMapToString(c.GetRespHeaders())
		responseBody := string(c.Body())

		_, err := middleware.StoreAPIRequestRespond(util.Ctx,
			reqTime, requestParams, requestHeaders, requestBody,
			responseTime, responseStatus, responseHeaders, responseBody)

		if err != nil {
			return err
		}

		return c.SendString("Hello, World!")
	})

	app.Post("/fetchgeo", func(c *fiber.Ctx) error {

		reqTime := time.Now()
		requestParams := ""
		requestHeaders := util.ConvertMapToString(c.GetReqHeaders())
		requestBody := string(c.Body())

		//
		err := service.HandlerData(0)
		if err != nil {
			return err
		}

		//
		responseTime := time.Now()
		responseStatus := c.Response().StatusCode()
		responseHeaders := util.ConvertMapToString(c.GetRespHeaders())
		responseBody := string(c.Body())

		_, err = middleware.StoreAPIRequestRespond(util.Ctx,
			reqTime, requestParams, requestHeaders, requestBody,
			responseTime, responseStatus, responseHeaders, responseBody)

		if err != nil {
			return err
		}

		//
		return c.SendString("fetched all data!")
	})

	app.Post("/fetchgeo/:number", func(c *fiber.Ctx) error {
		numberValue := c.Params("number")
		number, _ := strconv.Atoi(numberValue)

		reqTime := time.Now()
		requestParams := ""
		requestHeaders := util.ConvertMapToString(c.GetReqHeaders())
		requestBody := string(c.Body())

		//
		err := service.HandlerData(number)
		if err != nil {
			return err
		}

		//
		responseTime := time.Now()
		responseStatus := c.Response().StatusCode()
		responseHeaders := util.ConvertMapToString(c.GetRespHeaders())
		responseBody := string(c.Body())

		_, err = middleware.StoreAPIRequestRespond(util.Ctx,
			reqTime, requestParams, requestHeaders, requestBody,
			responseTime, responseStatus, responseHeaders, responseBody)

		if err != nil {
			return err
		}

		return c.SendString("fetched " + numberValue + "first data!")
	})

	//multiple filter
	app.Get("/earthquake", func(c *fiber.Ctx) error {
		limitValue := c.Query("limit")
		offsetValue := c.Query("offset")
		placeValue := c.Query("place")

		if limitValue == "" {
			limitValue = "10"
		}

		if offsetValue == "" {
			offsetValue = "1"
		}

		//Middleware
		reqTime := time.Now()
		requestParams := limitValue + ", " + offsetValue
		requestHeaders := util.ConvertMapToString(c.GetReqHeaders())
		requestBody := string(c.Body())

		//handler
		var result error
		limit, err := strconv.ParseUint(limitValue, 10, 32)

		if err != nil {
			return err
		}

		offset, err := strconv.ParseUint(offsetValue, 10, 32)

		if err != nil {
			return err
		}

		if placeValue != "" {
			requestParams += "," + placeValue
			result = service.GetEarthquakeFilteredByLocation(c, int(limit), int(offset), placeValue)
		} else {
			result = service.GetEarthquakeByLimitOffset(c, int(limit), int(offset))
		}

		responseTime := time.Now()
		responseStatus := c.Response().StatusCode()
		responseHeaders := util.ConvertMapToString(c.GetRespHeaders())
		responseBody := string(c.Body())

		_, err = middleware.StoreAPIRequestRespond(util.Ctx,
			reqTime, requestParams, requestHeaders, requestBody,
			responseTime, responseStatus, responseHeaders, responseBody)

		if err != nil {
			return err
		}

		return result
	})

	//service.GetAllEarthquake
	app.Get("/earthquake/all", func(c *fiber.Ctx) error {

		//middleware
		reqTime := time.Now()
		requestParams := ""
		requestHeaders := util.ConvertMapToString(c.GetReqHeaders())
		requestBody := string(c.Body())

		//handle
		result := service.GetAllEarthquake(c)

		//
		responseTime := time.Now()
		responseStatus := c.Response().StatusCode()
		responseHeaders := util.ConvertMapToString(c.GetRespHeaders())
		responseBody := string(c.Body())

		_, err := middleware.StoreAPIRequestRespond(util.Ctx,
			reqTime, requestParams, requestHeaders, requestBody,
			responseTime, responseStatus, responseHeaders, responseBody)

		if err != nil {
			return err
		}

		return result
	})

	//filter eathquake by ID
	app.Get("/earthquake/:id", func(c *fiber.Ctx) error {

		//middleware
		reqTime := time.Now()
		requestParams := ""
		requestHeaders := util.ConvertMapToString(c.GetReqHeaders())
		requestBody := string(c.Body())

		//
		idValue := utils.CopyString(c.Params("id"))

		id, err := strconv.ParseUint(idValue, 10, 32)

		if err != nil {
			return err
		}

		result := service.GetEarthquakeFilteredByID(c, int(id))

		//
		responseTime := time.Now()
		responseStatus := c.Response().StatusCode()
		responseHeaders := util.ConvertMapToString(c.GetRespHeaders())
		responseBody := string(c.Body())

		_, err = middleware.StoreAPIRequestRespond(util.Ctx,
			reqTime, requestParams, requestHeaders, requestBody,
			responseTime, responseStatus, responseHeaders, responseBody)

		if err != nil {
			return err
		}
		return result
	})

	app.Get("/getrequestbyweek/:number", func(c *fiber.Ctx) error {
		numberValue := c.Params("number")
		number, err := strconv.ParseUint(numberValue, 10, 32)

		if err != nil {
			return err
		}

		//middleware
		reqTime := time.Now()
		requestParams := numberValue
		requestHeaders := util.ConvertMapToString(c.GetReqHeaders())
		requestBody := string(c.Body())

		//
		result := service.GetRequestByWeek(c, int(number))
		//
		responseTime := time.Now()
		responseStatus := c.Response().StatusCode()
		responseHeaders := util.ConvertMapToString(c.GetRespHeaders())
		responseBody := string(c.Body())

		_, err = middleware.StoreAPIRequestRespond(util.Ctx,
			reqTime, requestParams, requestHeaders, requestBody,
			responseTime, responseStatus, responseHeaders, responseBody)

		if err != nil {
			return err
		}

		return result
	})

	app.Get("/getrequestbyhour/:number", func(c *fiber.Ctx) error {
		numberValue := c.Params("number")
		number, err := strconv.ParseUint(numberValue, 10, 32)

		if err != nil {
			return err
		}

		//middleware
		reqTime := time.Now()
		requestParams := numberValue
		requestHeaders := util.ConvertMapToString(c.GetReqHeaders())
		requestBody := string(c.Body())

		//
		result := service.GetRequestByHour(c, int(number))
		//
		responseTime := time.Now()
		responseStatus := c.Response().StatusCode()
		responseHeaders := util.ConvertMapToString(c.GetRespHeaders())
		responseBody := string(c.Body())

		_, err = middleware.StoreAPIRequestRespond(util.Ctx,
			reqTime, requestParams, requestHeaders, requestBody,
			responseTime, responseStatus, responseHeaders, responseBody)

		if err != nil {
			return err
		}

		return result
	})

}
