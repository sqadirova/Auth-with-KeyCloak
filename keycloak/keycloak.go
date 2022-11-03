package keycloak

import "github.com/Nerzal/gocloak/v12"

type Keycloak struct {
	Gocloak      *gocloak.GoCloak // keycloak client
	ClientId     string           // clientId specified in Keycloak
	ClientSecret string           // client secret specified in Keycloak
	Realm        string           // realm specified in Keycloak
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		Gocloak:      gocloak.NewClient("http://192.168.8.91:8080"),
		ClientId:     "my-go-service",
		ClientSecret: "An8Uofmv6DJDP63pNoYVouZsGLtEWqfF",
		Realm:        "test_authorization",
	}
}
