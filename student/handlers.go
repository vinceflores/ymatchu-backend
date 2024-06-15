package student

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Student struct {
	DB *gorm.DB
}

func (student *Student) Filter(c *fiber.Ctx) error {
	// get the student id from the request
	// get the preferences from the request
	// save the preferences
	// send array of matching listings
	// student.DB.Preload()	
	return 	c.Status(200).JSON(fiber.Map{
		"message": "Filtering Request",
	})
}

func (student *Student) ListingDetails(c *fiber.Ctx) error {
	// get the student id from the request
	// get the listing id from the request
	// get the listing details
	// return the listing details

	return c.Status(200).JSON(fiber.Map{
		"message": "Listing Details",
	})
}
