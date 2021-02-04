package main

// Dosya bilgilerini almak.

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	fileInfo, err = os.Stat("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Folder? (Dictionary?):", fileInfo.IsDir())
	fmt.Println("System Interface Type:", fileInfo.Sys())
}
