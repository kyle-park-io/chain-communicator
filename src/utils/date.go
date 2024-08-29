package utils

import (
	"time"
)

func CurrentTime() int64 {
	cT := time.Now().Round(0).UTC().UnixNano()
	// cT := time.Now().UnixNano()
	// big := new(big.Int)
	// big.SetInt64(cT)
	// return big.Bytes()

	return cT
}
