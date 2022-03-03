package conversions

import "encoding/binary"

func Btoi(binarySlice []byte) uint64 {
	return binary.BigEndian.Uint64(binarySlice)
}
