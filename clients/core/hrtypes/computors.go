package hrtypes

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/clients/core/nodetypes"
	"github.com/qubic/go-qubic/common"
)

type ComputorList struct {
	Epoch        uint32
	Identities   []common.Identity
	SignatureHex string
	DigestHex    string
}

func (cl *ComputorList) FromNodeType(m nodetypes.Computors) error {
	cc := computorsConverter{comps: m}

	list, err := cc.toType()
	if err != nil {
		return errors.Wrap(err, "converting to type")
	}

	*cl = list

	return nil
}

type computorsConverter struct {
	comps nodetypes.Computors
}

func (c *computorsConverter) toType() (ComputorList, error) {
	identities, err := common.PubKeysToIdentities(c.comps.PubKeys[:], false)
	if err != nil {
		return ComputorList{}, errors.Wrap(err, "converting pubKeys to identities")
	}

	digest, err := c.getDigest()
	if err != nil {
		return ComputorList{}, errors.Wrap(err, "creating computors digest")
	}

	return ComputorList{
		Epoch:        uint32(c.comps.Epoch),
		Identities:   identities,
		SignatureHex: hex.EncodeToString(c.comps.Signature[:]),
		DigestHex:    hex.EncodeToString(digest[:]),
	}, nil
}

func (c *computorsConverter) getDigest() ([32]byte, error) {
	serialized, err := common.BinarySerializeLE(c.comps)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "serializing data")
	}

	// remove signature from computors data
	computorsData := serialized[:len(serialized)-64]
	digest, err := common.K12Hash(computorsData)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hashing computors data")
	}

	return digest, nil
}
