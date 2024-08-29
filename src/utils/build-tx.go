package utils

import (
	"strconv"

	"google.golang.org/protobuf/proto"

	"chain-communicator/logger"
	"chain-communicator/proto-go/trx"
)

func BuildTransaction(nonce, from, to, amount, gas, gasPrice string, type2 int32, payload []byte) (*trx.TrxProto, error) {

	pb := &trx.TrxProto{}
	pb.Version = 1
	pb.Time = CurrentTime()

	var err error
	if pb.Nonce, err = strconv.ParseUint(nonce, 10, 64); err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	if pb.From, err = HexAddressToBytes(from); err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	if pb.To, err = HexAddressToBytes(to); err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	if pb.XAmount, err = BnStringToBytes(amount); err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	if pb.Gas, err = strconv.ParseUint(gas, 10, 64); err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	if pb.XGasPrice, err = BnStringToBytes(gasPrice); err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	pb.Type = type2
	if len(payload) > 0 {
		message := &trx.TrxPayloadContractProto{
			XData: payload,
		}
		serializedData, err := proto.Marshal(message)
		if err != nil {
			logger.Log.Error(err)
			return nil, err
		}
		pb.XPayload = serializedData
	}

	return pb, nil
}
