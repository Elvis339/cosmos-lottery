package types

import (
	"crypto/sha256"
	"fmt"
)

func Hash(lotteryTx []LotteryTransaction) []byte {
	str := ""
	h := sha256.New()

	for _, tx := range lotteryTx {
		str += fmt.Sprintf("%d||%d||%s||%d", tx.Id, tx.Bet.Amount.Uint64(), tx.CreatedBy, tx.LotteryId)
	}
	h.Write([]byte(str))
	bs := h.Sum(nil)

	return bs
}
