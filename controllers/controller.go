package controllers

import (
	"antrian/models"

	"github.com/gofiber/fiber/v2"
)

/*
func Index(c *fiber.Ctx) error {

	var ant []models.Ant

	models.DB.Db.Find(&ant)

	return c.Status(fiber.StatusOK).JSON(ant)

}*/

func Create(c *fiber.Ctx) error {

	antri := new(models.Antrian)

	if err := c.BodyParser(antri); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Create(&antri)

	return c.Status(fiber.StatusCreated).JSON(antri)
}

func Show(c *fiber.Ctx) error {

	ant := &models.Antrian{}
	id := c.Params("id")

	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ant)
}

func Update(c *fiber.Ctx) error {

	ant := &models.Antrian{}
	id := c.Params("id")

	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	if err := c.BodyParser(ant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	ant.Seq = ant.Seq + 1
	models.DB.Db.Save(ant)

	return c.Status(fiber.StatusOK).JSON(ant)

}

func Reset(c *fiber.Ctx) error {

	ant := &models.Antrian{}
	id := c.Params("id")

	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	if err := c.BodyParser(ant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	ant.Seq = 0
	models.DB.Db.Save(ant)

	return c.Status(fiber.StatusOK).JSON(ant)
}

/*
func Delete(c *fiber.Ctx) error {

	antri := &models.Ant{}
	id := c.Params("id")

	if err := models.DB.Db.First(antri, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	models.DB.Db.Delete(antri, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Delete D"})
}*/
