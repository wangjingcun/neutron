package types

import (
	"context"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
)

type BankKeeper interface {
	// Methods imported from bank should be defined here
	GetDenomMetaData(ctx context.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx context.Context, denomMetaData banktypes.Metadata)

	HasSupply(ctx context.Context, denom string) bool

	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error

	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
	HasBalance(ctx context.Context, addr sdk.AccAddress, amt sdk.Coin) bool
}

type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI
}

// BankHooks event hooks
type BankHooks interface {
	TrackBeforeSend(ctx context.Context, from, to sdk.AccAddress, amount sdk.Coins)       // Must be before any send is executed
	BlockBeforeSend(ctx context.Context, from, to sdk.AccAddress, amount sdk.Coins) error // Must be before any send is executed
}

type ContractKeeper interface {
	Sudo(ctx context.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
	GetContractInfo(ctx context.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}

// ChannelKeeper defines the expected IBC channel keeper
type ChannelKeeper interface {
	GetAllChannelsWithPortPrefix(ctx sdk.Context, portPrefix string) (channels []channeltypes.IdentifiedChannel)
}

// TransferKeeper defines the expected IBC Transfer keeper
type TransferKeeper interface {
	GetPort(ctx sdk.Context) string
}
