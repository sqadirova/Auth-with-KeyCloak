package auth

import (
	"AuthorizationWithKeycloak/keycloak"
	"context"
	"net/http"
)

type IAuthService interface {
	//signOut(id string) (string, int, error)
	//isValid(enteredPassword, userPassword string) error
	signIn(userReqBody SignInReq) (string, int, error)
}

type AuthService struct {
	keycloak *keycloak.Keycloak
}

func GetNewAuthService(keycloak *keycloak.Keycloak) IAuthService {
	return &AuthService{keycloak: keycloak}
}

//func (a *AuthService) signOut(id string) (string, int, error) {
//	_, err := uuid.FromString(id)
//
//	if err != nil {
//		return "", 404, response.ErrInvalidId
//	}
//
//	user, err := userPack.NewUserService.GetUserByID(id)
//
//	if err != nil {
//		return "", 404, err
//	}
//
//	user.AccessToken = nil
//
//	err = a.repository.updateAccessToken(user)
//
//	if err != nil {
//		return "", 500, response.ErrUnexpected
//	}
//
//	return user.ID.String(), 0, nil
//}

func (auth *AuthService) signIn(userReqBody SignInReq) (string, int, error) {
	jwt, err := auth.keycloak.Gocloak.Login(context.Background(),
		auth.keycloak.ClientId,
		auth.keycloak.ClientSecret,
		auth.keycloak.Realm,
		userReqBody.Username,
		userReqBody.Password)

	if err != nil {
		return "", http.StatusForbidden, err
	}

	return jwt.AccessToken, 0, nil
}

//func (a *AuthService) getSignedJWTToken(userID uuid.UUID) (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//	claims["id"] = userID
//	claims["iat"] = time.Now().Unix()
//	claims["exp"] = time.Now().Unix() +
//		config.InfraConfiguration.Jwt.AccessTokenExpire
//	tokenString, err := token.SignedString(
//		[]byte(config.InfraConfiguration.Jwt.SecretKey))
//
//	if err != nil {
//		log.Println(err)
//		return "", err
//	}
//
//	return tokenString, nil
//}

//func (a *AuthService) isValid(enteredPassword, userPassword string) error {
//	passwordIsValid := bcrypt.CompareHashAndPassword([]byte(userPassword),
//		[]byte(enteredPassword))
//
//	if passwordIsValid != nil {
//		return errors.New("incorrect_password")
//	}
//
//	return nil
//}
