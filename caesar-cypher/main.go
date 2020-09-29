package main

import (
	"fmt"
)

func isEncryptable(char rune) bool {
	return char >= 65 && char <= 90
}

func (c *Cipher) shift(message string, forward bool) (result string) {
	var keyIdx int = 0
	var charToRune rune

	for _, char := range message {
		charToRune = rune(char)
		if isEncryptable(charToRune) {
			if forward {
				charToRune = ((charToRune + c.key[keyIdx%len(c.key)] + 26) % 26) + 'A'
			} else {
				charToRune = ((charToRune - c.key[keyIdx%len(c.key)] + 26) % 26) + 'A'
			}
			keyIdx += 1
		}

		result += string(charToRune)
	}
	return
}

func (c *Cipher) EncryptMessage(message string) (encrypted string) {
	return c.shift(message, true)
}

func (c *Cipher) DecryptMessage(message string) (decrypted string) {
	return c.shift(message, false)
}

type Cipher struct {
	key []rune
}

func main() {
	key := []rune("BACKENDARCHITECTURE")
	cipher := Cipher{key: key}
	enc := cipher.EncryptMessage("OWNED BY JOHN DOE")
	fmt.Println(enc)
	fmt.Println(cipher.DecryptMessage(enc))
	fmt.Println(cipher.DecryptMessage("HO KC EJHSFOL, IGH TXUCPZ FWX XB SLRA DQML"))
}
