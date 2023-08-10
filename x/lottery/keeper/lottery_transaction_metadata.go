package keeper

/*
LotteryTransactionMetadata is an in-memory data structure designed to manage lottery transactions efficiently.

- Key: Represents the user's address.
- Value: Represents the `Id` of the corresponding transaction detail.

The primary goals for this structure are:
  - Efficiently handle unique transactions: If the same user submits multiple lottery transactions, only the most recent one is considered valid. This design ensures that the counter does not increment upon repeated submissions from the same user.
  - Offer quick look-up and set operations for handling user transactions such as replacing initial transaction with the most recent one while retaining same order.

The disadvantage of this design are
  - Memory consumption, possible mitigation strategy would be to use LSM-tree data structure but for this use-case it's okay to use map
  - Synchronisation in a distributed setup, ensuring consistency across nodes can be challenging with in-memory data.
*/
type LotteryTransactionMetadata struct {
	state map[string]uint64
}

func NewLotteryTransactionMetadata() LotteryTransactionMetadata {
	return LotteryTransactionMetadata{state: make(map[string]uint64)}
}

func (m *LotteryTransactionMetadata) Get(key string) (uint64, bool) {
	index, ok := m.state[key]

	if !ok {
		return 0, false
	}

	return index, true
}

func (m *LotteryTransactionMetadata) Set(key string, value uint64) {
	m.state[key] = value
}

func (m *LotteryTransactionMetadata) GetState() map[string]uint64 {
	return m.state
}

func (m *LotteryTransactionMetadata) Prune() {
	m.state = make(map[string]uint64)
}
