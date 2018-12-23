package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	key := []byte("0123456789012345")
	iv := []byte("9876543210987654")
	msg := []byte("Hello AES, my old friend")

	enc, err := Encrypt(msg, key, iv)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", enc)
}

// Encrypt encrypts the plaintext with key and iv using AES-CFB and returns it in base64.
func Encrypt(plaintext, key, iv []byte) ([]byte, error) {
	// Create ciphertext.
	ciphertext := make([]byte, len(plaintext))
	// Create AES cipher.
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// Get AES-CFB stream encrypted.
	stream := cipher.NewCFBEncrypter(aesBlock, iv)
	// Encrypt the msg and store the results in ciphertext.
	stream.XORKeyStream(ciphertext, plaintext)
	// Base64 encode it.
	encodedCiphertext := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(encodedCiphertext, ciphertext)
	return encodedCiphertext, nil
}
