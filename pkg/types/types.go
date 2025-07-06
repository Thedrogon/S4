package types

import "math/big"

// Share represents a single (x, y) point in Shamir’s Secret Sharing
type Share struct {
    X *big.Int
    Y *big.Int
}

// Config holds parameters for splitting a secret
type Config struct {
    Secret    string // Input secret
    Shares    int    // Number of shares (n)
    Threshold int    // Threshold (k)
}