package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPlaceBet = "place_bet"

var _ sdk.Msg = &MsgPlaceBet{}

func NewMsgPlaceBet(creator string, bet uint64) *MsgPlaceBet {
	return &MsgPlaceBet{
		Creator: creator,
		Bet:     bet,
	}
}

func (msg *MsgPlaceBet) Route() string {
	return RouterKey
}

func (msg *MsgPlaceBet) Type() string {
	return TypeMsgPlaceBet
}

func (msg *MsgPlaceBet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPlaceBet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPlaceBet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.GetBet() < 1 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "bet cannot be negative or zero")
	}

	bet := sdk.NewInt64Coin(TokenDenom, int64(msg.GetBet()))

	if bet.IsLT(MinBet) {
		return ErrMinBet.Wrapf(fmt.Sprintf("min. place bet is %d, you sent %d", MinBet.Amount.Uint64(), bet.Amount.Uint64()))
	}

	if !bet.IsLTE(MaxBet) {
		return ErrMaxBet.Wrapf(fmt.Sprintf("max. place bet is %d you sent %d", MaxBet.Amount.Uint64(), bet.Amount.Uint64()))
	}

	return nil
}
