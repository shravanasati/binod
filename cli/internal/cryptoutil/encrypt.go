package cryptoutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Encrypt encrypts the given text.
func Encrypt(key []byte, text string) ([]byte, error) {
	// * creating a cipher block
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	// * creating a new block of bytes
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	// * creating a nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// * encrypting the text
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)
	return ciphertext, nil
}

// func main() {
//     key := generateRandomKey()
//     text := "Hello, World!"
//     encryptedText, err := encrypt(key, text)
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("Encrypted:", string(encryptedText))

//     decryptedText, err := decrypt(key, encryptedText)
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("Decrypted:", decryptedText)
// }
