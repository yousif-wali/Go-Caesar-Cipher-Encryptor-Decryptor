package main

import (
	"fmt"
	"strings"
)

// CaesarCipher struct represents the Caesar cipher with a shift
type CaesarCipher struct {
	Shift int
}

// NewCaesarCipher creates a new CaesarCipher with a given shift
func NewCaesarCipher(shift int) *CaesarCipher {
	return &CaesarCipher{Shift: shift}
}

// Encrypt encrypts the plaintext using the Caesar cipher
func (c *CaesarCipher) Encrypt(plaintext string) string {
	return c.shiftText(plaintext, c.Shift)
}

// Decrypt decrypts the ciphertext using the Caesar cipher
func (c *CaesarCipher) Decrypt(ciphertext string) string {
	return c.shiftText(ciphertext, -c.Shift)
}

// shiftText shifts each letter in the text by the given shift amount
func (c *CaesarCipher) shiftText(text string, shift int) string {
	shift = (shift%26 + 26) % 26 // Ensure the shift is within 0-25
	var shiftedText strings.Builder
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			// Shift lowercase letters
			shiftedChar := 'a' + (char-'a'+rune(shift))%26
			shiftedText.WriteRune(shiftedChar)
		} else if char >= 'A' && char <= 'Z' {
			// Shift uppercase letters
			shiftedChar := 'A' + (char-'A'+rune(shift))%26
			shiftedText.WriteRune(shiftedChar)
		} else {
			// Non-alphabetic characters are unchanged
			shiftedText.WriteRune(char)
		}
	}
	return shiftedText.String()
}

func main() {
	cipher := NewCaesarCipher(3) // Shift of 3
	plaintext := "Hello, World!"
	ciphertext := cipher.Encrypt(plaintext)
	fmt.Println("Encrypted:", ciphertext)

	decryptedText := cipher.Decrypt(ciphertext)
	fmt.Println("Decrypted:", decryptedText)
}
