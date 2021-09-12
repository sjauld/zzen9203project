// package publickey generates public keys of a given strength
package publickey

import (
	"crypto/rand"
	"log"
	"math/big"
)

// PublicKey represents the encrypt half of an RSA keypair
type PublicKey struct {
	E int
	N int
}

// New generates and returns a new publickey of a given size s
func New(s int) {

}

// this method is not efficient for large euler totients
func bruteForceInverseModM(a, m int) int {
	for x := 1; ; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
}

func extendedEuclideanInverseModM(a, m int) int {
	oldR, r := a, m
	oldS, s := 1, 0
	oldT, t := 0, 1

	for {
		quotient := oldR / r
		oldR, r = r, oldR%r
		oldS, s = s, oldS-quotient*s
		oldT, t = t, oldT-quotient*t
		if r == 0 {
			break
		}
	}

	if oldS < 0 {
		return oldS + m
	}
	return oldS
}

// ref: https://www.geeksforgeeks.org/check-two-numbers-co-prime-not/
func gcd(a, b int) int {
	// everything divides 0
	if a == 0 || b == 0 {
		return 0
	}

	if a == b {
		return a
	}

	if a > b {
		return gcd(a-b, b)
	}

	return gcd(a, b-a)
}

func findCoprime(i int) int {
	log.Printf("[DEBUG] finding coprime of %d", i)
	for {
		// just try random numbers until we find one
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(i)))
		log.Printf("[DEBUG] trying %d", n)
		if gcd(i, int(n.Int64())) == 1 {
			return int(n.Int64())
		}
	}
}

func generateKeys(p, q int) (n, m, d, e int) {
	// modulus
	n = p * q
	log.Printf("[DEBUG] modulus %d", n)
	// euler totient
	m = (p - 1) * (q - 1)
	log.Printf("[DEBUG] Euler totient %d", m)
	// first key
	a := findCoprime(m)
	// second key
	b := extendedEuclideanInverseModM(a, m)
	// the smallest key should be the encryption key
	if a > b {
		d = a
		e = b
	} else {
		d = b
		e = a
	}
	log.Printf("[DEBUG] decryption key %d", d)
	log.Printf("[DEBUG] encryption key %d", e)
	return
}
