package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"cosmos-lottery/testutil/network"
	"cosmos-lottery/testutil/nullify"
	"cosmos-lottery/x/lottery/client/cli"
	"cosmos-lottery/x/lottery/types"
)

func networkWithActiveLotteryObjects(t *testing.T) (*network.Network, types.ActiveLottery) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	activeLottery := &types.ActiveLottery{}
	nullify.Fill(&activeLottery)
	state.ActiveLottery = activeLottery
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.ActiveLottery
}

func TestShowActiveLottery(t *testing.T) {
	net, obj := networkWithActiveLotteryObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.ActiveLottery
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowActiveLottery(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetActiveLotteryResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ActiveLottery)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.ActiveLottery),
				)
			}
		})
	}
}
