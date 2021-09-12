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

	ans := CrackPublicKey(&testKey)
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
	ans := CrackPublicKey(&privKey.PublicKey)
	log.Printf("[INFO] ans %+v", ans)
	if e := privKey.D; e.Cmp(ans.D) != 0 {
		t.Errorf("Expected %v, got %v", e, ans.D)
	}
}

var result *rsa.PrivateKey

func benchmarkCrackPublicKey(bits int, b *testing.B) {
	var res *rsa.PrivateKey
	for i := 0; i < b.N; i++ {
		privKey, err := rsa.GenerateKey(rand.Reader, bits)
		if err != nil {
			b.Fatal(err)
		}

		res = CrackPublicKey(&privKey.PublicKey)
		// log.Printf("[INFO] PubKey: %+v, PrivKey: %+v", pubKey, privKey)
	}

	result = res
}

func BenchmarkCrackPublicKey_20bits(b *testing.B) { benchmarkCrackPublicKey(20, b) }
func BenchmarkCrackPublicKey_30bits(b *testing.B) { benchmarkCrackPublicKey(30, b) }
func BenchmarkCrackPublicKey_40bits(b *testing.B) { benchmarkCrackPublicKey(40, b) }
func BenchmarkCrackPublicKey_50bits(b *testing.B) { benchmarkCrackPublicKey(50, b) }
func BenchmarkCrackPublicKey_60bits(b *testing.B) { benchmarkCrackPublicKey(60, b) }
func BenchmarkCrackPublicKey_70bits(b *testing.B) { benchmarkCrackPublicKey(70, b) }
