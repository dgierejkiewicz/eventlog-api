package routes

import (
		__query "eventlog/cmd/query"
		"eventlog/entity"
		"eventlog/service/event"
		"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, service event.Service) {
		app.Get("/v1/events/type::type", getEventsByTypeHandler(service))
		app.Get("/v1/events/from::from/to::to", getEventsByTimeRangeHandler(service))
		app.Get("/v1/events/type::type/from::from/to::to", getEventsByTypeAndTimeRangeHandler(service))
		app.Post("/v1/event", createEventHandler(service))
}

func createEventHandler(service event.Service) fiber.Handler {
		return func(c *fiber.Ctx) error {
				var requestBody entity.Event
				err := c.BodyParser(&requestBody)
				if err != nil {
						return c.JSON(&fiber.Map{
								"success": false,
								"error":   err,
						})
				}
				result, dberr := service.CreateEvent(&requestBody)
				return c.JSON(&fiber.Map{
						"data":  result,
						"error": dberr,
				})
		}
}

func getEventsByTypeHandler(service event.Service) fiber.Handler {
		return func(c *fiber.Ctx) error {
				var requestParam string
				requestParam = c.Params("type")
				query, err := __query.NewTypeQuery(requestParam)
				if err != nil {
						return c.JSON(&fiber.Map{
								"success": false,
								"error":   err,
						})
				}
				result, dberr := service.FetchEventsByType(query)
				return c.JSON(&fiber.Map{
						"data":  result,
						"error": dberr,
				})
		}
}

func getEventsByTimeRangeHandler(service event.Service) fiber.Handler {
		return func(c *fiber.Ctx) error {
				requestParamFrom := c.Params("from")
				requestParamTo := c.Params("to")
				query, err := __query.NewTimeRangeQuery(
						requestParamFrom,
						requestParamTo)
				if err != nil {
						return c.JSON(&fiber.Map{
								"success": false,
								"error":   err,
						})
				}
				result, dberr := service.FetchEventsByTimeRange(query)
				return c.JSON(&fiber.Map{
						"data":  result,
						"error": dberr,
				})
		}
}

func getEventsByTypeAndTimeRangeHandler(service event.Service) fiber.Handler {
		return func(c *fiber.Ctx) error {
				requestParamType := c.Params("type")
				requestParamFrom := c.Params("from")
				requestParamTo := c.Params("to")
				query, err := __query.NewTypeTimeRangeQuery(
						requestParamType,
						requestParamFrom,
						requestParamTo)
				if err != nil {
						return c.JSON(&fiber.Map{
								"success": false,
								"error":   err,
						})
				}
				result, dberr := service.FetchEventsByTypeAndTimeRange(query)
				return c.JSON(&fiber.Map{
						"data":  result,
						"error": dberr,
				})
		}
}
