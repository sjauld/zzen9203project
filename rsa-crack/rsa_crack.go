package rsacrack

import (
	"crypto/rsa"
	"math/big"
)

var bigOne = big.NewInt(1)
var bigZero = big.NewInt(0)
var bigMinusOne = big.NewInt(-1)
var bigMinusTwo = big.NewInt(-2)

func CrackPublicKey(pubkey *rsa.PublicKey) *rsa.PrivateKey {
	modulus := pubkey.N
	// log.Printf("[INFO] factorising %+v", modulus)

	start := big.NewInt(0)
	start.Sqrt(modulus)

	// Make it odd
	if start.Bit(0) == 0 {
		start.Add(start, bigMinusOne)
	}
	for i := start; i.Cmp(bigOne) == 1; i.Add(i, bigMinusTwo) {
		guess := *i
		// if checkForObviousComposites(&guess) {
		// 	continue
		// }

		if big.NewInt(0).Mod(modulus, &guess).Cmp(bigZero) == 0 {
			// log.Printf("[INFO] found factor %v", i)
			return calculatePrivateKey(pubkey, i)
		}
	}

	return nil
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

func checkForObviousComposites(num *big.Int) bool {
	return checkDivisibleBy3Or5(num)
}

const ascii0 = 48
const ascii5 = 53

func checkDivisibleBy3Or5(num *big.Int) bool {
	n := []rune(num.String())
	l := len(n)
	if n[l-1] == ascii5 {
		return true
	}
	// var digitSum int
	// for i := 0; i < l; i++ {
	// 	digitSum += int(n[i]) - ascii0
	// }

	return false
}
