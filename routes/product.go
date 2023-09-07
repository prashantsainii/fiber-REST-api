package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/prashantsainii/fiber-REST-api/database"
	"github.com/prashantsainii/fiber-REST-api/models"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}				// empty slice which will contain user of type models.User

	database.Database.Db.Find(&products)		// will put all user in our slice
	responseProducts := []Product{}				// empty slice which will contain user of type User

	for _, product := range products {			// filling up the slice with all users
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct) 
	}

	return c.Status(200).JSON(responseProducts)	// returning all users
}

// using the helper function this time

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("Product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product		

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// database.Database.Db.Find(&product, "id = ?", id)
	// if product.ID == 0 {
	// 	return c.Status(400).JSON("Product does not exist")
	// }

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)

}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// database.Database.Db.Find(&product, "id = ?", id)
	// if product.ID == 0 {
	// 	return c.Status(400).JSON("Product does not exist")
	// }

	type UpdateProduct struct {
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// database.Database.Db.Find(&product, "id = ?", id)
	// if product.ID == 0 {
	// 	return c.Status(400).JSON("Product does not exist")
	// }

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted Product")
}