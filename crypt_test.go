package crypt

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	k := []byte("0123456789012345")
	m := []byte("A quick brown fox jumped over the lazy dog.")

	e, err := Encrypt(k, m)
	if err != nil {
		t.Error(err)
	}
	log.Printf("enrypted: %s\n", base64.StdEncoding.EncodeToString(e))
	log.Printf("Key: %s\n", string(k))

	d, err := Decrypt(k, e)
	if err != nil {
		t.Error(err)
	}
	log.Printf("Decrypted: %s\n", string(d))
}
