package marshal

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"

	"chain-communicator/proto-go/trx"
)

func ProtoMarshal(msg *trx.TrxProto) ([]byte, error) {
	bp, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return bp, nil
}

func ProtoUnmarshal(bp []byte, msg *trx.TrxProto) (*trx.TrxProto, error) {
	err := proto.Unmarshal(bp, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func JsonMarshal(msg any) ([]byte, error) {
	bj, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return bj, nil
}

func JsonUnmarshal(bj []byte, msg any) ([]byte, error) {
	err := json.Unmarshal(bj, msg)
	if err != nil {
		return nil, err
	}
	return bj, nil
}
