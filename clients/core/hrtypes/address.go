package hrtypes

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/core/nodetypes"
	"github.com/qubic/go-qubic/common"
)

type AddressData struct {
	ID                         common.Identity
	IncomingAmount             int64
	OutgoingAmount             int64
	NumberOfIncomingTransfers  uint32
	NumberOfOutgoingTransfers  uint32
	LatestIncomingTransferTick uint32
	LatestOutgoingTransferTick uint32
}

type AddressInfo struct {
	AddressData   AddressData
	Tick          uint32
	SpectrumIndex int32
	Siblings      []common.Identity
}

func (ai *AddressInfo) FromNodeType(m nodetypes.AddressInfo) error {
	aic := addressInfoConverter{rawAddressInfo: m}
	converted, err := aic.toType()
	if err != nil {
		return errors.Wrap(err, "converting to type")
	}

	*ai = converted

	return nil
}

type addressInfoConverter struct {
	rawAddressInfo nodetypes.AddressInfo
}

func (c *addressInfoConverter) toType() (AddressInfo, error) {
	id, err := common.GetIDFrom32Bytes(c.rawAddressInfo.AddressData.PublicKey, false)
	if err != nil {
		return AddressInfo{}, errors.Wrapf(err, "getting address id from pubkey hex: %s", hex.EncodeToString(c.rawAddressInfo.AddressData.PublicKey[:]))
	}

	siblings := make([]common.Identity, 0, nodetypes.SpectrumDepth)
	for _, sibling := range c.rawAddressInfo.Siblings {
		if sibling == [32]byte{} {
			continue
		}
		siblingID, err := common.GetIDFrom32Bytes(sibling, false)
		if err != nil {
			return AddressInfo{}, errors.Wrapf(err, "getting address id from sibling hex: %s", hex.EncodeToString(sibling[:]))
		}
		siblings = append(siblings, siblingID)
	}

	return AddressInfo{
		AddressData: AddressData{
			ID:                         id,
			IncomingAmount:             c.rawAddressInfo.AddressData.IncomingAmount,
			OutgoingAmount:             c.rawAddressInfo.AddressData.OutgoingAmount,
			NumberOfIncomingTransfers:  c.rawAddressInfo.AddressData.NumberOfIncomingTransfers,
			NumberOfOutgoingTransfers:  c.rawAddressInfo.AddressData.NumberOfOutgoingTransfers,
			LatestIncomingTransferTick: c.rawAddressInfo.AddressData.LatestIncomingTransferTick,
			LatestOutgoingTransferTick: c.rawAddressInfo.AddressData.LatestOutgoingTransferTick,
		},
		Tick:          c.rawAddressInfo.Tick,
		SpectrumIndex: c.rawAddressInfo.SpectrumIndex,
		Siblings:      siblings,
	}, nil
}
