package cosmoslite

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

func (c *CosmosLite) ParamDeposit() (result gov.DepositParams, err error) {
	err = c.query(fmt.Sprintf("custom/%s/params/%s", gov.QuerierRoute, gov.ParamDeposit), nil, &result)
	return result, err
}

func (c *CosmosLite) ParamVoting() (result gov.VotingParams, err error) {
	err = c.query(fmt.Sprintf("custom/%s/params/%s", gov.QuerierRoute, gov.ParamVoting), nil, &result)
	return result, err
}

func (c *CosmosLite) ParamTallying() (result gov.TallyParams, err error) {
	err = c.query(fmt.Sprintf("custom/%s/params/%s", gov.QuerierRoute, gov.ParamTallying), nil, &result)
	return result, err
}

func (c *CosmosLite) Proposal(proposalID uint64) (result gov.Proposal, err error) {
	bytes, err := c.cdc.MarshalJSON(gov.NewQueryProposalParams(proposalID))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposal), bytes, &result)
	return result, err
}

func (c *CosmosLite) Proposals(status gov.ProposalStatus, limit uint64, voter, depositor sdk.AccAddress) (result gov.Proposals, err error) {
	bytes, err := c.cdc.MarshalJSON(gov.NewQueryProposalsParams(status, limit, voter, depositor))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryProposals), bytes, &result)
	return result, err
}

func (c *CosmosLite) Deposit(proposalID uint64, depositor sdk.AccAddress) (result gov.Deposit, err error) {
	bytes, err := c.cdc.MarshalJSON(gov.NewQueryDepositParams(proposalID, depositor))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryDeposit), bytes, &result)
	return result, err
}

func (c *CosmosLite) Deposits(proposalID uint64) (result gov.Deposits, err error) {
	bytes, err := c.cdc.MarshalJSON(gov.NewQueryProposalParams(proposalID))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryDeposits), bytes, &result)
	return result, err
}

func (c *CosmosLite) Vote(proposalID uint64, voter sdk.AccAddress) (result gov.Vote, err error) {
	bytes, err := c.cdc.MarshalJSON(gov.NewQueryVoteParams(proposalID, voter))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVote), bytes, &result)
	return result, err
}

func (c *CosmosLite) Votes(proposalID uint64) (result gov.Votes, err error) {
	bytes, err := c.cdc.MarshalJSON(gov.NewQueryProposalParams(proposalID))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryVotes), bytes, &result)
	return result, err
}

func (c *CosmosLite) Tally(proposalID uint64) (result gov.TallyResult, err error) {
	bytes, err := c.cdc.MarshalJSON(gov.NewQueryProposalParams(proposalID))
	if err != nil {
		return result, err
	}

	err = c.query(fmt.Sprintf("custom/%s/%s", gov.QuerierRoute, gov.QueryTally), bytes, &result)
	return result, err
}
