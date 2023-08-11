package keeper

import (
	"cosmos-lottery/x/lottery/types"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	zero                    = sdk.NewInt64Coin("token", 0)
	defaultSerializationKey = "0||0"
)

/*
LotteryTransactionMetadata is an in-memory data structure designed to manage lottery transactions efficiently.

The primary goals for this structure are:
  - Efficiently handle unique transactions: If the same user submits multiple lottery transactions, only the most recent one is considered valid. This design ensures that the counter does not increment upon repeated submissions from the same user.
  - Offer quick look-up and set operations for handling user transactions such as replacing initial transaction with the most recent one while retaining same order.
  - O(1) access of sum of all bets, minimal bet, and maximal bet

The disadvantage of this design are
  - Memory consumption, possible mitigation strategy would be to use LSM-tree data structure but for this use-case it's okay to use map
  - Synchronisation in a distributed setup, ensuring consistency across nodes can be challenging with in-memory data.
*/
type LotteryTransactionMetadata struct {
	betSum sdk.Coin // Total sum of all the bets.
	minBet string   // Encoded details of the minimum bet.
	maxBet string   // Encoded details of the maximum bet.

	// addressToBet efficiently identifies the previous placed bet by the address.
	// It is utilized so that `betSum` reflects accurate calculations when a
	// user updates their bet.
	addressToBet map[string]sdk.Coin

	// addressToLotteryTxId maps user addresses to their lottery transaction IDs.
	// It aids in quick lookups for user transaction details.
	addressToLotteryTxId map[string]uint64
}

func NewLotteryTransactionMetadata() LotteryTransactionMetadata {
	return LotteryTransactionMetadata{
		betSum:               zero,
		minBet:               defaultSerializationKey,
		maxBet:               defaultSerializationKey,
		addressToBet:         make(map[string]sdk.Coin),
		addressToLotteryTxId: make(map[string]uint64),
	}
}

// Set updates or inserts a new lottery transaction into the metadata.
// It efficiently manages the betSum, minBet, and maxBet while accounting
// for any changes or new bets made by users.
func (m *LotteryTransactionMetadata) Set(lotteryTx types.LotteryTransaction) {
	address := lotteryTx.GetCreatedBy()
	found, _ := m.GetLotteryTransactionId(address)
	currBet := m.GetBet(address)

	if !found && currBet.IsEqual(zero) {
		m.addressToLotteryTxId[address] = lotteryTx.LotteryId
	}

	m.betSum = m.betSum.Sub(currBet)
	m.betSum = m.betSum.Add(lotteryTx.Bet)

	// Upsert address to bet
	m.addressToBet[address] = lotteryTx.Bet

	bet := lotteryTx.Bet
	isMinBetSet, minAmount, _ := m.GetMinBet()

	if !isMinBetSet == true {
		m.minBet = m.encodeBet(lotteryTx)
	} else {
		if bet.IsLT(minAmount) {
			m.minBet = m.encodeBet(lotteryTx)
		}
	}

	isMaxBetSet, maxAmount, _ := m.GetMaxBet()

	if !isMaxBetSet {
		m.maxBet = m.encodeBet(lotteryTx)
	} else {
		if bet.IsGTE(maxAmount) {
			m.maxBet = m.encodeBet(lotteryTx)
		}
	}
}

// Prune resets the LotteryTransactionMetadata to its default state.
// This is typically used after a lottery round concludes.
func (m *LotteryTransactionMetadata) Prune() {
	m.minBet = defaultSerializationKey
	m.maxBet = defaultSerializationKey
	m.betSum = zero
	m.addressToBet = make(map[string]sdk.Coin)
	m.addressToLotteryTxId = make(map[string]uint64)
}

func (m *LotteryTransactionMetadata) encodeBet(lotteryTx types.LotteryTransaction) string {
	amountStr := strconv.FormatUint(lotteryTx.Bet.Amount.Uint64(), 10)
	addr := lotteryTx.GetCreatedBy()

	return fmt.Sprintf("%s||%s", amountStr, addr)
}

func (m *LotteryTransactionMetadata) decodeBet(serialized string) (sdk.Coin, string) {
	slice := strings.Split(serialized, "||")
	parseInt, err := strconv.ParseInt(slice[0], 10, 64)
	if err != nil {
		panic(err)
	}
	amount := sdk.NewInt64Coin("token", parseInt)

	return amount, slice[1]
}

func (m *LotteryTransactionMetadata) GetBetSum() sdk.Coin {
	return m.betSum
}

func (m *LotteryTransactionMetadata) GetMinBet() (bool, sdk.Coin, string) {
	serialized := m.minBet
	amount, address := m.decodeBet(serialized)

	fmt.Println("GET_MIN_BET = amount=", amount)
	fmt.Println("GET_MIN_BET address=", address)

	if amount.IsEqual(zero) && address == "0" {
		fmt.Println("GET_MIN_BET returning = false, zero, 0", amount)
		return false, zero, "0"
	}

	return true, amount, address
}
func (m *LotteryTransactionMetadata) GetMaxBet() (bool, sdk.Coin, string) {
	serialized := m.maxBet
	amount, address := m.decodeBet(serialized)

	if amount.IsEqual(zero) && address == "0" {
		return false, zero, "0"
	}

	return true, amount, address
}

func (m *LotteryTransactionMetadata) GetBet(address string) sdk.Coin {
	bet, ok := m.addressToBet[address]
	if !ok {
		return zero
	}
	return bet
}

func (m *LotteryTransactionMetadata) GetLotteryTransactionId(address string) (bool, uint64) {
	lotteryTxId, ok := m.addressToLotteryTxId[address]
	if !ok {
		return false, 0
	}
	return true, lotteryTxId
}
