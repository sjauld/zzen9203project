package main

import (
	"crypto/rand"
	"log"
	"math/big"
)

// this method is not efficient for large euler totients
func bruteForceInverseModM(a, m int) int {
	for x := 1; ; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
}

func extendedEuclideanInverseModM(a, m *big.Int) *big.Int {
	oldR := big.NewInt(0).Set(a)
	r := big.NewInt(0).Set(m)
	oldS, s := big.NewInt(1), big.NewInt(0)
	oldT, t := big.NewInt(0), big.NewInt(1)

	for {
		quotient := big.NewInt(0).Div(oldR, r)

		oldR, r = r, big.NewInt(0).Mod(oldR, r)
		oldS, s = s, big.NewInt(0).Sub(oldS, big.NewInt(0).Mul(quotient, s))
		oldT, t = t, big.NewInt(0).Sub(oldT, big.NewInt(0).Mul(quotient, t))
		if r.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}

	if oldS.Cmp(big.NewInt(0)) == -1 {
		return big.NewInt(0).Add(oldS, m)
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

func findPrivateKey(p, q, e *big.Int) *big.Int {
	// euler totient
	m := big.NewInt(0).Mul(p.Sub(p, big.NewInt(1)), q.Sub(q, big.NewInt(1)))

	return extendedEuclideanInverseModM(e, m)
}
