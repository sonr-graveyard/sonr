package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonrhq/sonr/x/identity/types"
)

func (k msgServer) CreateDIDDocument(goCtx context.Context, msg *types.MsgCreateDIDDocument) (*types.MsgCreateDIDDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDIDDocument(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var dIDDocument = types.DIDDocument{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetDIDDocument(
		ctx,
		dIDDocument,
	)
	return &types.MsgCreateDIDDocumentResponse{}, nil
}

func (k msgServer) UpdateDIDDocument(goCtx context.Context, msg *types.MsgUpdateDIDDocument) (*types.MsgUpdateDIDDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDIDDocument(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var dIDDocument = types.DIDDocument{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetDIDDocument(ctx, dIDDocument)

	return &types.MsgUpdateDIDDocumentResponse{}, nil
}

func (k msgServer) DeleteDIDDocument(goCtx context.Context, msg *types.MsgDeleteDIDDocument) (*types.MsgDeleteDIDDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDIDDocument(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDIDDocument(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteDIDDocumentResponse{}, nil
}
