package auth

import (
	"AuthorizationWithKeycloak/keycloak"
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v12"
	"log"
	"net/http"
)

type IAuthService interface {
	signOut(refreshToken string) (bool, int, error)
	signIn(userReqBody SignInReq) (*SignInResp, int, error)
	testKeyCloakService(token string) ([]*gocloak.Role, error)
	getUserInfoFromToken(token string) (*UserMeResp, error)
	//testKeyCloakService(token string) ([]*gocloak.User, error)
}

type AuthService struct {
	keycloak *keycloak.Keycloak
}

func GetNewAuthService(keycloak *keycloak.Keycloak) IAuthService {
	return &AuthService{keycloak: keycloak}
}

//func (auth *AuthService) testKeyCloakService(token string) ([]*gocloak.User, error) {
//	log.Println("token: ", token)
//	enabled := true
//	users, err := auth.keycloak.Gocloak.GetUsers(context.Background(), token, auth.keycloak.Realm, gocloak.GetUsersParams{Enabled: &enabled})
//	fmt.Println("users: ", users)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	return users, nil
//}

func (auth *AuthService) testKeyCloakService(token string) ([]*gocloak.Role, error) {
	log.Println("token: ", token)
	//enabled := true
	//users, err := auth.keycloak.Gocloak.GetUsers(context.Background(), token, auth.keycloak.Realm, gocloak.GetUsersParams{Enabled: &enabled})
	//fmt.Println("users: ", users)
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}

	role, err := auth.keycloak.Gocloak.GetRealmRoles(context.Background(), token, auth.keycloak.Realm, gocloak.GetRoleParams{})
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (auth *AuthService) signOut(refreshToken string) (bool, int, error) {
	err := auth.keycloak.Gocloak.Logout(context.Background(), auth.keycloak.ClientId, auth.keycloak.ClientSecret, auth.keycloak.Realm, refreshToken)

	if err != nil {
		log.Println(err)
		return false, http.StatusForbidden, errors.New("unexpected_error")
	}

	return true, 0, nil
}

func (auth *AuthService) signIn(userReqBody SignInReq) (*SignInResp, int, error) {
	jwt, err := auth.keycloak.Gocloak.Login(context.Background(),
		auth.keycloak.ClientId,
		auth.keycloak.ClientSecret,
		auth.keycloak.Realm,
		userReqBody.Username,
		userReqBody.Password)

	if err != nil {
		return nil, http.StatusForbidden, err
	}

	fmt.Println("jwt.RefreshToken: ", jwt.RefreshToken)
	fmt.Println("jwt.IDToken: ", jwt.IDToken)

	return &SignInResp{AccessToken: jwt.AccessToken, RefreshToken: jwt.RefreshToken}, 0, nil
}

func (auth *AuthService) getUserInfoFromToken(token string) (*UserMeResp, error) {
	userInfo, err := auth.keycloak.Gocloak.GetUserInfo(context.Background(), token, auth.keycloak.Realm)
	if err != nil {
		return nil, errors.New("unexpected_error")
	}

	fmt.Println("userInfo: ", userInfo)

	roles, err := auth.keycloak.Gocloak.GetRealmRolesByUserID(context.Background(), token, auth.keycloak.Realm, *userInfo.Sub)
	if err != nil {
		return nil, errors.New("unexpected_error")
	}

	log.Println("roles: ", roles)

	var userRole *gocloak.Role
	for _, role := range roles {
		if *role.Composite == false {
			userRole = role
		}
	}

	fmt.Println("userRole: ", userRole)

	return &UserMeResp{
			Id:        *userInfo.Sub,
			Firstname: *userInfo.GivenName,
			Lastname:  *userInfo.FamilyName,
			Role: RolesResp{
				Id:       *userRole.ID,
				RoleType: *userRole.Name,
			},
			Username: *userInfo.PreferredUsername,
		},
		nil
}
