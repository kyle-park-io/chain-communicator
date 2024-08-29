package types

import (
	"io"

	"github.com/ethereum/go-ethereum/rlp"
)

type TrxPayloadContract struct {
	Data []byte `json:"data"`
}

func (tx *TrxPayloadContract) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, tx.Data)
}

func (tx *TrxPayloadContract) DecodeRLP(s *rlp.Stream) error {
	bz, err := s.Bytes()
	if err != nil {
		return err
	}
	tx.Data = bz
	return nil
}
