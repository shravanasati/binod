package cryptoutil

import (
	"math/rand"
	"path/filepath"
	"strings"

	"github.com/shravanasati/binod/cli/internal/fsutil"
)

// generateRandomKey generates a random 32 bit signing key for all encryption-decryption tasks.
func generateRandomKey() []byte {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var b strings.Builder
	for i := 0; i < 32; i++ {
		b.WriteRune(letters[rand.Intn(len(letters))])
	}
	return []byte(b.String())
}

// getKey fetches the signing key from the local file system, and if it doesn't find one, generates one.
func getKey() []byte {
	rootDir := fsutil.GetBinodRootDir()
	keyLocation := filepath.Join(rootDir, "data", "key.dat")
	var key []byte

	// key doesnt exist
	if !fsutil.Exists(keyLocation) {
		key = generateRandomKey()
		fsutil.WriteToFile(string(key), keyLocation)

	} else { // key exists
		key = []byte(fsutil.ReadFile(keyLocation))
	}

	return key
}
