package handlers

import (
	"app/cmd/internal/postgresRepo"
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequestUploadUserImageHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println("erro aqui")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse image"})
	}

	imageUrl, err := saveImage(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = saveImageUrlOnUser(c, imageUrl)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully"})
}

func saveImage(c *fiber.Ctx, imageFile *multipart.FileHeader) (string, error) {
	file := imageFile

	// Ensure the ../public/img directory exists
	if _, err := os.Stat("../public/img"); os.IsNotExist(err) {
		err = os.Mkdir("../public/img", os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// Save the file
	savePath := fmt.Sprintf("../public/img/%s", image)
	err := c.SaveFile(file, savePath)
	if err != nil {
		fmt.Println("erro aqui 2")
		return "", err
	}

	baseUrl := os.Getenv("BASE_URL")
	imageUrl := fmt.Sprintf("%s/images/%s", baseUrl, image)

	return imageUrl, nil

}

func saveImageUrlOnUser(c *fiber.Ctx, imageUrl string) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["email"] == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	email := claims["email"].(string)

	ctx, cancel, dbConnection, err := postgresRepo.InitContextAndDB()
	defer cancel()
	defer dbConnection.Close()

	if err != nil {
		return err
	}

	var createImageRequest postgresRepo.InsertImageOnUserByEmailParams
	createImageRequest.UserEmail = email
	createImageRequest.ImageUrl = imageUrl

	queries := postgresRepo.New(dbConnection)
	_, err = queries.InsertImageOnUserByEmail(ctx, createImageRequest)

	if err != nil {
		return err
	}

	return nil
}
