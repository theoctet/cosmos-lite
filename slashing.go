package cosmoslite

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
)

func (c *CosmosLite) QuerySlashingParameters() (result slashing.Params, err error) {
	err = c.query(fmt.Sprintf("custom/%s/%s", slashing.QuerierRoute, slashing.QueryParameters), nil, &result)
	return result, err
}

func (c *CosmosLite) QuerySigningInfo(consAddr sdk.ConsAddress) (result slashing.ValidatorSigningInfo, err error) {
	bytes, err := c.cdc.MarshalJSON(slashing.NewQuerySigningInfoParams(consAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", slashing.QuerierRoute, slashing.QuerySigningInfo), bytes, &result)
	return result, err
}

func (c *CosmosLite) QuerySigningInfos(page, limit int) (result []slashing.ValidatorSigningInfo, err error) {
	bytes, err := c.cdc.MarshalJSON(slashing.NewQuerySigningInfosParams(page, limit))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", slashing.QuerierRoute, slashing.QuerySigningInfos), bytes, &result)
	return result, err
}
