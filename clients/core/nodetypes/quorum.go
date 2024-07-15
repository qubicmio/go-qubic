package nodetypes

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/internal/connector"
	"io"
)

const (
	MinimumQuorumVotes = 451
)

const (
	QuorumTickTypeRequest  = 14
	QuorumTickTypeResponse = 3
)

type QuorumTickVote struct {
	ComputorIndex uint16
	Epoch         uint16
	Tick          uint32

	Millisecond uint16
	Second      uint8
	Minute      uint8
	Hour        uint8
	Day         uint8
	Month       uint8
	Year        uint8

	PreviousResourceTestingDigest uint64
	SaltedResourceTestingDigest   uint64

	PreviousSpectrumDigest [32]byte
	PreviousUniverseDigest [32]byte
	PreviousComputerDigest [32]byte

	SaltedSpectrumDigest [32]byte
	SaltedUniverseDigest [32]byte
	SaltedComputerDigest [32]byte

	TxDigest                 [32]byte
	ExpectedNextTickTxDigest [32]byte

	Signature [SignatureSize]byte
}

type QuorumVotes []QuorumTickVote

func (qv *QuorumVotes) UnmarshallFromReader(r io.Reader) error {
	for {
		var header connector.RequestResponseHeader
		err := binary.Read(r, binary.BigEndian, &header)
		if err != nil {
			return errors.Wrap(err, "reading header")
		}

		if header.Type == connector.EndResponse {
			break
		}

		var qtd QuorumTickVote
		if header.Type != QuorumTickTypeResponse {
			return errors.Errorf("Invalid header type, expected %d, found %d", QuorumTickTypeResponse, header.Type)
		}

		err = binary.Read(r, binary.LittleEndian, &qtd)
		if err != nil {
			return errors.Wrap(err, "reading quorum tick data from reader")
		}

		*qv = append(*qv, qtd)
	}

	return nil
}
