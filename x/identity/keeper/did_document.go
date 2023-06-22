package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/sonr/x/identity/types"
)

// SetDIDDocument set a specific dIDDocument in the store from its index
func (k Keeper) SetDIDDocument(ctx sdk.Context, dIDDocument types.DIDDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))
	b := k.cdc.MustMarshal(&dIDDocument)
	store.Set(types.DIDDocumentKey(
		dIDDocument.Index,
	), b)
}

// GetDIDDocument returns a dIDDocument from its index
func (k Keeper) GetDIDDocument(
	ctx sdk.Context,
	index string,

) (val types.DIDDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))

	b := store.Get(types.DIDDocumentKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDIDDocument removes a dIDDocument from the store
func (k Keeper) RemoveDIDDocument(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))
	store.Delete(types.DIDDocumentKey(
		index,
	))
}

// GetAllDIDDocument returns all dIDDocument
func (k Keeper) GetAllDIDDocument(ctx sdk.Context) (list []types.DIDDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DIDDocumentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DIDDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
