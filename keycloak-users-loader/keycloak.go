package main

import (
	"context"

	"github.com/Nerzal/gocloak/v8"
)

var client gocloak.GoCloak;

func authenticateToKeycloak() *gocloak.JWT {
	client = gocloak.NewClient(conf.Keycloak.URL)
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, conf.Keycloak.Login, conf.Keycloak.Password, conf.Keycloak.AdminRealm)
	if err != nil {
		panic(err)
	}
	return token
}

func createUser(token *gocloak.JWT, firstName string, lastName string, email string, username string) {
	ctx := context.Background()
	user := gocloak.User{
		FirstName: gocloak.StringP(firstName),
		LastName:  gocloak.StringP(lastName),
		Email:     gocloak.StringP(email),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP(username),
		RequiredActions: requiredActions(),
	}
	
	_, err := client.CreateUser(ctx, token.AccessToken, conf.Keycloak.TargetRealm, user)
	if err != nil {
		panic(err)
	}
}