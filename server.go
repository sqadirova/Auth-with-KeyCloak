package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"AuthorizationWithKeycloak/config"
	_ "AuthorizationWithKeycloak/docs"
	"AuthorizationWithKeycloak/routes"
	swagger "github.com/gofiber/swagger"
	"log"
)

// @title Inventory Management API
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	//migrateService := migrate.NewMigrateService(config.DB)
	//err := migrateService.MigrateModelsToDB()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	v1 := app.Group("/api/v1")

	routeService := routes.NewFiberService(v1)
	routeService.Auth()
	routeService.User()
	//routeService.Role()
	routeService.SwaggerRoute()

	v1.Get("/swagger/*", swagger.HandlerDefault)

	err := app.Listen(fmt.Sprintf(":%v", config.InfraConfiguration.App.Port))

	if err != nil {
		log.Println(err)
	}
}
