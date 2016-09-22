package authentication

import (
	"fmt"
	"testing"
)

func TestSpec(t *testing.T) {

	salt := generateSalt()
	if len(salt) != SaltLength {
		fmt.Printf("le sel ne fait pas la bonne taille actual %v, expected %v\n\n", len(salt), SaltLength)
	}
	p := "boomchuckalucka"
	password := CreatePassword(p)
	fmt.Printf("CreatePassword : %v\n\n", password)
	passStruct := new(Password)
	fmt.Printf("PAss struct %v\n\n", passStruct)

}
