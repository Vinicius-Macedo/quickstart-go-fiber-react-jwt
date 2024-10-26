package handlers

import (
	"app/cmd/internal/helpers"
	"app/cmd/internal/postgresRepo"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

func RequestUserHandler(c *fiber.Ctx) error {
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

	queries := postgresRepo.New(dbConnection)
	userData, err := queries.GetUserByEmail(ctx, email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"name": userData.Name, "username": userData.Username, "email": userData.Email, "image": userData.ImageUrl.String})
}

func RequestLogoutHandler(c *fiber.Ctx) error {
	c.ClearCookie("jwt")
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.SendStatus(fiber.StatusNoContent)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RequestLoginHandler(c *fiber.Ctx) error {
	var loginReq LoginRequest

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel, dbConnection, err := postgresRepo.InitContextAndDB()
	defer cancel()
	defer dbConnection.Close()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	queries := postgresRepo.New(dbConnection)

	user, err := queries.GetUserByEmail(ctx, loginReq.Email)
	if err != nil {
		// Dummy hash to prevent timing attacks
		helpers.VerifyPassword(loginReq.Password, "$2a$14$dummyHashdummyHashdummyHashdummyHashdummyHashdummyHashdummyHashdummyHash")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error a": err.Error()})
	}

	if !helpers.VerifyPassword(loginReq.Password, user.Password) {
		fmt.Println("Password does not match", loginReq.Password)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(helpers.GetEnv("JWT_SECRET", "secret")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.SendStatus(fiber.StatusOK)

}

func User(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

func CreateUser(c *fiber.Ctx) error {
	var createUserReq postgresRepo.CreateUserParams

	fmt.Println("CreateUser")

	if err := c.BodyParser(&createUserReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel, dbConnection, err := postgresRepo.InitContextAndDB()
	defer cancel()
	defer dbConnection.Close()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	createUserReq.Password, err = helpers.HashPassword(createUserReq.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	queries := postgresRepo.New(dbConnection)

	// check if user already exists
	_, err = queries.CheckUserExistsByEmail(ctx, createUserReq.Email)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email já existe"})
	}

	_, err = queries.CheckUserExistsByUsername(ctx, createUserReq.Username)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Usuário já existe"})
	}

	newUser, err := queries.CreateUser(ctx, createUserReq)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": newUser})

}

type PasswordResetRequest struct {
	Email string `json:"email"`
}

func RequestPasswordResetHandler(c *fiber.Ctx) error {
	var passwordResetReq PasswordResetRequest

	if err := c.BodyParser(&passwordResetReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	ctx, cancel, dbConnection, err := postgresRepo.InitContextAndDB()
	defer cancel()
	defer dbConnection.Close()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	queries := postgresRepo.New(dbConnection)
	user, err := queries.GetUserByEmail(ctx, passwordResetReq.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	})

	// Sign the token
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// Create the password reset link
	baseURL := os.Getenv("FRONTEND_URL")
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", baseURL, tokenString)

	// Email content
	subject := "Password Reset Request"
	body := fmt.Sprintf("<p>To reset your password, click the link below:</p><p><a href=\"%s\">Reset Password</a></p>", resetLink)

	// Send the email
	err = helpers.SendEmail(passwordResetReq.Email, subject, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email"})
	}

	return c.JSON(fiber.Map{"message": "Password reset email sent successfully"})
}

type TokenAndNewPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func RequestHandleTokenAndNewPasswordHandler(c *fiber.Ctx) error {
	var tokenAndNewPasswordReq TokenAndNewPasswordRequest

	if err := c.BodyParser(&tokenAndNewPasswordReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Parse the token
	token, err := jwt.Parse(tokenAndNewPasswordReq.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims := token.Claims.(jwt.MapClaims)

	// Get the email from the token
	email := claims["email"].(string)

	// Hash the new password
	hashedPassword, err := helpers.HashPassword(tokenAndNewPasswordReq.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash the password"})
	}

	ctx, cancel, dbConnection, err := postgresRepo.InitContextAndDB()
	defer cancel()
	defer dbConnection.Close()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	queries := postgresRepo.New(dbConnection)

	newData := postgresRepo.UpdateUserPasswordByEmailParams{
		Email:    email,
		Password: hashedPassword,
	}

	err = queries.UpdateUserPasswordByEmail(ctx, newData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Password updated successfully"})

}
