package jwt

import jwtLib "github.com/dgrijalva/jwt-go"

type Claims struct {
	// User id, used to retrieve from database.
	ID string `json:"id"`
	jwtLib.StandardClaims
}
