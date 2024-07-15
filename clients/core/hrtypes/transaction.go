package hrtypes

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/core/nodetypes"
	"github.com/qubic/go-qubic/common"
)

type Transaction struct {
	ID           common.Identity
	SourceID     common.Identity
	DestID       common.Identity
	Amount       int64
	Tick         uint32
	InputType    uint16
	InputSize    uint16
	InputHex     string
	SignatureHex string
	Digest       [32]byte
}

func (tx *Transaction) FromNodeType(m nodetypes.Transaction) error {
	txc := txConverter{rawTx: m}
	converted, err := txc.toType()
	if err != nil {
		return errors.Wrap(err, "converting raw tx to type")
	}

	*tx = converted

	return nil
}

type txConverter struct {
	rawTx nodetypes.Transaction
}

func (c *txConverter) toType() (Transaction, error) {
	digest, err := c.rawTx.Digest()
	if err != nil {
		return Transaction{}, errors.Wrap(err, "getting tx digest")
	}

	id, err := common.GetIDFrom32Bytes(digest, true)
	if err != nil {
		return Transaction{}, errors.Wrap(err, "getting tx id")
	}

	sourceID, err := common.GetIDFrom32Bytes(c.rawTx.SourcePublicKey, false)
	if err != nil {
		return Transaction{}, errors.Wrap(err, "getting tx source id")
	}

	destID, err := common.GetIDFrom32Bytes(c.rawTx.DestinationPublicKey, false)
	if err != nil {
		return Transaction{}, errors.Wrap(err, "getting tx dest id")
	}

	return Transaction{
		ID:           id,
		SourceID:     sourceID,
		DestID:       destID,
		Amount:       c.rawTx.Amount,
		Tick:         c.rawTx.Tick,
		InputType:    c.rawTx.InputType,
		InputSize:    c.rawTx.InputSize,
		InputHex:     hex.EncodeToString(c.rawTx.Input),
		SignatureHex: hex.EncodeToString(c.rawTx.Signature[:]),
		Digest:       digest,
	}, nil
}

type Transactions []Transaction

func (txs *Transactions) FromNodeType(m nodetypes.Transactions) error {
	convertedTxs := make([]Transaction, len(m))
	for i, mTx := range m {
		var tx Transaction
		err := tx.FromNodeType(mTx)
		if err != nil {
			return errors.Wrapf(err, "calling FromNodeType tx index: %d", i)
		}
		convertedTxs[i] = tx
	}
	*txs = convertedTxs

	return nil
}

type TransactionsStatus struct {
	CurrentTickOfNode      uint32
	Tick                   uint32
	TxCount                uint32
	StatusPerTransactionID map[common.Identity]bool
}

func (ts *TransactionsStatus) FromNodeType(m nodetypes.TransactionStatus) error {
	tsc := transactionsStatusConverter{rawTxStatus: m}
	converted, err := tsc.toType()
	if err != nil {
		return errors.Wrap(err, "converting raw tx status to type")
	}

	*ts = converted

	return nil
}

type transactionsStatusConverter struct {
	rawTxStatus nodetypes.TransactionStatus
}

func (c *transactionsStatusConverter) toType() (TransactionsStatus, error) {
	statuses := make(map[common.Identity]bool)

	for index, digest := range c.rawTxStatus.TransactionDigests {
		id, err := common.GetIDFrom32Bytes(digest, true)
		if err != nil {
			return TransactionsStatus{}, errors.Wrapf(err, "getting tx id for tx with digest hex: %s", hex.EncodeToString(digest[:]))
		}

		moneyFlew := c.getMoneyFlewFromBits(index)
		statuses[id] = moneyFlew
	}

	return TransactionsStatus{
		CurrentTickOfNode:      c.rawTxStatus.CurrentTickOfNode,
		Tick:                   c.rawTxStatus.Tick,
		TxCount:                c.rawTxStatus.TxCount,
		StatusPerTransactionID: statuses,
	}, nil
}

func (c *transactionsStatusConverter) getMoneyFlewFromBits(digestIndex int) bool {
	pos := digestIndex / 8
	bitIndex := digestIndex % 8

	return c.getNthBit(pos, bitIndex)
}

func (c *transactionsStatusConverter) getNthBit(inputPos, bitIndex int) bool {
	input := c.rawTxStatus.MoneyFlew[inputPos]
	// Shift the input byte to the right by the bitIndex positions
	// This isolates the bit at the bitIndex position at the least significant bit position
	shifted := input >> bitIndex

	// Extract the least significant bit using a bitwise AND operation with 1
	// If the least significant bit is 1, the result will be 1; otherwise, it will be 0
	return shifted&1 == 1
}

func NewSimpleTransferTransaction(sourceID, destinationID string, amount int64, targetTick uint32) (nodetypes.Transaction, error) {
	srcID := common.Identity(sourceID)
	destID := common.Identity(destinationID)
	srcPubKey, err := srcID.ToPubKey(false)
	if err != nil {
		return nodetypes.Transaction{}, errors.Wrap(err, "converting src id string to pubkey")
	}
	destPubKey, err := destID.ToPubKey(false)
	if err != nil {
		return nodetypes.Transaction{}, errors.Wrap(err, "converting dest id string to pubkey")
	}

	return nodetypes.Transaction{
		SourcePublicKey:      srcPubKey,
		DestinationPublicKey: destPubKey,
		Amount:               amount,
		Tick:                 targetTick,
		InputType:            0,
		InputSize:            0,
		Input:                nil,
	}, nil
}
