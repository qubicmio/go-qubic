package nodetypes

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/internal/connector"
	"io"
)

const (
	NumberOfTransactionsPerTick = 1024
)

const (
	CurrentTickInfoTypeRequest  = 27
	CurrentTickInfoTypeResponse = 28
	TickDataTypeResponse        = 8
	TickDataTypeRequest         = 16
)

type TickData struct {
	ComputorIndex      uint16
	Epoch              uint16
	Tick               uint32
	Millisecond        uint16
	Second             uint8
	Minute             uint8
	Hour               uint8
	Day                uint8
	Month              uint8
	Year               uint8
	UnionData          [256]byte
	Timelock           [32]byte
	TransactionDigests [NumberOfTransactionsPerTick][32]byte
	ContractFees       [1024]int64
	Signature          [SignatureSize]byte
}

func (td *TickData) UnmarshallFromReader(r io.Reader) error {
	var header connector.RequestResponseHeader

	err := binary.Read(r, binary.BigEndian, &header)
	if err != nil {
		return errors.Wrap(err, "reading tick data from reader")
	}

	if header.Type == connector.EndResponse {
		return nil
	}

	if header.Type != TickDataTypeResponse {
		return errors.Errorf("Invalid header type, expected %d, found %d", TickDataTypeResponse, header.Type)
	}

	err = binary.Read(r, binary.LittleEndian, td)
	if err != nil {
		return errors.Wrap(err, "reading tick data from reader")
	}

	return nil
}

func (td *TickData) IsEmpty() bool {
	if td == nil {
		return true
	}

	return *td == TickData{}
}

type TickInfo struct {
	TickDuration            uint16
	Epoch                   uint16
	Tick                    uint32
	NumberOfAlignedVotes    uint16
	NumberOfMisalignedVotes uint16
	InitialTick             uint32
}

func (ti *TickInfo) UnmarshallFromReader(r io.Reader) error {
	var header connector.RequestResponseHeader

	err := binary.Read(r, binary.BigEndian, &header)
	if err != nil {
		return errors.Wrap(err, "reading header")
	}

	if header.Type != CurrentTickInfoTypeResponse {
		return errors.Errorf("Invalid header type, expected %d, found %d", CurrentTickInfoTypeResponse, header.Type)
	}

	err = binary.Read(r, binary.LittleEndian, ti)
	if err != nil {
		return errors.Wrap(err, "reading tick data from reader")
	}
	return nil
}
