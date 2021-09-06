package fsutil

import (
	"bufio"
	"log"
	"os"
)

// WriteToFile writes the given content to the given filepath.
func WriteToFile(content, path string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to write to file %s\n%v", path, err)
	}

	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		log.Fatalf("unable to write to file %s\n%v", path, err)
	}
}

// ReadFile reads the given file.
func ReadFile(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("unable to real file %s\n%v", file, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	text := ""
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	return text
}
