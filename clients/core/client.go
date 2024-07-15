package core

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/core/hrtypes"
	"github.com/qubic/go-qubic/clients/core/nodetypes"
	"github.com/qubic/go-qubic/common"
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

func (c *Client) GetTickInfo(ctx context.Context) (hrtypes.TickInfo, error) {
	var result nodetypes.TickInfo

	err := c.connector.PerformCoreRequest(ctx, nodetypes.CurrentTickInfoTypeRequest, nil, &result)
	if err != nil {
		return hrtypes.TickInfo{}, errors.Wrap(err, "handling chainRequest")
	}

	var ti hrtypes.TickInfo
	err = ti.FromNodeType(result)
	if err != nil {
		return hrtypes.TickInfo{}, errors.Wrap(err, "converting from raw model")
	}

	return ti, nil
}

func (c *Client) GetAddressInfo(ctx context.Context, id string) (hrtypes.AddressInfo, error) {
	identity := common.Identity(id)
	pubKey, err := identity.ToPubKey(false)
	if err != nil {
		return hrtypes.AddressInfo{}, errors.Wrap(err, "converting identity to public key")
	}

	var result nodetypes.AddressInfo
	err = c.connector.PerformCoreRequest(ctx, nodetypes.BalanceTypeRequest, pubKey, &result)
	if err != nil {
		return hrtypes.AddressInfo{}, errors.Wrap(err, "handling chainRequest")
	}

	var ai hrtypes.AddressInfo
	err = ai.FromNodeType(result)
	if err != nil {
		return hrtypes.AddressInfo{}, errors.Wrap(err, "converting from raw model")
	}

	return ai, nil
}

func (c *Client) GetComputors(ctx context.Context) (hrtypes.ComputorList, error) {
	var result nodetypes.Computors

	err := c.connector.PerformCoreRequest(ctx, nodetypes.ComputorsTypeRequest, nil, &result)
	if err != nil {
		return hrtypes.ComputorList{}, errors.Wrap(err, "handling chainRequest")
	}

	var cl hrtypes.ComputorList
	err = cl.FromNodeType(result)
	if err != nil {
		return hrtypes.ComputorList{}, errors.Wrap(err, "converting from raw model")
	}

	return cl, nil
}

func (c *Client) GetTickQuorumVotes(ctx context.Context, tickNumber uint32) (hrtypes.TickQuorumVotes, error) {
	tickInfo, err := c.GetTickInfo(ctx)
	if err != nil {
		return hrtypes.TickQuorumVotes{}, errors.Wrap(err, "getting tick info")
	}

	if tickInfo.Tick < tickNumber {
		return hrtypes.TickQuorumVotes{}, errors.Errorf("Requested tick %d is in the future. Latest tick is: %d", tickNumber, tickInfo.Tick)
	}

	request := struct {
		Tick      uint32
		VoteFlags [(nodetypes.NumberOfComputors + 7) / 8]byte
	}{Tick: tickNumber}

	var result nodetypes.QuorumVotes

	err = c.connector.PerformCoreRequest(ctx, nodetypes.QuorumTickTypeRequest, request, &result)
	if err != nil {
		return hrtypes.TickQuorumVotes{}, errors.Wrap(err, "handling chainRequest")
	}

	var tqv hrtypes.TickQuorumVotes
	err = tqv.FromNodeType(result)
	if err != nil {
		return hrtypes.TickQuorumVotes{}, errors.Wrap(err, "converting from raw model")
	}

	return tqv, nil
}

func (c *Client) GetTickData(ctx context.Context, tickNumber uint32) (hrtypes.TickData, error) {
	tickInfo, err := c.GetTickInfo(ctx)
	if err != nil {
		return hrtypes.TickData{}, errors.Wrap(err, "getting tick info")
	}

	if tickInfo.Tick < tickNumber {
		return hrtypes.TickData{}, errors.Errorf("Requested tick %d is in the future. Latest tick is: %d", tickNumber, tickInfo.Tick)
	}

	request := struct{ Tick uint32 }{Tick: tickNumber}

	var result nodetypes.TickData
	err = c.connector.PerformCoreRequest(ctx, nodetypes.TickDataTypeRequest, request, &result)
	if err != nil {
		return hrtypes.TickData{}, errors.Wrap(err, "handling chainRequest")
	}

	var td hrtypes.TickData
	err = td.FromNodeType(result)
	if err != nil {
		return hrtypes.TickData{}, errors.Wrap(err, "converting from raw model")
	}

	return td, nil
}

func (c *Client) GetTickTransactions(ctx context.Context, tickNumber uint32) (hrtypes.Transactions, error) {
	tickData, err := c.GetTickData(ctx, tickNumber)
	if err != nil {
		return hrtypes.Transactions{}, errors.Wrap(err, "getting tick data")
	}

	nrTx := len(tickData.TransactionIDs)
	if nrTx == 0 {
		return hrtypes.Transactions{}, nil
	}

	requestTickTransactions := struct {
		Tick             uint32
		TransactionFlags [nodetypes.NumberOfTransactionsPerTick / 8]uint8
	}{Tick: tickNumber}

	for i := 0; i < (nrTx+7)/8; i++ {
		requestTickTransactions.TransactionFlags[i] = 0
	}
	for i := (nrTx + 7) / 8; i < nodetypes.NumberOfTransactionsPerTick/8; i++ {
		requestTickTransactions.TransactionFlags[i] = 1
	}

	var result nodetypes.Transactions
	err = c.connector.PerformCoreRequest(ctx, nodetypes.TickTransactionsTypeRequest, requestTickTransactions, &result)
	if err != nil {
		return nil, errors.Wrap(err, "handling chainRequest")
	}

	txs := make(hrtypes.Transactions, len(result))
	err = txs.FromNodeType(result)
	if err != nil {
		return txs, errors.Wrap(err, "converting from raw model")
	}

	return txs, nil
}

func (c *Client) GetTickTransactionsStatus(ctx context.Context, tickNumber uint32) (hrtypes.TransactionsStatus, error) {
	tickInfo, err := c.GetTickInfo(ctx)
	if err != nil {
		return hrtypes.TransactionsStatus{}, errors.Wrap(err, "getting tick info")
	}

	if tickInfo.Tick < tickNumber {
		return hrtypes.TransactionsStatus{}, errors.Errorf("Requested tick %d is in the future. Latest tick is: %d", tickNumber, tickInfo.Tick)
	}

	request := struct {
		Tick uint32
	}{
		Tick: tickNumber,
	}

	var result nodetypes.TransactionStatus
	err = c.connector.PerformCoreRequest(ctx, nodetypes.TxStatusTypeRequest, request, &result)
	if err != nil {
		return hrtypes.TransactionsStatus{}, errors.Wrap(err, "handling chainRequest")
	}

	var txStatus hrtypes.TransactionsStatus

	err = txStatus.FromNodeType(result)
	if err != nil {
		return hrtypes.TransactionsStatus{}, errors.Wrap(err, "converting from raw model")
	}

	return txStatus, nil
}
