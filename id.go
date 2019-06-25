package snowflake

import (
	"encoding/base64"
	"encoding/binary"
	"strconv"
	"time"
)

// ID uint64
type ID uint64

// Uint64 returns an uint64 of the snowflake ID
func (i ID) Uint64() uint64 {
	return uint64(i)
}

// ParseInt64 converts an uint64 into a snowflake ID
func ParseInt64(id uint64) ID {
	return ID(id)
}

// String returns a string of the snowflake ID
func (i ID) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(id string) (ID, error) {
	i, err := strconv.ParseUint(id, 10, 64)
	return ID(i), err
}

// Base2 returns a string base2 of the snowflake ID
func (i ID) Base2() string {
	return strconv.FormatUint(uint64(i), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(id string) (ID, error) {
	i, err := strconv.ParseUint(id, 2, 64)
	return ID(i), err
}

// Base64 returns a base64 string of the snowflake ID
func (i ID) Base64() string {
	return base64.StdEncoding.EncodeToString(i.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(id string) (ID, error) {
	b, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return 0, err
	}
	return ParseBytes(b)
}

// Bytes returns a byte slice of the snowflake ID
func (i ID) Bytes() []byte {
	return []byte(i.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(id []byte) (ID, error) {
	i, err := strconv.ParseUint(string(id), 10, 64)
	return ID(i), err
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (i ID) IntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(i))
	return b
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(id [8]byte) ID {
	return ID(uint64(binary.BigEndian.Uint64(id[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
func (i ID) Time() int64 {
	ts := int64(uint64(i) >> timestampLeftShiftBits)
	return epoch.Add(time.Duration(ts) * time.Millisecond).UnixNano()
}

// Node returns an uint64 of the snowflake ID node number
func (i ID) Node() uint64 {
	return (uint64(i) >> workeridLeftShiftBits) & (-1 ^ (-1 << nodeidBits))
}

// Sequence returns an uint64 of the snowflake sequence number
func (i ID) Sequence() uint64 {
	return uint64(i) & sequenceMask
}
