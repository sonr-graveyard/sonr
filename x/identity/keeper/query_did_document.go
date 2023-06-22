package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonrhq/sonr/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DIDDocumentAll(goCtx context.Context, req *types.QueryAllDIDDocumentRequest) (*types.QueryAllDIDDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var dIDDocuments []types.DIDDocument
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	dIDDocumentStore := prefix.NewStore(store, types.KeyPrefix(types.DIDDocumentKeyPrefix))

	pageRes, err := query.Paginate(dIDDocumentStore, req.Pagination, func(key []byte, value []byte) error {
		var dIDDocument types.DIDDocument
		if err := k.cdc.Unmarshal(value, &dIDDocument); err != nil {
			return err
		}

		dIDDocuments = append(dIDDocuments, dIDDocument)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDIDDocumentResponse{DIDDocument: dIDDocuments, Pagination: pageRes}, nil
}

func (k Keeper) DIDDocument(goCtx context.Context, req *types.QueryGetDIDDocumentRequest) (*types.QueryGetDIDDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDIDDocument(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDIDDocumentResponse{DIDDocument: val}, nil
}
