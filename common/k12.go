package common

import (
	"github.com/cloudflare/circl/xof/k12"
	"github.com/pkg/errors"
)

func K12Hash(data []byte) ([32]byte, error) {
	h := k12.NewDraft10([]byte{}) // Using K12 for hashing, equivalent to KangarooTwelve(temp, 96, h, 64).
	_, err := h.Write(data)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "k12 hashing")
	}

	var out [32]byte
	_, err = h.Read(out[:])
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "reading k12 digest")
	}

	return out, nil
}
