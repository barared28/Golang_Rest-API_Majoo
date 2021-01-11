package handlers

import (
	"fmt"
	"test/server/database"
	"test/server/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Login for auth login user
func Login(ctx *fiber.Ctx) error {
	type loginInput struct {
		Username string `json:"username" `
		Password string `json:"password" `
	}
	var input loginInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	user := &model.User{}
	result := database.DB.Where(&model.User{Username: input.Username, Password: input.Password}).First(&user)
	if result.Error != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)

	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Success Login", "data": t})
}

// Register for auth login user
func Register(ctx *fiber.Ctx) error {
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
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	pathFile := fmt.Sprintf("./uploads/%s", file.Filename)
	ctx.SaveFile(file, pathFile)
	user := model.User{Username: input.Username, Password: input.Password, NamaLengkap: input.NamaLengkap, Foto: pathFile}
	database.DB.Create(&user)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)

	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Success Register", "data": t})
}
