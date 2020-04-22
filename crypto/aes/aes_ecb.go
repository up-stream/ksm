package aes

import (
	"crypto/aes"
	"errors"
)

func EncryptWithECB(key []byte, plainText []byte) ([]byte, error) {
	blockSize := aes.BlockSize
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(plainText)%blockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	encrypted := make([]byte, len(plainText))
	dst := encrypted
	for len(plainText) > 0 {
		c.Encrypt(dst, plainText[:blockSize])
		plainText = plainText[blockSize:]
		dst = dst[blockSize:]
	}

	return encrypted, nil
}
