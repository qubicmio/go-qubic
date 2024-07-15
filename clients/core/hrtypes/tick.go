package hrtypes

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/core/nodetypes"
	"github.com/qubic/go-qubic/common"
	"time"
)

type TickData struct {
	ComputorIndex  uint16
	Epoch          uint16
	TickNumber     uint32
	Timestamp      int64
	VarStruct      []byte
	TimeLock       []byte
	TransactionIDs []common.Identity
	ContractFees   []int64
	SignatureHex   string
}

func (td *TickData) FromNodeType(m nodetypes.TickData) error {
	tdc := tickDataConverter{rawTickData: m}
	tickData, err := tdc.toType()
	if err != nil {
		return errors.Wrap(err, "tick data to type")
	}

	*td = tickData

	return nil
}

type tickDataConverter struct {
	rawTickData nodetypes.TickData
}

func (tdc *tickDataConverter) toType() (TickData, error) {
	if tdc.rawTickData.IsEmpty() {
		return TickData{}, nil
	}

	date := time.Date(2000+int(tdc.rawTickData.Year), time.Month(tdc.rawTickData.Month), int(tdc.rawTickData.Day), int(tdc.rawTickData.Hour), int(tdc.rawTickData.Minute), int(tdc.rawTickData.Second), 0, time.UTC)
	date.Add(time.Duration(tdc.rawTickData.Millisecond) * time.Millisecond)

	transactionIds, err := common.PubKeysToIdentities(tdc.rawTickData.TransactionDigests[:], true)
	if err != nil {
		return TickData{}, errors.Wrap(err, "getting transaction ids from digests")
	}
	return TickData{
		ComputorIndex:  tdc.rawTickData.ComputorIndex,
		Epoch:          tdc.rawTickData.Epoch,
		TickNumber:     tdc.rawTickData.Tick,
		Timestamp:      date.UnixMilli(),
		VarStruct:      tdc.rawTickData.UnionData[:],
		TimeLock:       tdc.rawTickData.Timelock[:],
		TransactionIDs: transactionIds,
		ContractFees:   contractFeesToProto(tdc.rawTickData.ContractFees),
		SignatureHex:   hex.EncodeToString(tdc.rawTickData.Signature[:]),
	}, nil
}

func contractFeesToProto(contractFees [1024]int64) []int64 {
	protoContractFees := make([]int64, 0, len(contractFees))
	for _, fee := range contractFees {
		if fee == 0 {
			continue
		}
		protoContractFees = append(protoContractFees, fee)
	}
	return protoContractFees
}

type TickInfo struct {
	TickDuration            uint16
	Epoch                   uint16
	Tick                    uint32
	NumberOfAlignedVotes    uint16
	NumberOfMisalignedVotes uint16
	InitialTick             uint32
}

func (ti *TickInfo) FromNodeType(m nodetypes.TickInfo) error {
	*ti = TickInfo{
		TickDuration:            m.TickDuration,
		Epoch:                   m.Epoch,
		Tick:                    m.Tick,
		NumberOfAlignedVotes:    m.NumberOfAlignedVotes,
		NumberOfMisalignedVotes: m.NumberOfMisalignedVotes,
		InitialTick:             m.InitialTick,
	}

	return nil
}
