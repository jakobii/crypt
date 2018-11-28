package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Encrypt will encrypt and byte array into with aes encrption with the supplied key.
func Encrypt(key []byte, object []byte) (b []byte, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	b = make([]byte, aes.BlockSize+len(object))
	o := b[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, o)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, o)
	stream.XORKeyStream(b[aes.BlockSize:], object)
	return b, nil
}
