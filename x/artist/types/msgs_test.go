package types

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	addrs = []sdk.AccAddress{
		sdk.AccAddress("test1"),
		sdk.AccAddress("test2"),
	}
)

// test ValidateBasic for MsgCreateValidator
func TestMsgCreateArtist(t *testing.T) {
	tests := []struct {
		name       string
		ownerAddr  sdk.AccAddress
		expectPass bool
	}{
		{"Freddy Mercury", addrs[0], true},
		{"", addrs[0], false},
		{"U2", nil, false},
		{"Vasco Rossi", addrs[1], true},
		{"Bob Marley", sdk.AccAddress{}, false},
		{strings.Repeat("#", MaxNameLength*2), addrs[0], false},
	}

	for i, tc := range tests {
		msg := NewMsgCreateArtist(
			NewMeta(tc.name),
			tc.ownerAddr,
		)

		if tc.expectPass {
			require.NoError(t, msg.ValidateBasic(), "test: %v", i)
		} else {
			require.Error(t, msg.ValidateBasic(), "test: %v", i)
		}
	}
}
