package main

import (
	"testing"
)

func TestCypher(t *testing.T) {
	key := []rune("ASDQWEZXC")
	cipher := Cipher{key: key}
	message := "OWNED BY JOHN DOE"
	encryptedMessage := cipher.EncryptMessage(message)
	decryptedMessage := cipher.DecryptMessage(encryptedMessage)

	if encryptedMessage == message {
		t.Error("Error with EncryptMessage")
	}
	if decryptedMessage != message {
		t.Errorf("Error with DecryptMessage: got %s, expected %s", decryptedMessage, message)
	}

	key = []rune("BACKENDARCHITECTURE")
	encryptedMessage = "HO KC EJHSFOL, IGH TXUCPZ FWX XB SLRA DQML"
	expected := "GO IS AWESOME, AND REALLY FUN TO PLAY WITH"
	cipher = Cipher{key: key}
	decryptedMessage = cipher.DecryptMessage(encryptedMessage)
	if decryptedMessage != expected {
		t.Errorf("Got %s, expected %s", decryptedMessage, expected)
	}

}
