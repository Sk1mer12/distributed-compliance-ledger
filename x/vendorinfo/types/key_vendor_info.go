package types

import "encoding/binary"

var _ binary.ByteOrder

// VendorInfoKeyPrefix is the prefix to retrieve all VendorInfo.
var VendorInfoKeyPrefix = []byte{0x01} // prefix for each key to a pending account

// VendorInfoKey returns the store key to retrieve a VendorInfo from the index fields.
func VendorInfoKey(
	vendorID int32,
) []byte {
	var key []byte

	vendorIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(vendorIDBytes, uint32(vendorID))
	key = append(key, vendorIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
