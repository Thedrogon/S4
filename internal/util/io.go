package util


import (
    "encoding/json"
    "github.com/Thedrogon/S4/pkg/types"
    "os"
)

// SaveShares saves shares to a JSON file
func SaveShares(shares []types.Share, filename string) error {
    data, err := json.Marshal(shares)
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

// LoadShares loads shares from a JSON file
func LoadShares(filename string) ([]types.Share, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    var shares []types.Share
    return shares, json.Unmarshal(data, &shares)
}