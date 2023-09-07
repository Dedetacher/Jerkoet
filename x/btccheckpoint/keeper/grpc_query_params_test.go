package keeper_test

import (
	"math/rand"
	"testing"

	testkeeper "github.com/babylonchain/babylon/testutil/keeper"
	"github.com/babylonchain/babylon/x/btccheckpoint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// FuzzParamsQuery fuzzes queryClient.Params
// 1. Generate random param
// 2. When EpochInterval is 0, ensure `Validate` returns an error
// 3. Randomly set the param via query and check if the param has been updated
func FuzzParamsQuery(f *testing.F) {
	f.Add(uint64(11111), uint64(3232), int64(23))
	f.Add(uint64(22222), uint64(444), int64(330))
	f.Add(uint64(22222), uint64(12333), int64(101))

	f.Fuzz(func(t *testing.T, btcConfirmationDepth uint64, checkpointFinalizationTimeout uint64, seed int64) {
		r := rand.New(rand.NewSource(seed))

		// params generated by fuzzer
		params := types.DefaultParams()
		params.BtcConfirmationDepth = btcConfirmationDepth
		params.CheckpointFinalizationTimeout = checkpointFinalizationTimeout

		// test the case of BtcConfirmationDepth == 0
		// after that, change BtcConfirmationDepth to a random non-zero value
		if btcConfirmationDepth == 0 {
			// validation should not pass with zero EpochInterval
			require.Error(t, params.Validate())
			params.BtcConfirmationDepth = uint64(r.Int())
		}

		// test the case of CheckpointFinalizationTimeout == 0
		// after that, change CheckpointFinalizationTimeout to a random non-zero value
		if checkpointFinalizationTimeout == 0 {
			// validation should not pass with zero EpochInterval
			require.Error(t, params.Validate())
			params.CheckpointFinalizationTimeout = uint64(r.Int())
		}

		if btcConfirmationDepth >= checkpointFinalizationTimeout {
			// validation should not pass with BtcConfirmationDepth >= CheckpointFinalizationTimeout
			require.Error(t, params.Validate())

			// swap the values so we can continue the test
			params.CheckpointFinalizationTimeout = btcConfirmationDepth
			params.BtcConfirmationDepth = checkpointFinalizationTimeout
		}

		keeper, ctx := testkeeper.NewBTCCheckpointKeeper(t, nil, nil, nil, nil)
		wctx := sdk.WrapSDKContext(ctx)

		// if setParamsFlag == 0, set params
		setParamsFlag := r.Intn(2)
		if setParamsFlag == 0 {
			if err := keeper.SetParams(ctx, params); err != nil {
				panic(err)
			}
		}
		req := types.QueryParamsRequest{}
		resp, err := keeper.Params(wctx, &req)
		require.NoError(t, err)
		// if setParamsFlag == 0, resp.Params should be changed, otherwise default
		if setParamsFlag == 0 {
			require.Equal(t, params, resp.Params)
		} else {
			require.Equal(t, types.DefaultParams(), resp.Params)
		}
	})
}
