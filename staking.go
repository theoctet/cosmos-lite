package cosmoslite

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func (c *CosmosLite) QueryValidators(page, limit int, status string) (result staking.Validators, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryValidatorsParams(page, limit, status))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryValidators), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryValidator(validatorAddr sdk.ValAddress) (result staking.Validator, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryValidatorParams(validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryValidator), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryValidatorDelegations(validatorAddr sdk.ValAddress) (result staking.DelegationResponses, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryValidatorParams(validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryValidatorDelegations), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryValidatorUnbondingDelegations(validatorAddr sdk.ValAddress) (result staking.UnbondingDelegations, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryValidatorParams(validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryValidatorDelegations), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryDelegation(delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress) (result staking.DelegationResponse, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryBondsParams(delegatorAddr, validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegation), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryUnbondingDelegation(delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress) (result staking.UnbondingDelegation, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryBondsParams(delegatorAddr, validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryUnbondingDelegation), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryDelegatorDelegations(delegatorAddr sdk.AccAddress) (result staking.DelegationResponses, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryDelegatorParams(delegatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorDelegations), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryDelegatorUnbondingDelegations(delegatorAddr sdk.AccAddress) (result staking.UnbondingDelegations, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryDelegatorParams(delegatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorUnbondingDelegations), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryRedelegations(delegatorAddr sdk.AccAddress) (result staking.RedelegationResponses, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryDelegatorParams(delegatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryRedelegations), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryDelegatorValidators(delegatorAddr sdk.AccAddress) (result staking.Validators, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryDelegatorParams(delegatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorValidators), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryDelegatorValidator(delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress) (result staking.Validator, err error) {
	bytes, err := c.cdc.MarshalJSON(staking.NewQueryBondsParams(delegatorAddr, validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryDelegatorValidator), bytes, &result)
	return result, err
}

func (c *CosmosLite) QueryPool() (result staking.Pool, err error) {
	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryPool), nil, &result)
	return result, err
}

func (c *CosmosLite) QueryStakingParameters() (result staking.Pool, err error) {
	err = c.query(fmt.Sprintf("custom/%s/%s", staking.QuerierRoute, staking.QueryParameters), nil, &result)
	return result, err
}
