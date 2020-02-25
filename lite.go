package cosmoslite

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/rpc/client"
)

type CosmosLite struct {
	cdc *codec.Codec
	RPC *client.HTTP
}

func NewCosmosLite(cdc *codec.Codec, remote string) *CosmosLite {
	return &CosmosLite{
		cdc: cdc,
		RPC: client.NewHTTP(remote, "/websocket"),
	}
}

func (c *CosmosLite) query(path string, data common.HexBytes, result interface{}) error {
	res, err := c.RPC.ABCIQuery(path, data)
	if err != nil || res.Response.Value == nil {
		return err
	}

	return c.cdc.UnmarshalJSON(res.Response.Value, result)
}

func (c *CosmosLite) querySubspace(store string, data common.HexBytes) (result common.KVPairs, err error) {
	res, err := c.RPC.ABCIQuery(fmt.Sprintf("/store/%s/%s", store, "subspace"), data)
	if err != nil || res.Response.Value == nil {
		return result, err
	}

	err = c.cdc.UnmarshalBinaryLengthPrefixed(res.Response.Value, &result)
	return result, err
}
