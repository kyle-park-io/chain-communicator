package encode

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"

	pb "chain-communicator/proto-go/test_trx"
	"chain-communicator/utils"
)

func TestA(t *testing.T) {
	a, err := rlp.EncodeToBytes([][]byte{{1}, {2}})
	if err != nil {
		fmt.Println(err)
	}

	b := utils.BytesToHexString(a)
	fmt.Println(b)

	tS := base64.StdEncoding.EncodeToString(a)
	fmt.Println(tS)
}

type Trx struct {
	Version uint32 `json:"version,omitempty"`
	Nonce   uint64 `json:"nonce"`
	Empty   string `json:"empty"`
}

func EncodeRlp1(msg *pb.TestTrxProto) ([]byte, error) {
	test, err := rlp.EncodeToBytes(msg)
	if err != nil {
		return nil, err
	}
	return test, nil
}

func EncodeRlp2(msg *Trx) ([]byte, error) {
	test, err := rlp.EncodeToBytes(msg)
	if err != nil {
		return nil, err
	}
	return test, nil
}

func TestEncode(t *testing.T) {
	// &{1 1695623550277741000 0 35BA39ABEDFACE7FEC2EBC221E8E48F00206731F 7D2CC4E75E4E52FBBEC3315264FF4557E28FAD0F 1000 4000 250000000000 1 0x1068334c0 8F5B25F1244918E85E73832342B19869A072A1FB666F612787738AF17AB437CA6E9E995DDA1520AE495E2C98A5F8BEEB0721C6E63DF570FA7EEF6095F65F46C201}

	test := &pb.TestTrxProto{}
	test.Version = 1
	test.Nonce = 2

	eD, err := EncodeRlp1(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(eD)

	test2 := &Trx{}
	test2.Version = 1
	test2.Nonce = 2
	eD2, err := EncodeRlp2(test2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(eD2)
}
