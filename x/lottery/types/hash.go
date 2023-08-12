package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

func Hash(lotteryTx []LotteryTransaction) []byte {
	h := sha256.New()
	buf := new(bytes.Buffer)

	for _, tx := range lotteryTx {
		binary.Write(buf, binary.BigEndian, tx.Id)
		binary.Write(buf, binary.BigEndian, tx.Bet.Amount.Uint64())
		buf.WriteString(tx.CreatedBy)
		binary.Write(buf, binary.BigEndian, tx.LotteryId)

		h.Write(buf.Bytes())
		buf.Reset() // Reset the buffer for the next iteration
	}

	return h.Sum(nil)
}
