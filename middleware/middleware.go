package middleware

import (
	"AuthorizationWithKeycloak/keycloak"
	"context"
	"errors"
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
		return false, "", errors.New("bearer_token_missing")
	}

	// call Keycloak API to verify the access token
	result, err := newKeyCloak.Gocloak.RetrospectToken(context.Background(), token, newKeyCloak.ClientId, newKeyCloak.ClientSecret, newKeyCloak.Realm)
	log.Println("result: ", result)

	if err != nil {
		return false, "", errors.New("invalid_token")
	}

	//jwt, _, err := newKeyCloak.Gocloak.DecodeAccessToken(context.Background(), token, newKeyCloak.Realm)
	//log.Println("jwt: ", jwt)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//jwtToken, _ := json.Marshal(jwt)
	//fmt.Printf("jwtToken: %v\n", string(jwtToken))

	// check if the token isn't expired and valid
	if !*result.Active {
		return false, "", errors.New("token_expired")
	}

	return true, token, nil
}
