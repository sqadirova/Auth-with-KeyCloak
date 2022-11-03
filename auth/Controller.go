package auth

import (
	"AuthorizationWithKeycloak/keycloak"
	"AuthorizationWithKeycloak/middleware"
	"AuthorizationWithKeycloak/response"
	"github.com/go-playground/validator"
	fiber "github.com/gofiber/fiber/v2"
	"net/http"
)

var NewAuthService IAuthService
var validate *validator.Validate

func init() {
	NewAuthService = GetNewAuthService(keycloak.NewKeycloak())
	validate = validator.New()
}

func Get(ctx *fiber.Ctx) error {
	//get userID from request header
	isValid, token, err := middleware.VerifyToken(ctx.GetReqHeaders()["Authorization"])

	if err != nil || isValid == false {
		return ctx.Status(http.StatusUnauthorized).
			JSON(response.Response{
				Key:     err.Error(),
				Message: response.GetErrorResponse(err.Error()),
			})
	}

	userID, statusCode, err := NewAuthService.signOut(userId)

	if err != nil {
		return ctx.Status(statusCode).
			JSON(response.Response{
				Key:     err.Error(),
				Message: response.GetErrorResponse(err.Error()),
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(SignOutDTO{UserID: userID})
}

// SignIn godoc
// @Summary Sign in the user to system.
// @Tags auth
// @Accept json
// @Produce json
// @Param input   body  SignInReq   true  "Sign In Req"
// @Success 200 {object} SignInResp
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /auth/sign-in [post]
func SignIn(ctx *fiber.Ctx) error {
	var userReqBody SignInReq

	if err := ctx.BodyParser(&userReqBody); err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(response.Response{ //nolint:wsl
			Key:     "invalid_data",
			Message: response.GetErrorResponse("invalid_data"),
		})
	}

	ctx.Set("Content-Type", "application/json")

	err := validate.Struct(userReqBody)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(response.Response{
				Key:     "invalid_data",
				Message: response.GetErrorResponse("invalid_data")})
	}

	token, statusCode, err := NewAuthService.signIn(userReqBody)

	if err != nil {
		return ctx.Status(statusCode).
			JSON(response.Response{
				Key:     err.Error(),
				Message: response.GetErrorResponse(err.Error()),
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(SignInResp{Token: token})
}

//// SignOut godoc
//// @Summary Sign out the user from system.
//// @Tags auth
//// @Accept json
//// @Produce json
//// @Param Authorization header string true "Bearer"
//// @Success 200 {object} SignOutDTO
//// @Failure 400 {object} response.Response
//// @Router  /auth/sign-out [post]
//func SignOut(ctx *fiber.Ctx) error {
//	//get userID from request header
//	isValid, token, err := middleware.VerifyToken(ctx.GetReqHeaders()["Authorization"])
//
//	if err != nil || isValid == false {
//		return ctx.Status(http.StatusUnauthorized).
//			JSON(response.Response{
//				Key:     err.Error(),
//				Message: response.GetErrorResponse(err.Error()),
//			})
//	}
//
//	userID, statusCode, err := NewAuthService.signOut(userId)
//
//	if err != nil {
//		return ctx.Status(statusCode).
//			JSON(response.Response{
//				Key:     err.Error(),
//				Message: response.GetErrorResponse(err.Error()),
//			})
//	}
//
//	return ctx.Status(fiber.StatusOK).JSON(SignOutDTO{UserID: userID})
//}
