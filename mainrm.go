package main

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func maina() {
	salt := "du sel"

	dk, a := scrypt.Key([]byte("some password"), []byte(salt), 16384, 8, 1, 32)

	fmt.Printf("dk : %T %v | a : %T %v", dk, dk, a, a)
}
