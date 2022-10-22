package main

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	feegranttype "github.com/cosmos/cosmos-sdk/x/feegrant"
)

func FirmaChainFeegrantMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {

	switch msg := cosmosMsg.(type) {

	case *feegranttype.MsgGrantAllowance:
		var stringArray = []string{msg.Grantee, msg.Granter}
		return stringArray, nil

	case *feegranttype.MsgRevokeAllowance:
		var stringArray = []string{msg.Grantee, msg.Granter}
		return stringArray, nil
	}

	return nil, nil
}
