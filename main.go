package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()

		if file.IsDir() {
			fileName += "/"
		}

		fmt.Println(fileName)
	}
}
