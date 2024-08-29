package encode

import (
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/rlp"

	"chain-communicator/logger"
	"chain-communicator/proto-go/trx"
	"chain-communicator/types"
	"chain-communicator/utils"
)

func EncodeRlp(msg *trx.TrxProto, payload []byte) ([]byte, error) {
	v := reflect.ValueOf(msg)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	var data [][]byte
	for i := 3; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() != reflect.Slice {
			switch field.Kind() {
			case reflect.Int32:
				str := strconv.Itoa(int(field.Interface().(int32)))
				sB, err := utils.BnStringToBytes(str)
				if err != nil {
					logger.Log.Error(err)
					return nil, err
				}
				data = append(data, sB)
			case reflect.Uint32:
				str := strconv.Itoa(int(field.Interface().(uint32)))
				sB, err := utils.BnStringToBytes(str)
				if err != nil {
					logger.Log.Error(err)
					return nil, err
				}
				data = append(data, sB)
			case reflect.Int64:
				str := strconv.Itoa(int(field.Interface().(int64)))
				sB, err := utils.BnStringToBytes(str)
				if err != nil {
					logger.Log.Error(err)
					return nil, err
				}
				data = append(data, sB)
			case reflect.Uint64:
				str := strconv.Itoa(int(field.Interface().(uint64)))
				sB, err := utils.BnStringToBytes(str)
				if err != nil {
					logger.Log.Error(err)
					return nil, err
				}
				data = append(data, sB)
			}
		} else {
			if i == 12 && len(payload) > 0 {
				pS := &types.TrxPayloadContract{
					Data: payload,
				}
				e, err := rlp.EncodeToBytes(pS)
				if err != nil {
					logger.Log.Error(err)
					return nil, err
				}
				data = append(data, e)
			} else {
				data = append(data, field.Bytes())
			}
		}
	}

	eB, err := rlp.EncodeToBytes(data)
	if err != nil {
		return nil, err
	}
	return eB, nil
}

func DecodeRlp(br []byte, msg interface{}) (interface{}, error) {
	err := rlp.DecodeBytes(br, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func EncodeRlpInterface(msg interface{}) ([]byte, error) {
	test, err := rlp.EncodeToBytes(msg)
	if err != nil {
		return nil, err
	}
	return test, nil
}
