package authentication

// From : https://trihackeat.wordpress.com/2014/10/11/758/
// This will handle all aspects of authenticating users in our system
// For password managing/salting I used:
// http://austingwalters.com/building-a-web-server-in-go-salting-passwords/

import (
	"crypto/rand"
	"errors"
	"log"
	"strings"

	"github.com/tipounet/go-bank/stringUtils"

	"golang.org/x/crypto/bcrypt"
)

const (
	// MaxSaltLength : la taille max du sel
	MaxSaltLength = 56
	// EncryptCost  On a scale of 3 - 31, how intense Bcrypt should be
	EncryptCost = 14
)

// Password This is returned when a new hash + salt combo is generated
type Password struct {
	Hash string
	Salt string
}

func (p Password) String() string {
	return "Password : {\n\tsalt : " + p.Salt + "\n\thash : " + p.Hash + "}"
}

// this handles taking a raw user password and making in into something safe for
// storing in our DB
func hashPassword(saltedPass string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(saltedPass), EncryptCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPass)
}

// Handles merging together the salt and the password
func combine(salt string, rawPass string) string {
	// concat salt + password
	pieces := []string{salt, rawPass}
	// Attention : changer la glue casse les mdp existant. Il faut prévoir un mode pour initialiser la base !
	saltedPassword := strings.Join(pieces, "")
	return saltedPassword
}

// Generates a random salt using DevNull
func generateSalt(pwdLen int) string {
	length := MaxSaltLength
	if pwdLen <= MaxSaltLength {
		length = MaxSaltLength - pwdLen
	}

	salt, _ := stringUtils.RandGen(length, stringUtils.All, "", "")
	return salt
}

// Generates a random salt using DevNull
func generateSaltOld() string {
	// Read in data
	data := make([]byte, MaxSaltLength)
	_, err := rand.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// Convert to a string
	salt := string(data[:])
	return salt
}

//CreatePassword Handles create a new hash/salt combo from a raw password as inputted
// by the user
func CreatePassword(rawPass string) (*Password, error) {
	if len(rawPass) > 64 {
		log.Printf("Le mot de passe est trop long (oui c'est parfois une question de taille :/) : %v", rawPass)
	} else {
		password := new(Password)
		password.Salt = generateSalt(len(rawPass))
		saltedPass := combine(password.Salt, rawPass)
		password.Hash = hashPassword(saltedPass)
		return password, nil
	}
	return nil, errors.New("La mot de passe doit faire au maximum 64 caractères")
}

// PasswordMatch Checks whether or not the correct password has been provided
func PasswordMatch(guess string, password *Password) (ok bool, err error) {
	saltedGuess := combine(password.Salt, guess)
	// compare to the real deal
	if err = bcrypt.CompareHashAndPassword([]byte(password.Hash), []byte(saltedGuess)); err != nil {
		ok = false
	} else {
		ok = true
	}
	return
}
