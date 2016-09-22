package service

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	hmacSampleSecret = []byte("s3cr3tK3y à moi tout seul que je connais !")
)

// JWTService Objet d'accès au chose consernant jwt
type JWTService struct{}

// GenerateToken : création du token JWT à partir de l'email fournit
// From : https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
func (JWTService) GenerateToken(userMail string) (retour string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usermail": userMail,
		// TODO : now ?
		"nbf": time.Date(1981, 9, 24, 23, 2, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	retour, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}
	return
}

// ParseToken : vérifie et récupère l'information contenu dans le jeton
// TODO : ben a faire ;)
func (JWTService) ParseToken(token string) (userMail string, ok bool) {
	return
}
