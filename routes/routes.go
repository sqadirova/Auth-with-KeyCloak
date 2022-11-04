package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	//"AuthorizationWithKeycloak/CONSTANTS"
	"AuthorizationWithKeycloak/auth"
	"AuthorizationWithKeycloak/config"
	"AuthorizationWithKeycloak/docs"
	//"AuthorizationWithKeycloak/middleware"
	//"AuthorizationWithKeycloak/user"
)

type Interface interface {
	Auth()
	User()
	SwaggerRoute()
	//Role()
}

type RouteService struct {
	Server fiber.Router
}

func NewFiberService(server fiber.Router) Interface {
	return &RouteService{Server: server}
}

func (a *RouteService) Auth() {
	a.Server.Post("/auth/sign-in", auth.SignIn)
	a.Server.Post("/auth/testKeycloak", auth.TestKeycloak)
	a.Server.Post("/auth/sign-out", auth.SignOut)
}

//	func (a *RouteService) Role() {
//		a.Server.Get("/user/roles", middleware.IsAuthorized(CONSTANTS.ADMIN), user.GetAllRoles)
//		a.Server.Get("/user/roles/:id", middleware.IsAuthorized(CONSTANTS.ADMIN), user.GetOneRole)
//	}
func (a *RouteService) User() {
	//a.Server.Get("/users", middleware.IsAuthorized(CONSTANTS.ADMIN), user.GetAllUsers)
	//a.Server.Post("/user", middleware.IsAuthorized(CONSTANTS.ADMIN), user.CreateUser)
	//a.Server.Put("/user/:id", middleware.IsAuthorized(CONSTANTS.ADMIN), user.UpdateUser)
	a.Server.Get("/user/me", auth.GetUserMe)
}

func (r *RouteService) SwaggerRoute() {
	docs.SwaggerInfo.Host = config.InfraConfiguration.Swagger.Host
	docs.SwaggerInfo.Schemes = []string{config.InfraConfiguration.Swagger.Scheme}

	r.Server.Get("/swagger/*", swagger.HandlerDefault) // default
}
