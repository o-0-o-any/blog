package models

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
