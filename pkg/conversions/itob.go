package conversions

import "encoding/binary"

func Itob(valueToConvert uint64) []byte {
	binSlc := make([]byte, 8)
	binary.BigEndian.PutUint64(binSlc, valueToConvert)
	return binSlc
}
