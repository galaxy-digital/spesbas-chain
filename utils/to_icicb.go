package utils

import "math/big"

// ToDbas number of DBAS to Wei
func ToDbas(dbas uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(dbas), big.NewInt(1e18))
}
