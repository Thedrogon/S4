package shamirsalgo

import (
	"fmt"
	"math/big"

	"github.com/Thedrogon/S4/internal/util"
	"github.com/Thedrogon/S4/pkg/types"
)

// GenerateShares splits a secret into n shares with threshold k
func GenerateShares(secret *big.Int, n, k int) ([]types.Share, error) {
    if k > n || k < 1 {
        return nil, fmt.Errorf("invalid parameters")
    }
    coeffs := make([]*big.Int, k)
    coeffs[0] = new(big.Int).Set(secret) // Secret is constant term
    for i := 1; i < k; i++ {
        coeffs[i], _ = util.RandomInt() // Random coefficients
    }
    shares := make([]types.Share, n)
    for i := 0; i < n; i++ {
        x := big.NewInt(int64(i + 1))
        y := evaluatePolynomial(coeffs, x)
        shares[i] = types.Share{X: x, Y: y}
    }
    return shares, nil
}

// evaluatePolynomial computes f(x) for a polynomial
func evaluatePolynomial(coeffs []*big.Int, x *big.Int) *big.Int {
    result := new(big.Int)
	//TODO : COMPUTATION PART
    // Compute f(x) = a_0 + a_1*x + a_2*x^2 + ...
    return result
}