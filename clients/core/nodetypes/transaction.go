package nodetypes

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"github.com/pkg/errors"
	"github.com/qubic/go-qubic/common"
	"github.com/qubic/go-qubic/internal/connector"
	"io"
)

const (
	TickTransactionsTypeRequest  = 29
	TickTransactionsTypeResponse = 24
	TxStatusTypeRequest          = 201
	TxStatusTypeResponse         = 202
)

type Transaction struct {
	SourcePublicKey      [32]byte
	DestinationPublicKey [32]byte
	Amount               int64
	Tick                 uint32
	InputType            uint16
	InputSize            uint16
	Input                []byte
	Signature            [64]byte
}

func (tx *Transaction) GetUnsignedDigest() ([32]byte, error) {
	serialized, err := tx.MarshallBinary()
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "marshalling tx data")
	}

	// create digest with data without signature
	digest, err := common.K12Hash(serialized[:len(serialized)-64])
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hashing tx data")
	}

	return digest, nil
}

func (tx *Transaction) MarshallBinary() ([]byte, error) {
	var buff bytes.Buffer
	_, err := buff.Write(tx.SourcePublicKey[:])
	if err != nil {
		return nil, errors.Wrap(err, "writing source public key to buffer")
	}

	_, err = buff.Write(tx.DestinationPublicKey[:])
	if err != nil {
		return nil, errors.Wrap(err, "writing destination public key to buffer")
	}
	err = binary.Write(&buff, binary.LittleEndian, tx.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "writing amount to buf")
	}

	err = binary.Write(&buff, binary.LittleEndian, tx.Tick)
	if err != nil {
		return nil, errors.Wrap(err, "writing tick to buf")
	}

	err = binary.Write(&buff, binary.LittleEndian, tx.InputType)
	if err != nil {
		return nil, errors.Wrap(err, "writing input type to buf")
	}

	err = binary.Write(&buff, binary.LittleEndian, tx.InputSize)
	if err != nil {
		return nil, errors.Wrap(err, "writing input size to buf")
	}

	_, err = buff.Write(tx.Input)
	if err != nil {
		return nil, errors.Wrap(err, "writing input to buffer")
	}

	_, err = buff.Write(tx.Signature[:])
	if err != nil {
		return nil, errors.Wrap(err, "writing signature to buffer")
	}

	return buff.Bytes(), nil
}

func (tx *Transaction) UnmarshallBinary(r io.Reader) error {
	err := binary.Read(r, binary.LittleEndian, &tx.SourcePublicKey)
	if err != nil {
		return errors.Wrap(err, "reading source public key from reader")
	}

	err = binary.Read(r, binary.LittleEndian, &tx.DestinationPublicKey)
	if err != nil {
		return errors.Wrap(err, "reading destination public key from reader")
	}

	err = binary.Read(r, binary.LittleEndian, &tx.Amount)
	if err != nil {
		return errors.Wrap(err, "reading amount from reader")
	}

	err = binary.Read(r, binary.LittleEndian, &tx.Tick)
	if err != nil {
		return errors.Wrap(err, "reading tick from reader")
	}

	err = binary.Read(r, binary.LittleEndian, &tx.InputType)
	if err != nil {
		return errors.Wrap(err, "reading input type from reader")
	}

	err = binary.Read(r, binary.LittleEndian, &tx.InputSize)
	if err != nil {
		return errors.Wrap(err, "reading input size from reader")
	}

	tx.Input = make([]byte, tx.InputSize)
	err = binary.Read(r, binary.LittleEndian, &tx.Input)
	if err != nil {
		return errors.Wrap(err, "reading input from reader")
	}

	err = binary.Read(r, binary.LittleEndian, &tx.Signature)
	if err != nil {
		return errors.Wrap(err, "reading signature from reader")
	}

	return nil
}

func (tx *Transaction) Digest() ([32]byte, error) {
	serialized, err := tx.MarshallBinary()
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "marshalling tx data")
	}

	digest, err := common.K12Hash(serialized)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hashing tx data")
	}

	return digest, nil
}

func (tx *Transaction) EncodeToBase64() (string, error) {
	txPacket, err := tx.MarshallBinary()
	if err != nil {
		return "", errors.Wrap(err, "binary marshalling")
	}

	return base64.StdEncoding.EncodeToString(txPacket[:]), nil
}

type Transactions []Transaction

func (txs *Transactions) UnmarshallFromReader(r io.Reader) error {
	for {
		var header connector.RequestResponseHeader
		err := binary.Read(r, binary.BigEndian, &header)
		if err != nil {
			return errors.Wrap(err, "reading header")
		}

		if header.Type == connector.EndResponse {
			break
		}

		if header.Type != TickTransactionsTypeResponse {
			return errors.Errorf("Invalid header type, expected %d, found %d", TickTransactionsTypeResponse, header.Type)
		}

		var tx Transaction

		err = tx.UnmarshallBinary(r)
		if err != nil {
			return errors.Wrap(err, "unmarshalling transaction")
		}

		*txs = append(*txs, tx)
	}

	return nil
}

type TransactionStatus struct {
	CurrentTickOfNode  uint32
	Tick               uint32
	TxCount            uint32
	MoneyFlew          [(NumberOfTransactionsPerTick + 7) / 8]byte
	TransactionDigests [][32]byte
}

func (ts *TransactionStatus) UnmarshallFromReader(r io.Reader) error {
	var header connector.RequestResponseHeader

	err := binary.Read(r, binary.BigEndian, &header)
	if err != nil {
		return errors.Wrap(err, "reading header")
	}

	if header.Type != TxStatusTypeResponse {
		return errors.Errorf("Invalid header type, expected %d, found %d", TxStatusTypeResponse, header.Type)
	}

	err = binary.Read(r, binary.LittleEndian, &ts.CurrentTickOfNode)
	if err != nil {
		return errors.Wrap(err, "reading current tick of node")
	}

	err = binary.Read(r, binary.LittleEndian, &ts.Tick)
	if err != nil {
		return errors.Wrap(err, "reading tick")
	}

	err = binary.Read(r, binary.LittleEndian, &ts.TxCount)
	if err != nil {
		return errors.Wrap(err, "reading tx count")
	}

	err = binary.Read(r, binary.LittleEndian, &ts.MoneyFlew)
	if err != nil {
		return errors.Wrap(err, "reading reading money flew")
	}

	ts.TransactionDigests = make([][32]byte, ts.TxCount)
	err = binary.Read(r, binary.LittleEndian, &ts.TransactionDigests)
	if err != nil {
		return errors.Wrap(err, "reading tx digests")
	}

	return nil
}
