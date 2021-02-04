package main

// Dosyayı yeniden isimlendirme ve taşıma.

import (
	"log"
	"os"
)

func main() {
	originalPath := "hi.txt"
	newPath := "./moved_folder/hi_renamed.txt"

	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
