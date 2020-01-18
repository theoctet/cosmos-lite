package cosmoslite

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
)

func (c *CosmosLite) QueryMintParameters() (result mint.Params, err error) {
	err = c.query(fmt.Sprintf("custom/%s/%s", mint.QuerierRoute, mint.QueryParameters), nil, &result)
	return result, err
}

func (c *CosmosLite) QueryInflation() (result sdk.Dec, err error) {
	err = c.query(fmt.Sprintf("custom/%s/%s", mint.QuerierRoute, mint.QueryInflation), nil, &result)
	return result, err
}

func (c *CosmosLite) QueryAnnualProvisions() (result sdk.Dec, err error) {
	err = c.query(fmt.Sprintf("custom/%s/%s", mint.QuerierRoute, mint.QueryAnnualProvisions), nil, &result)
	return result, err
}
