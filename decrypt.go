package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// Decrypt decrypts an aes byte array with the provided key.
func Decrypt(key []byte, n []byte) (b []byte, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(n) < aes.BlockSize {
		err = errors.New("n block size is too short")
		return nil, err
	}

	k := n[:aes.BlockSize]
	n = n[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, k)
	stream.XORKeyStream(n, n)
	b = n
	return b, nil
}
