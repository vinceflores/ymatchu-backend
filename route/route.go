package route

import (
	// go fiber
	"ymatchu-backend/student"

	"github.com/gofiber/fiber/v2"
)

func RouteInit() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("YMatchu API")
	})

	// landlord
	app.Post("/createListing", func(c *fiber.Ctx) error {
		return c.SendString("Create Listing")
	})

	// student
	app.Post("/student/filterRequest", student.Filter)
	// post listingDetails
	app.Post("/student/listingDetails/:id", student.ListingDetails)

	return app
}
