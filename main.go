package main

import (
	"fmt"

	_ "github.com/mehranhn/go_example_1/docs"
	"github.com/mehranhn/go_example_1/middlewares"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/mehranhn/go_example_1/config"
	"github.com/mehranhn/go_example_1/controllers"
	memoryimpinmemory "github.com/mehranhn/go_example_1/external/memory/implementations/inmemory"
	repositoryimppostgres "github.com/mehranhn/go_example_1/external/repositories/implementations/postgres"
	smsimpdummyconsole "github.com/mehranhn/go_example_1/external/sms/implementations/dummy-console"
	"github.com/mehranhn/go_example_1/services"
)

func getControllers(conf *config.AppConfig) (*controllers.AuthController, *controllers.UserController, error) {
	// extenrals
	db, err := repositoryimppostgres.NewPostgres(conf.DBConnection)
	if err != nil {
		return nil, nil, err
	}

	memory := memoryimpinmemory.NewInMemory()
	sms := smsimpdummyconsole.NewDummyConsole()

	// services
	authService := services.NewAuthService(
		&db,
		&memory,
		&sms,
		conf.RateLimitMaxAttempts,
		conf.RateLimitTTLDuration,
		conf.OTPTTLDuration,
		conf.JwtSecret,
		conf.JwtTTLDuration,
	)
	userService := services.NewUserService(&db)

	// controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)

	middlewares.DefaultJWTConfig(conf.JwtSecret)

	return &authController, &userController, nil
}

func root(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func healthCheck(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// @title User API
// @version 1.0
// @description This is a sample User API

// @contact.name API Support
// @contact.email mehranhhm@gmail.com

// @host localhost:3000
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	godotenv.Load(".env")

	conf, err := config.ReadConfig()
	if err != nil {
		panic(err.Error())
	}

	authController, userController, err := getControllers(conf)
	if err != nil {
		panic(err.Error())
	}

	jwtConfig := middlewares.DefaultJWTConfig(conf.JwtSecret)

	// this part could be improved by putting it into its seperated module but since this is a simple test crud app i'm not going to do it
	app := fiber.New()

	app.Use(logger.New())
	app.Use("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/", root)
	app.Get("/health", healthCheck)

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register-or-login", authController.RegistryOrLogin)
	auth.Post("/confirm-otp", authController.ConfirmOtp)

	user := api.Group("/user")
	user.Get("/", middlewares.JWTAuthMiddleware(jwtConfig), userController.GetList)
	user.Get("/:id", middlewares.JWTAuthMiddleware(jwtConfig), userController.GetUserByID)

	app.Listen(fmt.Sprintf(":%d", conf.Port))
}
