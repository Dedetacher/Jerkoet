package types

import (
	bstypes "github.com/babylonchain/babylon/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type BTCStakingKeeper interface {
	GetBTCValidator(ctx sdk.Context, valBTCPK []byte) (*bstypes.BTCValidator, error)
	HasBTCValidator(ctx sdk.Context, valBTCPK []byte) bool
	SlashBTCValidator(ctx sdk.Context, valBTCPK []byte) error
	GetVotingPower(ctx sdk.Context, valBTCPK []byte, height uint64) uint64
	GetVotingPowerTable(ctx sdk.Context, height uint64) map[string]uint64
	GetBTCStakingActivatedHeight(ctx sdk.Context) (uint64, error)
	RecordRewardDistCache(ctx sdk.Context)
	GetRewardDistCache(ctx sdk.Context, height uint64) (*bstypes.RewardDistCache, error)
	RemoveRewardDistCache(ctx sdk.Context, height uint64)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// IncentiveKeeper defines the expected interface needed to distribute rewards.
type IncentiveKeeper interface {
	RewardBTCStaking(ctx sdk.Context, height uint64, rdc *bstypes.RewardDistCache)
}
