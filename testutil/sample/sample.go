package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func GenUniqueAddresses(n int) []string {
	addresses := make([]string, n)
	hashMap := make(map[string]bool)

	for i := 0; i < n; i++ {
		addr := AccAddress()
		for hashMap[addr] == true {
			addr = AccAddress()
		}
		hashMap[addr] = true
		addresses[i] = addr
	}

	return addresses
}
