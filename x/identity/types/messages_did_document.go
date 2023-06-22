package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDIDDocument = "create_did_document"
	TypeMsgUpdateDIDDocument = "update_did_document"
	TypeMsgDeleteDIDDocument = "delete_did_document"
)

var _ sdk.Msg = &MsgCreateDIDDocument{}

func NewMsgCreateDIDDocument(
	creator string,
	index string,

) *MsgCreateDIDDocument {
	return &MsgCreateDIDDocument{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateDIDDocument) Route() string {
	return RouterKey
}

func (msg *MsgCreateDIDDocument) Type() string {
	return TypeMsgCreateDIDDocument
}

func (msg *MsgCreateDIDDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDIDDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDIDDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDIDDocument{}

func NewMsgUpdateDIDDocument(
	creator string,
	index string,

) *MsgUpdateDIDDocument {
	return &MsgUpdateDIDDocument{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateDIDDocument) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDIDDocument) Type() string {
	return TypeMsgUpdateDIDDocument
}

func (msg *MsgUpdateDIDDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDIDDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDIDDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDIDDocument{}

func NewMsgDeleteDIDDocument(
	creator string,
	index string,

) *MsgDeleteDIDDocument {
	return &MsgDeleteDIDDocument{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteDIDDocument) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDIDDocument) Type() string {
	return TypeMsgDeleteDIDDocument
}

func (msg *MsgDeleteDIDDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDIDDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDIDDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
