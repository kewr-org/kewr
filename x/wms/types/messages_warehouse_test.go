package types

import (
	"testing"

	"kewr/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateWarehouse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateWarehouse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateWarehouse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateWarehouse{
				Creator: sample.AccAddress(),
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

func TestMsgUpdateWarehouse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateWarehouse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateWarehouse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateWarehouse{
				Creator: sample.AccAddress(),
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

func TestMsgDeleteWarehouse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteWarehouse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteWarehouse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteWarehouse{
				Creator: sample.AccAddress(),
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
