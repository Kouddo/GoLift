package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
)

// TODO implement better error handling
func getNewToken(key []byte) string {

	t := jwt.New(jwt.SigningMethodHS256)
	s, err := t.SignedString(key)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	return s
}

// TODO implement this function
func validateToken(tokenString string, key []byte) bool {
	return false
}
