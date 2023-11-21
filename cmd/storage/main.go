package main

import (
	"fmt"
	"github.com/maksimbb52/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It's uploaded!", file)
}
