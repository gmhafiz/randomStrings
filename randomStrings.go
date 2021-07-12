package randomStrings

import (
	"crypto/rand"
	"log"
)




const  (
	Alpha string = "alpha"
	Number string = "number"
	AlphaNum string = "alphanum"
	All string = "all"
)


// MODIFIED from https://socketloop.com/tutorials/golang-how-to-generate-random-string
func RandStr(strSize int, randType string) []byte {

	var dictionary string
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specialCharacters := "~@#$%^&*()_+|[]{}';\"?/.,><"

	switch randType {
	case Alpha:
		dictionary = alphabet
	case Number:
		dictionary = numbers
	case AlphaNum:
		dictionary = numbers + alphabet
	case All:
		dictionary = numbers + alphabet + specialCharacters
	default:
		return nil
	}

	var bytes = make([]byte, strSize)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return bytes
}

