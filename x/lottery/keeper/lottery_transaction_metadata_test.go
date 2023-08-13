package keeper

import (
	"cosmos-lottery/testutil/sample"
	"cosmos-lottery/x/lottery/types"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	alice  = sample.AccAddress()
	bob    = sample.AccAddress()
	peggy  = sample.AccAddress()
	victor = sample.AccAddress()
)

func TestTransactionMetadata_MinAndMaxBet(t *testing.T) {
	m := NewLotteryTransactionMetadata()

	lotteryTx := []types.LotteryTransaction{
		sample.CreateLotteryTx(alice, 4),
		sample.CreateLotteryTx(bob, 1),
		sample.CreateLotteryTx(peggy, 3),
	}

	for _, tx := range lotteryTx {
		m.Set(tx)
	}

	foundMinBet, minBet, minBetAddr := m.GetMinBet()
	require.True(t, foundMinBet)
	require.Equal(t, int64(1), minBet.Amount.Int64())
	require.Equal(t, minBetAddr, bob)

	foundMaxBet, maxBet, maxBetAddr := m.GetMaxBet()
	require.True(t, foundMaxBet)
	require.Equal(t, int64(4), maxBet.Amount.Int64())
	require.Equal(t, maxBetAddr, alice)
}

// TestTransactionMetadata_BetSum bet sum should be updated if address updates it's bet
func TestTransactionMetadata_BetSum(t *testing.T) {
	m := NewLotteryTransactionMetadata()

	lotteryTx := []types.LotteryTransaction{
		sample.CreateLotteryTx(alice, 4),
		sample.CreateLotteryTx(bob, 2),
		sample.CreateLotteryTx(peggy, 3),
	}

	for _, tx := range lotteryTx {
		m.Set(tx)
	}

	require.Equal(t, int64(9), m.GetBetSum().Amount.Int64())

	m.Set(sample.CreateLotteryTx(bob, 1))
	require.Equal(t, int64(8), m.GetBetSum().Amount.Int64())

	m.Set(sample.CreateLotteryTx(bob, 3))
	require.Equal(t, int64(10), m.GetBetSum().Amount.Int64())
}

func TestTransactionMetadata_Prune(t *testing.T) {
	m := NewLotteryTransactionMetadata()

	lotteryTx := []types.LotteryTransaction{
		sample.CreateLotteryTx(alice, 4),
		sample.CreateLotteryTx(bob, 2),
		sample.CreateLotteryTx(peggy, 3),
	}

	for _, tx := range lotteryTx {
		m.Set(tx)
	}

	m.Prune()

	minBetSet, minBet, minBetAddr := m.GetMinBet()
	require.False(t, minBetSet)
	require.Equal(t, int64(0), minBet.Amount.Int64())
	require.Equal(t, "0", minBetAddr)

	maxBetSet, maxBet, maxetAddr := m.GetMaxBet()
	require.False(t, maxBetSet)
	require.Equal(t, int64(0), maxBet.Amount.Int64())
	require.Equal(t, "0", maxetAddr)

	bet := m.GetBet(alice)
	require.Equal(t, int64(0), bet.Amount.Int64())

	sum := m.GetBetSum()
	require.Equal(t, int64(0), sum.Amount.Int64())

	txIdFound, _ := m.GetLotteryTransactionId(alice)
	require.False(t, txIdFound)
}

func TestLotteryTransactionMetadata_RemoveLotteryTransactionId(t *testing.T) {
	m := NewLotteryTransactionMetadata()

	lotteryTx := []types.LotteryTransaction{
		sample.CreateLotteryTx(alice, 4),
		sample.CreateLotteryTx(bob, 2),
		sample.CreateLotteryTx(peggy, 3),
		sample.CreateLotteryTx(victor, 9),
	}

	for _, tx := range lotteryTx {
		m.Set(tx)
	}

	require.Equal(t, uint64(18), m.GetBetSum().Amount.Uint64())

	minBetSet, minBet, minBetAddr := m.GetMinBet()
	require.True(t, minBetSet)
	require.Equal(t, int64(2), minBet.Amount.Int64())
	require.Equal(t, bob, minBetAddr)

	maxBetSet, maxBet, maxetAddr := m.GetMaxBet()
	require.True(t, maxBetSet)
	require.Equal(t, int64(9), maxBet.Amount.Int64())
	require.Equal(t, victor, maxetAddr)

	m.RemoveLotteryTransactionId(bob)
	newMinBetSet, newMinBet, newMinBetAddr := m.GetMinBet()
	require.True(t, newMinBetSet)
	require.Equal(t, int64(3), newMinBet.Amount.Int64())
	require.Equal(t, peggy, newMinBetAddr)
	require.Equal(t, uint64(16), m.GetBetSum().Amount.Uint64())

	m.RemoveLotteryTransactionId(victor)
	newMaxBetSet, newMaxBet, newMaxetAddr := m.GetMaxBet()
	require.True(t, newMaxBetSet)
	require.Equal(t, int64(4), newMaxBet.Amount.Int64())
	require.Equal(t, alice, newMaxetAddr)
	fmt.Println(m.GetBetSum().Amount.Uint64())
	require.Equal(t, uint64(7), m.GetBetSum().Amount.Uint64())
}
