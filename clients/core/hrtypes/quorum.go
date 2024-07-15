package hrtypes

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/core/nodetypes"
	"github.com/qubic/go-qubic/common"
	"time"
)

type TickQuorumVotes struct {
	SharedVotes                []QuorumGroupedVotes
	SaltedVotesByComputorIndex map[uint16]QuorumSaltedVote
}

type QuorumGroupedVotes struct {
	Vote          QuorumSharedVote
	NumberOfVotes int
}

type QuorumSharedVote struct {
	Epoch                        uint16
	TickNumber                   uint32
	Timestamp                    int64
	PrevResourceTestingDigestHex string
	PrevSpectrumDigestHex        string
	PrevUniverseDigestHex        string
	PrevComputerDigestHex        string
	TxDigestHex                  string
}

type QuorumSaltedVote struct {
	SaltedResourceTestingDigestHex string
	SaltedSpectrumDigestHex        string
	SaltedUniverseDigestHex        string
	SaltedComputerDigestHex        string
	ExpectedNextTickTxDigestHex    string
	SignatureHex                   string
}

func (qd *TickQuorumVotes) FromNodeType(m nodetypes.QuorumVotes) error {
	qc := quorumConverter{quorumVotes: m}

	quorumData, err := qc.toType()
	if err != nil {
		return errors.Wrap(err, "converting to type")
	}

	*qd = quorumData

	return nil
}

type quorumConverter struct {
	quorumVotes nodetypes.QuorumVotes
}

func (qc *quorumConverter) toType() (TickQuorumVotes, error) {
	sharedVotes, err := qc.toSharedVotes()
	if err != nil {
		return TickQuorumVotes{}, errors.Wrap(err, "to shared votes")
	}

	saltedVotes := qc.toSaltedVotes()

	return TickQuorumVotes{
		SharedVotes:                sharedVotes,
		SaltedVotesByComputorIndex: saltedVotes,
	}, nil
}

func (qc *quorumConverter) toSaltedVotes() map[uint16]QuorumSaltedVote {
	saltedVotes := make(map[uint16]QuorumSaltedVote)

	for _, vote := range qc.quorumVotes {
		saltedVotes[vote.ComputorIndex] = qc.toSaltedVote(vote)
	}

	return saltedVotes
}

func (qc *quorumConverter) toSaltedVote(vote nodetypes.QuorumTickVote) QuorumSaltedVote {
	return QuorumSaltedVote{
		SaltedResourceTestingDigestHex: convertUint64ToHex(vote.SaltedResourceTestingDigest),
		SaltedSpectrumDigestHex:        hex.EncodeToString(vote.SaltedSpectrumDigest[:]),
		SaltedUniverseDigestHex:        hex.EncodeToString(vote.SaltedUniverseDigest[:]),
		SaltedComputerDigestHex:        hex.EncodeToString(vote.SaltedComputerDigest[:]),
		ExpectedNextTickTxDigestHex:    hex.EncodeToString(vote.ExpectedNextTickTxDigest[:]),
		SignatureHex:                   hex.EncodeToString(vote.Signature[:]),
	}
}

func (qc *quorumConverter) toSharedVotes() ([]QuorumGroupedVotes, error) {
	votesHeatMap := make(map[[32]byte][]QuorumSharedVote)
	for _, qv := range qc.quorumVotes {
		hv := heatmapVote{
			Epoch:                         qv.Epoch,
			Tick:                          qv.Tick,
			Millisecond:                   qv.Millisecond,
			Second:                        qv.Second,
			Minute:                        qv.Minute,
			Hour:                          qv.Hour,
			Day:                           qv.Day,
			Month:                         qv.Month,
			Year:                          qv.Year,
			PreviousResourceTestingDigest: qv.PreviousResourceTestingDigest,
			PreviousSpectrumDigest:        qv.PreviousSpectrumDigest,
			PreviousUniverseDigest:        qv.PreviousUniverseDigest,
			PreviousComputerDigest:        qv.PreviousComputerDigest,
			TxDigest:                      qv.TxDigest,
		}

		digest, err := hv.digest()
		if err != nil {
			return nil, errors.Wrap(err, "getting digest")
		}

		sv := qc.toSharedVote(qv)
		if votes, ok := votesHeatMap[digest]; !ok {
			votesHeatMap[digest] = []QuorumSharedVote{sv}
		} else {
			votesHeatMap[digest] = append(votes, sv)
		}
	}

	sharedVotes := make([]QuorumGroupedVotes, 0, len(votesHeatMap))
	for _, votes := range votesHeatMap {
		sharedVotes = append(sharedVotes, QuorumGroupedVotes{Vote: votes[0], NumberOfVotes: len(votes)})
	}

	return sharedVotes, nil
}

func (qc *quorumConverter) toSharedVote(vote nodetypes.QuorumTickVote) QuorumSharedVote {
	date := time.Date(2000+int(vote.Year), time.Month(vote.Month), int(vote.Day), int(vote.Hour), int(vote.Minute), int(vote.Second), 0, time.UTC)
	date.Add(time.Duration(vote.Millisecond) * time.Millisecond)

	return QuorumSharedVote{
		Epoch:                        vote.Epoch,
		TickNumber:                   vote.Tick,
		Timestamp:                    date.UnixMilli(),
		PrevResourceTestingDigestHex: convertUint64ToHex(vote.PreviousResourceTestingDigest),
		PrevSpectrumDigestHex:        hex.EncodeToString(vote.PreviousSpectrumDigest[:]),
		PrevUniverseDigestHex:        hex.EncodeToString(vote.PreviousUniverseDigest[:]),
		PrevComputerDigestHex:        hex.EncodeToString(vote.PreviousComputerDigest[:]),
		TxDigestHex:                  hex.EncodeToString(vote.TxDigest[:]),
	}
}

func (qc *quorumConverter) getVoteDigest(vote nodetypes.QuorumTickVote) ([32]byte, error) {
	// xor computor index with 8
	vote.ComputorIndex ^= 3

	sData, err := common.BinarySerializeLE(vote)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "serializing data")
	}

	tickData := sData[:len(sData)-64]
	digest, err := common.K12Hash(tickData)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hashing tick data")
	}

	return digest, nil
}

type heatmapVote struct {
	Epoch                         uint16
	Tick                          uint32
	Millisecond                   uint16
	Second                        uint8
	Minute                        uint8
	Hour                          uint8
	Day                           uint8
	Month                         uint8
	Year                          uint8
	PreviousResourceTestingDigest uint64
	PreviousSpectrumDigest        [32]byte
	PreviousUniverseDigest        [32]byte
	PreviousComputerDigest        [32]byte
	TxDigest                      [32]byte
}

func (hv *heatmapVote) digest() ([32]byte, error) {
	b, err := common.BinarySerializeLE(hv)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "serializing vote")
	}

	digest, err := common.K12Hash(b)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hashing vote")
	}

	return digest, nil
}

func convertUint64ToHex(value uint64) string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, value)
	return hex.EncodeToString(b)
}
