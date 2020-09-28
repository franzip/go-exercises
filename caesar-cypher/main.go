package main

import "fmt"

var cypherKey = []rune("BACKENDARCHITECTURE")

var cypherLength int = len(cypherKey)

func isEncryptable(char rune) bool {
	return char >= 65 && char <= 90
}

func encryptMessage(message string) (encrypted string) {
	var sum rune
	var charToRune rune
	var cypherIdx int = 0
	for i := 0; i < len(message); i++ {
		charToRune = rune(message[i])
		if isEncryptable(charToRune) {
			sum = charToRune + cypherKey[cypherIdx%cypherLength]
			cypherIdx += 1
			encrypted += string((sum+26)%26 + 'A')
		} else {
			encrypted += string(message[i])
		}
	}
	return
}

func decryptMessage(message string) (decrypted string) {
	var sum rune
	var charToRune rune
	var cypherIdx int = 0
	for i := 0; i < len(message); i++ {
		charToRune = rune(message[i])
		if isEncryptable(charToRune) {
			sum = (charToRune - cypherKey[cypherIdx%cypherLength])
			cypherIdx += 1
			decrypted += string((sum+26)%26 + 'A')
		} else {
			decrypted += string(message[i])
		}
	}
	return
}

func main() {
	enc := encryptMessage("JOHN DOE IS A VERY STUPID GUY")
	fmt.Println(enc)
	fmt.Println(decryptMessage(enc))
	fmt.Println(decryptMessage("HO KC EJHSFOL, IGH TXUCPZ FWX XB SLRA DQML"))
}
