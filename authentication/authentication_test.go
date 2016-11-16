package authentication

import (
	"log"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
)

func TestGenerateSalt(t *testing.T) {
	salt := generateSalt(0)
	if len(salt) != MaxSaltLength {
		t.Fatalf("le sel ne fait pas la bonne taille actual %v, expected %v\n\n", len(salt), MaxSaltLength)
	}
}

func TestCreatePassword(t *testing.T) {
	p := "boomchuckalucka"
	password, e := CreatePassword(p)
	if e != nil {
		t.Errorf("L'erreur %v est inattendue", p)
	}
	if password.Salt == "" {
		t.Errorf("Le grain de sel ne devrait pas être vide (%v)", password)
	}
	passStruct := new(Password)
	if passStruct.Salt != "" {
		t.Errorf("Le grain de sel  devrait être vide")
	}
}

func TestPasswordMatch(t *testing.T) {
	p := "boomchuckalucka"
	password, e := CreatePassword(p)
	if e != nil {
		t.Errorf("L'erreur %v est inattendue", p)
	}
	if ok, _ := PasswordMatch(p, password); !ok {
		t.Errorf("Le mot de passe doit match %v\n", p)
	}

	failedPwd := []string{"boomchuckalucka4", "bboomchuckalucka", "boomchuckalucka42", "_boomchuckalucka4_", "-boomchuckalucka4", "azerty"}
	for _, pwd := range failedPwd {
		if ok, _ := PasswordMatch(pwd, password); ok {
			t.Errorf("Le mot de passe ne doit pas correspondre %v != %v\n\n", p, pwd)
		}
	}

	if ok, _ := PasswordMatch(p, &Password{
		Hash: password.Hash,
		Salt: password.Salt,
	}); !ok {
		t.Errorf("Comparaison du mot de passe avec lui même ko (%v)", p)
	}
}

//
func TestSpec(t *testing.T) {

	Convey("Authentication Testing", t, func() {
		Convey("generateSalt()", func() {
			salt := generateSalt(0)
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, MaxSaltLength)

			salt = generateSalt(10)
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, MaxSaltLength-10)
		})

		Convey("combine()", func() {
			password := "boomchuckalucka"
			salt := generateSalt(len(password))
			expectedLength := len(salt) + len(password)
			combo := combine(salt, password)

			So(combo, ShouldNotBeBlank)
			So(len(combo), ShouldEqual, expectedLength)
			So(strings.HasPrefix(combo, salt), ShouldBeTrue)
		})

		Convey("hashPassword()", func() {
			pwd := "hershmahgersh"
			combo := combine(generateSalt(len(pwd)), pwd)
			hash := hashPassword(combo)
			So(hash, ShouldNotBeBlank)

			cost, err := bcrypt.Cost([]byte(hash))
			if err != nil {
				log.Print(err)
			}
			So(cost, ShouldEqual, EncryptCost)
		})

		Convey("CreatePassword()", func() {
			passString := "mmmPassword1"
			password, e := CreatePassword(passString)
			So(e, ShouldBeNil)

			passStruct := new(Password)
			So(password, ShouldHaveSameTypeAs, passStruct)
			So(password.Hash, ShouldNotBeBlank)
			So(password.Salt, ShouldNotBeBlank)
			So(len(password.Salt), ShouldEqual, MaxSaltLength-len(passString))
		})

		Convey("comparePassword", func() {
			password := "megaman49"
			passwordMeta, e := CreatePassword(password)
			So(e, ShouldBeNil)

			ok, _ := PasswordMatch(password, passwordMeta)
			So(ok, ShouldBeTrue)
			ok, _ = PasswordMatch("lolfail", passwordMeta)
			So(ok, ShouldBeFalse)
			ok, _ = PasswordMatch("Megaman49", passwordMeta)
			So(ok, ShouldBeFalse)
			ok, _ = PasswordMatch("megaman40", passwordMeta)
			So(ok, ShouldBeFalse)
			ok, _ = PasswordMatch("megaman48", passwordMeta)
			So(ok, ShouldBeFalse)
		})
		Convey("compare long Password", func() {
			password := "megaman490megaman490megaman490megaman490megaman490megaman490megaman490"
			passwordMeta, e := CreatePassword(password)
			So(passwordMeta, ShouldBeNil)
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldEqual, "La mot de passe doit faire au maximum 64 caractères")
		})
	})
}
