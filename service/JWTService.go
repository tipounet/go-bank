package service

import (
	"fmt"
	"log"
	"time"

	"github.com/tipounet/go-bank/configuration"
	"github.com/tipounet/go-bank/model"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	hmacSampleSecret []byte
)

func init() {
	hmacSampleSecret = []byte(configuration.GetConfiguration().JWT)
}

// JWTService Objet d'accès au chose consernant jwt
type JWTService struct{}

// GenerateToken : création du token JWT à partir de l'email fournit
// From : https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
// func (JWTService) GenerateToken(userMail string) (retour string) {
func (JWTService) GenerateToken(user model.User) (retour string) {
	claims := UserClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(20 * time.Minute).Unix(),
			Issuer:    "bankApp",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(hmacSampleSecret)
	log.Printf("Erreur %v\n", err)
	return ss
}

// UserClaims : structure contenant les infos que l'on passe dans le token jwt
type UserClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

// ParseToken : vérifie et récupère l'information contenu dans le jeton
func (JWTService) ParseToken(tokenString string) (user model.User, ok bool) {

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, validToken := token.Method.(*jwt.SigningMethodHMAC); !validToken {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	// attention ici validToken c'est pas le scope global c'est le scope du if et donc si on fait juste un return ça retourne le ok de la fonction qui lui vaut forcément false ici
	// donc on peu pas l'appeler ok
	if claims, validToken := token.Claims.(*UserClaims); validToken && token.Valid {
		user = claims.User
		ok = true
	} else {
		log.Println(err)
	}
	return
}
