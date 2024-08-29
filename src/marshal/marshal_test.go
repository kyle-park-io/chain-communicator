package marshal_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"chain-communicator/encode"
	pb "chain-communicator/proto-go/test_trx"
)

func TestMarshalTestTrxProto(t *testing.T) {
	// tP1
	tP := &pb.TestTrxProto{}
	tP.Version = 1
	tP.Nonce = 2
	bp, err := ProtoMarshal_Example(tP)
	require.NoError(t, err)
	fmt.Println(bp)

	// tP2
	tP2 := &pb.TestTrxProto{}
	tP2.Version = 1
	tP2.Nonce = 2
	tP2.Data = []byte("test")
	bp2, err := ProtoMarshal_Example(tP2)
	require.NoError(t, err)
	fmt.Println(bp2)

	// tP3
	tP3 := &pb.TestTrxProto{}
	tP3.Version = 1
	tP3.Nonce = 2
	eD, err := encode.EncodeRlpInterface([]byte("test"))
	require.NoError(t, err)
	fmt.Println([]byte("test"))
	fmt.Println(eD)
	tP3.Data = eD
	bp3, err := ProtoMarshal_Example(tP3)
	require.NoError(t, err)
	fmt.Println(bp3)
}

func ProtoMarshal_Example(msg *pb.TestTrxProto) ([]byte, error) {
	bp, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return bp, nil
}
