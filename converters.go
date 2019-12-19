package guid

import (
	"encoding/hex"
	"strings"
)

// ToGUID takes
func ToGUID(uuid [16]byte) string {
	var buf [36]byte
	hex.Encode(buf[0:2], uuid[3:4])
	hex.Encode(buf[2:4], uuid[2:3])
	hex.Encode(buf[4:6], uuid[1:2])
	hex.Encode(buf[6:8], uuid[0:1])
	buf[8] = '-'
	hex.Encode(buf[9:11], uuid[5:6])
	hex.Encode(buf[11:13], uuid[4:5])
	buf[13] = '-'
	hex.Encode(buf[14:16], uuid[7:8])
	hex.Encode(buf[16:18], uuid[6:7])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])
	return string(buf[:])
}

// FromGUID takes
func FromGUID(guid string) [16]byte {
	readerMem := make([]byte, 16)
	hex.Decode(readerMem, []byte(strings.ReplaceAll(guid, "-", "")))
	var asBytes [16]byte
	copy(asBytes[:], readerMem[:16])
	asBytes = swap(asBytes, 0, 3)
	asBytes = swap(asBytes, 1, 2)
	asBytes = swap(asBytes, 4, 5)
	asBytes = swap(asBytes, 6, 7)
	return asBytes
}

func swap(uuid [16]byte, swapOne uint8, swapTwo uint8) [16]byte {
	var returnGUID [16]byte
	copy(returnGUID[:], uuid[:])
	returnGUID[swapOne] = uuid[swapTwo]
	returnGUID[swapTwo] = uuid[swapOne]
	return returnGUID
}
