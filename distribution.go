package cosmoslite

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
)

func (c *CosmosLite) ParamCommunityTax() (result sdk.Dec, err error) {
	err = c.query(fmt.Sprintf("custom/%s/params/%s", distribution.QuerierRoute, distribution.ParamCommunityTax), nil, &result)
	return result, err
}

func (c *CosmosLite) ParamBaseProposerReward() (result sdk.Dec, err error) {
	err = c.query(fmt.Sprintf("custom/%s/params/%s", distribution.QuerierRoute, distribution.ParamBaseProposerReward), nil, &result)
	return result, err
}

func (c *CosmosLite) ParamBonusProposerReward() (result sdk.Dec, err error) {
	err = c.query(fmt.Sprintf("custom/%s/params/%s", distribution.QuerierRoute, distribution.ParamBonusProposerReward), nil, &result)
	return result, err
}

func (c *CosmosLite) ParamWithdrawAddrEnabled() (result bool, err error) {
	err = c.query(fmt.Sprintf("custom/%s/params/%s", distribution.QuerierRoute, distribution.ParamWithdrawAddrEnabled), nil, &result)
	return result, err
}

func (c *CosmosLite) ValidatorOutstandingRewards(validatorAddr sdk.ValAddress) (result distribution.ValidatorOutstandingRewards, err error) {
	bytes, err := c.cdc.MarshalJSON(distribution.NewQueryValidatorOutstandingRewardsParams(validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", distribution.QuerierRoute, distribution.QueryValidatorOutstandingRewards), bytes, &result)
	return result, err
}

func (c *CosmosLite) ValidatorCommission(validatorAddr sdk.ValAddress) (result distribution.ValidatorAccumulatedCommission, err error) {
	bytes, err := c.cdc.MarshalJSON(distribution.NewQueryValidatorCommissionParams(validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", distribution.QuerierRoute, distribution.QueryValidatorCommission), bytes, &result)
	return result, err
}

func (c *CosmosLite) ValidatorSlashes(validatorAddr sdk.ValAddress, startingHeight, endingHeight uint64) (result distribution.ValidatorSlashEvents, err error) {
	bytes, err := c.cdc.MarshalJSON(distribution.NewQueryValidatorSlashesParams(validatorAddr, startingHeight, endingHeight))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", distribution.QuerierRoute, distribution.QueryValidatorSlashes), bytes, &result)
	return result, err
}

func (c *CosmosLite) DelegationRewards(delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress) (result sdk.DecCoins, err error) {
	bytes, err := c.cdc.MarshalJSON(distribution.NewQueryDelegationRewardsParams(delegatorAddr, validatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", distribution.QuerierRoute, distribution.QueryDelegationRewards), bytes, &result)
	return result, err
}

func (c *CosmosLite) DelegatorTotalRewards(delegatorAddr sdk.AccAddress) (result distribution.QueryDelegatorTotalRewardsResponse, err error) {
	bytes, err := c.cdc.MarshalJSON(distribution.NewQueryDelegatorParams(delegatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", distribution.QuerierRoute, distribution.QueryDelegatorTotalRewards), bytes, &result)
	return result, err
}

func (c *CosmosLite) DelegatorValidators(delegatorAddr sdk.AccAddress) (result []sdk.ValAddress, err error) {
	bytes, err := c.cdc.MarshalJSON(distribution.NewQueryDelegatorParams(delegatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", distribution.QuerierRoute, distribution.QueryDelegatorValidators), bytes, &result)
	return result, err
}

func (c *CosmosLite) DelegatorWithdrawAddress(delegatorAddr sdk.AccAddress) (result sdk.AccAddress, err error) {
	bytes, err := c.cdc.MarshalJSON(distribution.NewQueryDelegatorWithdrawAddrParams(delegatorAddr))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", distribution.QuerierRoute, distribution.QueryWithdrawAddr), bytes, &result)
	return result, err
}

func (c *CosmosLite) CommunityPool() (result sdk.DecCoins, err error) {
	err = c.query(fmt.Sprintf("custom/%s/%s", distribution.QuerierRoute, distribution.QueryCommunityPool), nil, &result)
	return result, err
}
