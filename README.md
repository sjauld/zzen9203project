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
