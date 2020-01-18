package cosmoslite

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func (c *CosmosLite) Account(addr sdk.AccAddress) (result auth.Account, err error) {
	bytes, err := c.cdc.MarshalJSON(auth.NewQueryAccountParams(addr))
	if err != nil {
		return nil, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount), bytes, &result)
	return result, err
}
