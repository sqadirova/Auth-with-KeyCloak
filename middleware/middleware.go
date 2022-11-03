package middleware

import (
	"AuthorizationWithKeycloak/keycloak"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v12"
	"log"
	"strings"
)

var newKeyCloak = keycloak.NewKeycloak()

func ExtractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func VerifyToken(accessToken string) (bool, string, error) {
	// extract Bearer token
	token := ExtractBearerToken(accessToken)

	if token == "" {
		return false, "", errors.New("Bearer Token missing")
	}

	//// call Keycloak API to verify the access token
	result, err := newKeyCloak.Gocloak.RetrospectToken(context.Background(), token, newKeyCloak.ClientId, newKeyCloak.ClientSecret, newKeyCloak.Realm)
	log.Println("result: ", result)

	if err != nil {
		return false, "", errors.New(fmt.Sprintf("Invalid or malformed token: %s", err.Error()))
	}

	jwt, _, err := newKeyCloak.Gocloak.DecodeAccessToken(context.Background(), token, newKeyCloak.Realm)
	log.Println("jwt: ", jwt)

	if err != nil {
		return false, "", errors.New(fmt.Sprintf("Invalid or malformed token: %s", err.Error()))

	}

	jwtToken, _ := json.Marshal(jwt)
	fmt.Printf("token: %v\n", string(jwtToken))

	// check if the token isn't expired and valid
	if !*result.Active {
		return false, "", errors.New("Invalid or expired Token")
	}

	return true, string(jwtToken), nil
}

func GetUserInfoFromToken(accessToken string) (*gocloak.UserInfo, error) {
	// extract Bearer token
	token := ExtractBearerToken(accessToken)

	if token == "" {
		return nil, errors.New("Bearer Token missing")
	}

	//// call Keycloak API to verify the access token
	result, err := newKeyCloak.Gocloak.RetrospectToken(context.Background(), token, newKeyCloak.ClientId, newKeyCloak.ClientSecret, newKeyCloak.Realm)
	log.Println("result: ", result)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid or malformed token: %s", err.Error()))
	}

	userInfo, err := newKeyCloak.Gocloak.GetUserInfo(context.Background(), token, newKeyCloak.Realm)
	if err != nil {
		return nil, err
	}

	fmt.Println("userInfo: ", userInfo)

	// check if the token isn't expired and valid
	if !*result.Active {
		return nil, errors.New("Invalid or expired Token")
	}

	return userInfo, nil
}
