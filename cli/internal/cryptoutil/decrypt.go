package cryptoutil

import (
	"crypto/aes"
	"crypto/cipher"
)

// Decrypt decrypts the given text.
func Decrypt(key, text []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(text) < nonceSize {
		return "", err
	}

	nonce, text := text[:nonceSize], text[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, text, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
