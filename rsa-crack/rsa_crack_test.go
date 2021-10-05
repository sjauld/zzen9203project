package rsacrack

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"math/big"
	"testing"
)

func TestCrackKnownSmallPublicKey(t *testing.T) {
	testKey := rsa.PublicKey{
		N: big.NewInt(17741),
		E: 1087,
	}

	ans := CrackPublicKey(&testKey, 1)
	if e := big.NewInt(13759); e.Cmp(ans.D) != 0 {
		t.Errorf("Expected %v, got %v", e, ans.D)
	}
}

func TestCrackRandomLargePublicKey(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 65)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("[INFO] cracking %+v", privKey)
	ans := CrackPublicKey(&privKey.PublicKey, 1)
	log.Printf("[INFO] ans %+v", ans)
	if e := privKey.D; e.Cmp(ans.D) != 0 {
		t.Errorf("Expected %v, got %v", e, ans.D)
	}
}

func TestCrackExamKey(t *testing.T) {
	pubkey := &rsa.PublicKey{
		N: big.NewInt(33),
		E: 3,
	}

	ans := CrackPublicKey(pubkey, 1)
	log.Printf("[INFO] ans %+v", ans)
}

var result *rsa.PrivateKey

func benchmarkCrackPublicKey(bits, routines int, b *testing.B) {
	var res *rsa.PrivateKey
	for i := 0; i < b.N; i++ {
		privKey, err := rsa.GenerateKey(rand.Reader, bits)
		if err != nil {
			b.Fatal(err)
		}

		res = CrackPublicKey(&privKey.PublicKey, routines)
		// log.Printf("[INFO] PubKey: %+v, PrivKey: %+v", pubKey, privKey)
	}

	result = res
}

// func BenchmarkCrackPublicKey_20bits(b *testing.B) { benchmarkCrackPublicKey(20, 2, b) }
// func BenchmarkCrackPublicKey_30bits(b *testing.B) { benchmarkCrackPublicKey(30, 2, b) }
// func BenchmarkCrackPublicKey_40bits(b *testing.B) { benchmarkCrackPublicKey(40, b) }
// func BenchmarkCrackPublicKey_50bits(b *testing.B) { benchmarkCrackPublicKey(50, b) }
func BenchmarkCrackPublicKey_60bits_1routine(b *testing.B)   { benchmarkCrackPublicKey(60, 1, b) }
func BenchmarkCrackPublicKey_60bits_2routine(b *testing.B)   { benchmarkCrackPublicKey(60, 2, b) }
func BenchmarkCrackPublicKey_60bits_5routine(b *testing.B)   { benchmarkCrackPublicKey(60, 5, b) }
func BenchmarkCrackPublicKey_60bits_10routine(b *testing.B)  { benchmarkCrackPublicKey(60, 10, b) }
func BenchmarkCrackPublicKey_60bits_20routine(b *testing.B)  { benchmarkCrackPublicKey(60, 20, b) }
func BenchmarkCrackPublicKey_60bits_50routine(b *testing.B)  { benchmarkCrackPublicKey(60, 50, b) }
func BenchmarkCrackPublicKey_60bits_100routine(b *testing.B) { benchmarkCrackPublicKey(60, 100, b) }

// func BenchmarkCrackPublicKey_70bits(b *testing.B) { benchmarkCrackPublicKey(70, b) }

// func BenchmarkCrackPublicKey_80bits_1routine(b *testing.B)   { benchmarkCrackPublicKey(80, 1, b) }
// func BenchmarkCrackPublicKey_80bits_2routine(b *testing.B)   { benchmarkCrackPublicKey(80, 2, b) }
// func BenchmarkCrackPublicKey_80bits_5routine(b *testing.B)   { benchmarkCrackPublicKey(80, 5, b) }
// func BenchmarkCrackPublicKey_80bits_10routine(b *testing.B)  { benchmarkCrackPublicKey(80, 10, b) }
// func BenchmarkCrackPublicKey_80bits_20routine(b *testing.B)  { benchmarkCrackPublicKey(80, 20, b) }
// func BenchmarkCrackPublicKey_80bits_50routine(b *testing.B)  { benchmarkCrackPublicKey(80, 50, b) }
// func BenchmarkCrackPublicKey_80bits_100routine(b *testing.B) { benchmarkCrackPublicKey(80, 100, b) }
