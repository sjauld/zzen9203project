package rsacrack

import (
	"context"
	"crypto/rsa"
	"log"
	"math/big"
	"sync"
)

var bigTwo = big.NewInt(2)
var bigOne = big.NewInt(1)
var bigZero = big.NewInt(0)
var bigMinusOne = big.NewInt(-1)
var bigMinusTwo = big.NewInt(-2)

func guessFactor(c chan *rsa.PrivateKey, mu *sync.Mutex, ctx context.Context, pubkey *rsa.PublicKey, guess *big.Int) {
	for {
		// Wait for a lock, check for cancellation, then copy the current value of our guess and decrement pointer
		mu.Lock()
		select {
		case <-ctx.Done():
			mu.Unlock()
			return
		default:
			thisGuess := big.NewInt(0).Set(guess)
			guess.Sub(guess, bigTwo)
			mu.Unlock()

			if guess.Cmp(bigTwo) == -1 {
				// Factor not found :(
				log.Printf("[DEBUG] factor not found")
				c <- nil
				return
			}

			if big.NewInt(0).Mod(pubkey.N, thisGuess).Cmp(bigZero) == 0 {
				c <- calculatePrivateKey(pubkey, thisGuess)
				return
			}
		}
	}
}

func CrackPublicKey(pubkey *rsa.PublicKey, routines int) *rsa.PrivateKey {
	c := make(chan *rsa.PrivateKey)
	var mu sync.Mutex
	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	start := big.NewInt(0)
	start.Sqrt(pubkey.N)
	// Make it odd
	if start.Bit(0) == 0 {
		start.Add(start, bigMinusOne)
	}

	// launch some goroutines to start guessing
	for i := 0; i < routines; i++ {
		go guessFactor(c, &mu, ctx, pubkey, start)
	}

	return <-c
}

func calculatePrivateKey(pubkey *rsa.PublicKey, p *big.Int) *rsa.PrivateKey {
	modulus := pubkey.N
	q := big.NewInt(0).Div(modulus, p)
	pMinusOne := big.NewInt(0).Add(p, bigMinusOne)
	qMinusOne := big.NewInt(0).Add(q, bigMinusOne)
	eulerTotient := big.NewInt(0).Mul(pMinusOne, qMinusOne)
	e := big.NewInt(int64(pubkey.E))
	d := extendedEuclideanInverseModM(e, eulerTotient)
	return &rsa.PrivateKey{
		PublicKey: *pubkey,
		D:         d,
		Primes:    []*big.Int{p, q},
	}
}

func extendedEuclideanInverseModM(a, m *big.Int) *big.Int {
	oldR, r := a, m
	oldS, s := bigOne, bigZero
	oldT, t := bigZero, bigOne

	for {
		quotient := big.NewInt(0).Div(oldR, r)
		oldR, r = r, big.NewInt(0).Mod(oldR, r)
		oldS, s = s, big.NewInt(0).Sub(oldS, big.NewInt(0).Mul(quotient, s))
		oldT, t = t, big.NewInt(0).Sub(oldT, big.NewInt(0).Mul(quotient, t))
		if r.Cmp(bigZero) == 0 {
			break
		}
	}

	if oldS.Cmp(bigZero) == -1 {
		return oldS.Add(oldS, m)
	}
	return oldS
}
