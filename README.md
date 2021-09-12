## Stated goals

- Build a tool that spits out public keys of selected strengths (e.g. `e` and `N`)
- Build a tool to brute force `d` given an `e` and `N`
- Analyse the time taken to brute force `d`, based on different strengths, and
  considering future advances in processing speed taking Moore's law as an assumption
- Write a blog post that can teach others about my findings

## Public key generator

Since the purpose of this project is to test the strength of public keys,
instead of building my own public key generator I have decided to just use go's
crypto/rsa library to spit out the keys. This will make my analysis more
relevant since I will avoid inadvertantly creating weak keys through a poor
implementation of the key generator.

## Benchmarks

### Initial implementation (v0.0.1)

```
$ go test -bench=. -run Benchmark
goos: linux
goarch: amd64
pkg: github.com/sjauld/zzen9203project/rsa-crack
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkCrackPublicKey_20bits-12          19960             59321 ns/op
BenchmarkCrackPublicKey_30bits-12           8602            129309 ns/op
BenchmarkCrackPublicKey_40bits-12            588           1873675 ns/op
BenchmarkCrackPublicKey_50bits-12             37          66106342 ns/op
BenchmarkCrackPublicKey_60bits-12              1        1039506588 ns/op
BenchmarkCrackPublicKey_70bits-12              1      175639027873 ns/op
```

### Using (*big.Int).ProbablyPrime to filter out non-primes before modding

```
...
BenchmarkCrackPublicKey_20bits-12           6859            160158 ns/op
BenchmarkCrackPublicKey_30bits-12            494           2656407 ns/op
BenchmarkCrackPublicKey_40bits-12             19          86507040 ns/op
BenchmarkCrackPublicKey_50bits-12              1        2763515611 ns/op
[aborted]
```

### Checking if divisible by 3 or 5 before modding (v0.0.2)

```
...
BenchmarkCrackPublicKey_20bits-12          19557             63098 ns/op
BenchmarkCrackPublicKey_30bits-12           5175            197781 ns/op
BenchmarkCrackPublicKey_40bits-12            354           4393442 ns/op
BenchmarkCrackPublicKey_50bits-12             15         116088184 ns/op
BenchmarkCrackPublicKey_60bits-12              1        2982120056 ns/op
[aborted]
```

## Checking if divisible by 5 only before modding

```
...
BenchmarkCrackPublicKey_20bits-12          18866             62719 ns/op
BenchmarkCrackPublicKey_30bits-12           5601            217795 ns/op
BenchmarkCrackPublicKey_40bits-12            236           4861034 ns/op
BenchmarkCrackPublicKey_50bits-12             13         168084242 ns/op
BenchmarkCrackPublicKey_60bits-12              1        1626325714 ns/op
BenchmarkCrackPublicKey_70bits-12              1        28324123964 ns/op
```
