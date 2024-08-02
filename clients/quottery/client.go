package quottery

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/quottery/nodetypes"
	"github.com/qubic/go-qubic/internal/connector"
	qubicpb "github.com/qubic/go-qubic/proto/v1"
)

type Client struct {
	connector *connector.Connector
}

func NewClient(connector *connector.Connector) *Client {
	return &Client{
		connector: connector,
	}
}

func (c *Client) GetBetInfo(ctx context.Context, betID uint32) (*qubicpb.BetInfo, error) {
	rcf := connector.RequestContractFunction{
		ContractIndex: nodetypes.QuotteryContractID,
		InputType:     nodetypes.ViewID.BetInfo,
		InputSize:     4, // sizeof betID; uint32
	}

	request := struct {
		BetID uint32
	}{
		BetID: betID,
	}

	var result nodetypes.BetInfo
	err := c.connector.PerformSmartContractRequest(ctx, rcf, request, &result)
	if err != nil {
		return nil, errors.Wrap(err, "handling smart contract request")
	}

	bi, err := BetInfoConverter.ToProto(result)
	if err != nil {
		return nil, errors.Wrap(err, "converting from node type")
	}

	return bi, nil
}

func (c *Client) GetActiveBets(ctx context.Context) (*qubicpb.ActiveBets, error) {
	rcf := connector.RequestContractFunction{
		ContractIndex: nodetypes.QuotteryContractID,
		InputType:     nodetypes.ViewID.ActiveBet,
		InputSize:     8,
	}

	var result nodetypes.ActiveBets
	err := c.connector.PerformSmartContractRequest(ctx, rcf, nil, &result)
	if err != nil {
		return nil, errors.Wrap(err, "handling smart contract request")
	}

	ab := ActiveBetsConverter.ToProto(result)

	return ab, nil
}

func (c *Client) GetBasicInfo(ctx context.Context) (*qubicpb.BasicInfo, error) {
	rcf := connector.RequestContractFunction{
		ContractIndex: nodetypes.QuotteryContractID,
		InputType:     nodetypes.ViewID.BasicInfo,
		InputSize:     0,
	}

	var result nodetypes.BasicInfo
	err := c.connector.PerformSmartContractRequest(ctx, rcf, nil, &result)
	if err != nil {
		return nil, errors.Wrap(err, "handling smart contract request")
	}

	bi, err := BasicInfoConverter.ToProto(result)
	if err != nil {
		return nil, errors.Wrap(err, "converting from node type")
	}

	return bi, nil
}

func (c *Client) GetBettorsByBetOption(ctx context.Context, betID, betOption uint32) (*qubicpb.BetOptionBettors, error) {
	rcf := connector.RequestContractFunction{
		ContractIndex: nodetypes.QuotteryContractID,
		InputType:     nodetypes.ViewID.BetDetail,
		InputSize:     8, // sizeof betID + betOptions; 2 * uint32
	}

	request := struct {
		BetID     uint32
		BetOption uint32
	}{
		BetID:     betID,
		BetOption: betOption,
	}

	var result nodetypes.BetOptionDetail
	err := c.connector.PerformSmartContractRequest(ctx, rcf, request, &result)
	if err != nil {
		return nil, errors.Wrap(err, "handling smart contract request")
	}

	bob, err := BetOptionBettorsConverter.ToProto(result)
	if err != nil {
		return nil, errors.Wrap(err, "converting from node type")
	}

	return bob, nil
}
