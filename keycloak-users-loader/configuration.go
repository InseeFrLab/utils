package main

type keycloak struct {
	URL string
	Login string 
	Password string
	AdminRealm string
	TargetRealm string
}

type input struct {
	File string
}

type config struct {
	Keycloak keycloak
	Input input
}