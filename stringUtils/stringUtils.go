// Package stringUtils : from https://github.com/elgs/gostrgen
package stringUtils

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// Lower permet d'indiquer si l'on souhaite utiliser les lettres minuscules
var Lower = 1 << 0

// Upper permet d'indiquer si l'on souhaite utiliser les lettres majuscules
var Upper = 1 << 1

// Digit permet d'indiquer si l'on souhaite utiliser les chiffre
var Digit = 1 << 2

// Punct permet d'indiquer si l'on souhaite utiliser la ponctuation
var Punct = 1 << 3

// LowerUpper : indique l'on souhaite utiliser les lettres minucules et majuscules (raccourcis pour Lower | Upper)
var LowerUpper = Lower | Upper

// LowerDigit : indique l'on souhaite utiliser les lettres minucules et les chiffres (raccourcis pour Lower | Digit)
var LowerDigit = Lower | Digit

// UpperDigit indique l'on souhaite utiliser les lettres majuscules et les chiffres (raccourcis pour Upper | Digit)
var UpperDigit = Upper | Digit

// LowerUpperDigit indique que l'on souhaite utiliser les chiffres et les lettres minucules et majuscules (raccourcis pour LowerUpper | Digit)
var LowerUpperDigit = LowerUpper | Digit

// All : indique que l'on souhaites utilisers les lettres minucules et majuscules ainsi que les chiffres et la ponctuation
var All = LowerUpperDigit | Punct

var lower = "abcdefghijklmnopqrstuvwxyz"
var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digit = "0123456789"
var punct = "~!@#$%^&*()_+-="

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandGen : Génération d'une chaîne alléatoire.
// size : taille de la chaine finale
// set : caractère à utiliser, en fonction des "set" prédéfinie ci dessus, c'est un "ou" binaire entre chaque
// includes (string) : caractère supplémentaire à inclure (qui ne sont pas dans les set existants)
// exclude (string) caractère à exclure (par exemple o majuscule pour pas confondre avec zéro, ou l minuscule pour pas confondre avec un)
func RandGen(size int, set int, include string, exclude string) (string, error) {
	all := include
	if set&Lower > 0 {
		all += lower
	}
	if set&Upper > 0 {
		all += upper
	}
	if set&Digit > 0 {
		all += digit
	}
	if set&Punct > 0 {
		all += punct
	}

	lenAll := len(all)
	if len(exclude) >= lenAll {
		return "", errors.New("Too much to exclude.")
	}
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		b := all[rand.Intn(lenAll)]
		if strings.Contains(exclude, string(b)) {
			i--
			continue
		}
		buf[i] = b
	}
	return string(buf), nil
}
