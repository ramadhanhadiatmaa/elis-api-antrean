package controllers

import (
	"antrian/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error{
	
	var antrian []models.Antrian

	models.DB.Find(&antrian)

	return c.Status(fiber.StatusOK).JSON(antrian)

}

func Show(c *fiber.Ctx) error{

	id := c.Params("id")

	var antrian models.Antrian

	if err := models.DB.Model(&antrian).Where("id = ? ", id).First(&antrian).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data not found",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})

	}

	return c.JSON(antrian)
}

func Create(c *fiber.Ctx) error {

	var antrian models.Antrian

	if err := c.BodyParser(&antrian); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	antrian.Updated = time.Now()
	antrian.Created = time.Now()
	antrian.Num = 0

	if err := models.DB.Create(&antrian).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data Created",
	})
}

func Update(c *fiber.Ctx) error{

	id := c.Params("id")

	var antrian models.Antrian

	if err := c.BodyParser(&antrian); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	antrian.Updated = time.Now()
	antrian.Num = antrian.Num + 1

	if models.DB.Where("id = ?", id).Updates(&antrian).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data Updated",
	})
}

func Reset(c *fiber.Ctx) error{

	id := c.Params("id")

	var antrian models.Antrian

	if err := c.BodyParser(&antrian); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	antrian.Updated = time.Now()
	antrian.Num = 0

	if models.DB.Where("id = ?", id).Updates(&antrian).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data Updated",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	var antrian models.Antrian

	if models.DB.Where("id = ?", id).Delete(&antrian).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data Deleted",
	})
}