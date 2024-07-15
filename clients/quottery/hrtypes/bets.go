package hrtypes

import (
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/quottery/nodetypes"
	"github.com/qubic/go-qubic/common"
	"time"
)

type BetInfo struct {
	ID                      uint32
	NumberOfOptions         uint32
	Creator                 common.Identity
	Description             string
	Options                 []BetOption
	OraclesInfo             []OracleInfo
	OraclesVotes            []OracleVote
	OpenDate                time.Time
	CloseDate               time.Time
	EndDate                 time.Time
	MinimumBetAmount        uint64
	MaximumBetSlotPerOption uint32
}

type OracleInfo struct {
	ID            common.Identity
	FeePercentage float64
}

type OracleVote struct {
	OracleIndex int8
	WonOption   int8
}

type BetOption struct {
	Description string
	State       uint32
}

func (bi *BetInfo) FromNodeType(m nodetypes.BetInfo) error {
	bic := betInfoConverter{rawBetInfo: m}

	converted, err := bic.toType()
	if err != nil {
		return errors.Wrap(err, "converting to type")
	}
	*bi = converted

	return nil
}

type betInfoConverter struct {
	rawBetInfo nodetypes.BetInfo
}

func (bi *betInfoConverter) toType() (BetInfo, error) {
	var creatorID common.Identity
	err := creatorID.FromPubKey(bi.rawBetInfo.Creator, false)
	if err != nil {
		return BetInfo{}, errors.Wrap(err, "converting creator id")
	}

	description := string(bi.rawBetInfo.Description[:])

	maxNOption := 8
	if int(bi.rawBetInfo.NumberOfOptions) < maxNOption {
		maxNOption = int(bi.rawBetInfo.NumberOfOptions)
	}

	betOptions := make([]BetOption, 0, maxNOption)
	for i := 0; i < maxNOption; i++ {
		optDescription := string(bi.rawBetInfo.OptionDescription[i*32 : (i*32)+32])
		state := bi.rawBetInfo.CurrentBetState[i]
		betOptions = append(betOptions, BetOption{Description: optDescription, State: state})
	}

	oraclesInfo := make([]OracleInfo, 0, 8)
	oracleVotes := make([]OracleVote, 0, 8)
	for i := 0; i < 8; i++ {
		var providerPubKey [32]byte
		copy(providerPubKey[:], bi.rawBetInfo.OracleProvider[i*32:(i*32)+32])
		if providerPubKey == [32]byte{} {
			continue
		}

		var providerID common.Identity
		err := providerID.FromPubKey(providerPubKey, false)
		if err != nil {
			return BetInfo{}, errors.Wrapf(err, "converting provider id with pubkey: %s", providerPubKey)
		}
		oi := OracleInfo{
			ID:            providerID,
			FeePercentage: float64(bi.rawBetInfo.OracleFees[i]) / 100,
		}
		oraclesInfo = append(oraclesInfo, oi)

		if bi.rawBetInfo.BetResultWonOption[i] == -1 && bi.rawBetInfo.BetResultOpID[i] == -1 {
			continue
		}

		ov := OracleVote{
			OracleIndex: bi.rawBetInfo.BetResultOpID[i],
			WonOption:   bi.rawBetInfo.BetResultWonOption[i],
		}

		oracleVotes = append(oracleVotes, ov)
	}

	openDate := quotteryParseDate(bi.rawBetInfo.OpenDate)
	closeDate := quotteryParseDate(bi.rawBetInfo.CloseDate)
	endDate := quotteryParseDate(bi.rawBetInfo.EndDate)

	return BetInfo{
		ID:                      bi.rawBetInfo.ID,
		NumberOfOptions:         bi.rawBetInfo.NumberOfOptions,
		Creator:                 creatorID,
		Description:             description,
		Options:                 betOptions,
		OraclesInfo:             oraclesInfo,
		OraclesVotes:            oracleVotes,
		OpenDate:                openDate,
		CloseDate:               closeDate,
		EndDate:                 endDate,
		MinimumBetAmount:        bi.rawBetInfo.MinimumBetAmount,
		MaximumBetSlotPerOption: bi.rawBetInfo.MaximumBetSlotPerOption,
	}, nil
}

func quotteryParseDate(date uint32) time.Time {
	return time.Date(quotteryGetYear(date), time.Month(quotteryGetMonth(date)), quotteryGetDay(date), quotteryGetHour(date), quotteryGetMinute(date), quotteryGetSecond(date), 0, time.UTC)
}

func quotteryGetYear(data uint32) int {
	return 2000 + int((data>>26)+24)
}

func quotteryGetMonth(data uint32) int {
	return int((data >> 22) & 0b1111)
}

func quotteryGetDay(data uint32) int {
	return int((data >> 17) & 0b11111)
}

func quotteryGetHour(data uint32) int {
	return int((data >> 12) & 0b11111)
}

func quotteryGetMinute(data uint32) int {
	return int((data >> 6) & 0b111111)
}

func quotteryGetSecond(data uint32) int {
	return int(data & 0b111111)
}
