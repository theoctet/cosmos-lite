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

func (c *CosmosLite) QueryValidatorDelegations(validatorAddr sdk.ValAddress) (result staking.Delegations, err error) {
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

func (c *CosmosLite) QueryDelegation(delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress) (result staking.Delegation, err error) {
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

func (c *CosmosLite) QueryDelegatorDelegations(delegatorAddr sdk.AccAddress) (result staking.Delegations, err error) {
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

func (c *CosmosLite) QueryRedelegations(delegatorAddr sdk.AccAddress) (result staking.Redelegations, err error) {
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

// nolint:dupl
func (c *CosmosLite) QueryAllValidators() (result staking.Validators, err error) {
	kvs, err := c.querySubspace(staking.StoreKey, staking.ValidatorsKey)
	if err != nil {
		return result, err
	}

	result = make(staking.Validators, 0, kvs.Len())
	for i := 0; i < kvs.Len(); i++ {
		var validator staking.Validator
		if err := c.cdc.UnmarshalBinaryLengthPrefixed(kvs[i].Value, &validator); err != nil {
			return nil, err
		}

		result = append(result, validator)
	}

	return result, nil
}

// nolint:dupl
func (c *CosmosLite) QueryAllDelegations() (result staking.Delegations, err error) {
	kvs, err := c.querySubspace(staking.StoreKey, staking.DelegationKey)
	if err != nil {
		return result, err
	}

	result = make(staking.Delegations, 0, kvs.Len())
	for i := 0; i < kvs.Len(); i++ {
		var delegation staking.Delegation
		if err := c.cdc.UnmarshalBinaryLengthPrefixed(kvs[i].Value, &delegation); err != nil {
			return nil, err
		}

		result = append(result, delegation)
	}

	return result, nil
}
