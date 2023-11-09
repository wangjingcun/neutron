package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/neutron-org/neutron/x/contractmanager/types"
)

// AddContractFailure adds a specific failure to the store. The provided address is used to determine
// the failure ID and they both are used to create a storage key for the failure.
//
// WARNING: The errMsg string parameter is expected to be deterministic. It means that the errMsg
// must be OS/library version agnostic and carry a concrete defined error message. One of the good
// ways to do so is to redact error as it is done in SudoLimitWrapper Sudo method:
// https://github.com/neutron-org/neutron/blob/4ee803ff75bb8fdddee1ed62fbf8b771d8006347/x/contractmanager/ibc_middleware.go#L40.
// Another good way could be passing here some constant value.
func (k Keeper) AddContractFailure(ctx sdk.Context, address string, sudoPayload []byte, errMsg string) types.Failure {
	failure := types.Failure{
		Address:     address,
		SudoPayload: sudoPayload,
		Error:       errMsg,
	}
	nextFailureID := k.GetNextFailureIDKey(ctx, failure.GetAddress())
	failure.Id = nextFailureID

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&failure)
	store.Set(types.GetFailureKey(failure.GetAddress(), nextFailureID), bz)
	return failure
}

func (k Keeper) GetNextFailureIDKey(ctx sdk.Context, address string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetFailureKeyPrefix(address))
	iterator := sdk.KVStoreReversePrefixIterator(store, []byte{})
	defer iterator.Close()

	if iterator.Valid() {
		var val types.Failure
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		return val.Id + 1
	}

	return 0
}

// GetAllFailures returns all failures
func (k Keeper) GetAllFailures(ctx sdk.Context) (list []types.Failure) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ContractFailuresKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Failure
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetFailure(ctx sdk.Context, contractAddr sdk.AccAddress, id uint64) (*types.Failure, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFailureKey(contractAddr.String(), id)

	bz := store.Get(key)
	if bz == nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "no failure found for contractAddress = %s and failureId = %d", contractAddr.String(), id)
	}
	var res types.Failure
	k.cdc.MustUnmarshal(bz, &res)

	return &res, nil
}

// ResubmitFailure tries to call sudo handler for contract with same parameters as initially.
func (k Keeper) ResubmitFailure(ctx sdk.Context, contractAddr sdk.AccAddress, failure *types.Failure) error {
	if failure.SudoPayload == nil {
		return errorsmod.Wrapf(types.ErrIncorrectFailureToResubmit, "cannot resubmit failure without sudo payload; failureId = %d", failure.Id)
	}

	if _, err := k.wasmKeeper.Sudo(ctx, contractAddr, failure.SudoPayload); err != nil {
		return errorsmod.Wrapf(types.ErrFailedToResubmitFailure, "cannot resubmit failure; failureId = %d; err = %s", failure.Id, err)
	}

	// Cleanup failure since we resubmitted it successfully
	k.removeFailure(ctx, contractAddr, failure.Id)

	return nil
}

func (k Keeper) removeFailure(ctx sdk.Context, contractAddr sdk.AccAddress, id uint64) {
	store := ctx.KVStore(k.storeKey)
	failureKey := types.GetFailureKey(contractAddr.String(), id)
	store.Delete(failureKey)
}
