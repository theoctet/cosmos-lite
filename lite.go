package cosmoslite

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/rpc/client"
)

type CosmosLite struct {
	*client.HTTP
	cdc *codec.Codec
}

func (c *CosmosLite) query(path string, data common.HexBytes, result interface{}) error {
	res, err := c.ABCIQuery(path, data)
	if err != nil {
		return err
	}
	if res.Response.Value == nil {
		return nil
	}

	return c.cdc.UnmarshalJSON(res.Response.Value, result)
}
