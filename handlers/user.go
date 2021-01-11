package handlers

import (
	"fmt"
	"test/server/database"
	"test/server/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers for get all users
func GetAllUsers(ctx *fiber.Ctx) error {
	users := []model.User{}
	result := database.DB.Limit(10).Find(&users)
	if result.Error != nil {
		return ctx.SendStatus(fiber.StatusBadGateway)
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "Success Get All Users", "data": users})
}

// GetUserByID for get user by id
func GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := &model.User{}
	result := database.DB.Find(&user, id)
	if result.Error != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	if user.ID == 0 {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Success Get User", "data": user})
}

// UpdateUserByID for update user by id
func UpdateUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := &model.User{}
	result := database.DB.Find(&user, id)
	if result.Error != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	type registerInput struct {
		Username    string `form:"username" `
		Password    string `form:"password" `
		NamaLengkap string `form:"namaLengkap" `
	}
	var input registerInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	file, err := ctx.FormFile("foto")
	if err != nil {
		return err
	}
	pathFile := fmt.Sprintf("./uploads/%s", file.Filename)
	ctx.SaveFile(file, pathFile)

	userNew := model.User{Username: input.Username, Password: input.Password, NamaLengkap: input.NamaLengkap, Foto: pathFile}
	database.DB.Model(&user).Updates(userNew)

	return ctx.JSON(fiber.Map{"status": "success", "message": "Success Update User", "data": user})
}

// DeleteUserByID for delete user by id
func DeleteUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := &model.User{}
	result := database.DB.Find(&user, id)
	if result.Error != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	database.DB.Delete(&model.User{}, id)

	return ctx.JSON(fiber.Map{"status": "success", "message": "Success Delete User", "data": nil})
}
