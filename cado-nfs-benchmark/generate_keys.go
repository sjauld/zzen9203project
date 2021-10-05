package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"os"
)

// You can use this to create a bunch of keys
func generateKeys() {
	f, err := os.Create("keys3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for n := 256; n < 330; n++ {
		key, err := rsa.GenerateKey(rand.Reader, n)
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}

		line := fmt.Sprintf("%d,%d,%d\n", n, key.E, key.N)
		log.Println(line)

		f.WriteString(line)
	}
}
