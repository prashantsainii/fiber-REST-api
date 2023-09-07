package routes

import (
	// "errors"

	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/prashantsainii/fiber-REST-api/database"
	"github.com/prashantsainii/fiber-REST-api/models"
)

type User struct {
	// this is not the model User, see this as serializer

	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User does not exist")
	}
	return nil
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}				// empty slice which will contain user of type models.User

	database.Database.Db.Find(&users)		// will put all user in our slice
	responseUsers := []User{}				// empty slice which will contain user of type User

	for _, user := range users {			// filling up the slice with all users
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser) 
	}

	return c.Status(200).JSON(responseUsers)	// returning all users
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User		

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return c.Status(400).JSON("User does not exist")
	}

	responseUser := CreateResponseUser(user)
	
	return c.Status(200).JSON(responseUser)

}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return c.Status(400).JSON("User does not exist")
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return c.Status(400).JSON("User does not exist")
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted User")
}