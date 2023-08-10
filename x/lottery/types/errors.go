package types

// DONTCOVER

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrMinBet = sdkerrors.Register(ModuleName, 1100, fmt.Sprintf("min. place bet is %d", MinBet.Amount.Uint64()))
	ErrMaxBet = sdkerrors.Register(ModuleName, 1101, "")
)
