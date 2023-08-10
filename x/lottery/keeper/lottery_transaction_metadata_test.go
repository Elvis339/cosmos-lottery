package keeper

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestLotteryTransactionMetadata_Prune(t *testing.T) {
	ltMetadata := NewLotteryTransactionMetadata()

	for i := 0; i < 10; i++ {
		key := strconv.Itoa(i)
		ltMetadata.Set(key, uint64(i))

		entry, found := ltMetadata.Get(key)
		require.True(t, found)
		require.Equal(t, entry, uint64(i))
	}

	ltMetadata.Prune()

	for i := 0; i < 10; i++ {
		key := strconv.Itoa(i)
		_, exist := ltMetadata.Get(key)
		require.False(t, exist)
	}
}
