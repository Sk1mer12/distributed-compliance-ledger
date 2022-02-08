package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgApproveUpgrade_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgApproveUpgrade
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgApproveUpgrade{
				Creator: "invalid_address",
				Name:    testconstants.Plan.Name,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgApproveUpgrade{
				Creator: sample.AccAddress(),
				Name:    testconstants.Plan.Name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
