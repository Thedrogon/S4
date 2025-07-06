package util

import (
    "crypto/rand"
    "math/big"
)

// Prime is a large prime for modular arithmetic (e.g., 2^127 - 1)
var Prime = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 127), big.NewInt(1))

// ModMul multiplies two numbers modulo Prime
func ModMul(a, b *big.Int) *big.Int {
    product := new(big.Int).Mul(a, b)
    return new(big.Int).Mod(product, Prime)
}

// RandomInt generates a cryptographically secure random number
func RandomInt() (*big.Int, error) {
    return rand.Int(rand.Reader, Prime)
}