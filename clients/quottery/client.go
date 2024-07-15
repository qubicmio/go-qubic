package quottery

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/quottery/hrtypes"
	"github.com/qubic/go-qubic/clients/quottery/nodetypes"
	"github.com/qubic/go-qubic/internal/connector"
)

type Client struct {
	connector *connector.Connector
}

func NewClient(connector *connector.Connector) *Client {
	return &Client{
		connector: connector,
	}
}

func (c *Client) GetBetInfo(ctx context.Context, betID uint32) (hrtypes.BetInfo, error) {
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
		return hrtypes.BetInfo{}, errors.Wrap(err, "handling smart contract request")
	}

	var bi hrtypes.BetInfo
	err = bi.FromNodeType(result)
	if err != nil {
		return hrtypes.BetInfo{}, errors.Wrap(err, "converting from node type")
	}

	return bi, nil
}

func (c *Client) GetActiveBets(ctx context.Context) (nodetypes.ActiveBets, error) {
	rcf := connector.RequestContractFunction{
		ContractIndex: nodetypes.QuotteryContractID,
		InputType:     nodetypes.ViewID.ActiveBet,
		InputSize:     0,
	}

	var result nodetypes.ActiveBets
	err := c.connector.PerformSmartContractRequest(ctx, rcf, nil, &result)
	if err != nil {
		return nodetypes.ActiveBets{}, errors.Wrap(err, "handling smart contract request")
	}

	return result, nil
}

var QuotteryGetBetInfoRequest = connector.RequestContractFunction{
	ContractIndex: nodetypes.QuotteryContractID,
	InputType:     nodetypes.ViewID.BetInfo,
	InputSize:     4, // sizeof betID; uint32
}
