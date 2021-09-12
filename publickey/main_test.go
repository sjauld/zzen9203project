package publickey

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"testing"
)

func TestGoImplementation(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 24)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("[BOOM]")

	log.Printf("%+v", key.Public())
}
