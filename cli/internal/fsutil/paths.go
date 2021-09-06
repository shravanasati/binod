package fsutil

import (
	"log"
	"os"
	"path/filepath"
)

func Exists(name string) bool {
	_, e := os.Stat(name)
	return os.IsExist(e)
}

// GetBinodRootDir returns the path to ~/.binod. Creates the directory if they don't exist.
func GetBinodRootDir() string {
	dir, e := os.UserHomeDir()
	if e != nil {
		log.Fatalf("Unable to get the binod root dir! \n%v", e)
	}

	rootDir := filepath.Join(dir, ".binod")
	if !Exists(rootDir) {
		os.Mkdir(rootDir, os.ModePerm)
	}

	subDirs := []string{"data"}
	for _, v := range subDirs {
		if !Exists(filepath.Join(rootDir, v)) {
			os.Mkdir(rootDir, os.ModePerm)
		}
	}

	return rootDir
}
