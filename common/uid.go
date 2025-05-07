package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
	"strings"
)

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) UID {
	return UID{localID, objectType, shardID}
}

// Convert to consistent pointer receivers
func (uid *UID) String() string {
	val := uint64(uid.localID)<<28 + uint64(uid.objectType)<<18 + uint64(uid.shardID)
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid *UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid *UID) GetObjectType() int {
	return uid.objectType
}

func (uid *UID) GetShardID() uint32 {
	return uid.shardID
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FF), // Fixed to match the mask in String()
	}

	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (uid *UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(data []byte) error {
	decodedUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))

	if err != nil {
		return err
	}

	uid.localID = decodedUID.localID
	uid.objectType = decodedUID.objectType
	uid.shardID = decodedUID.shardID

	return nil
}

func (uid *UID) Value() (driver.Value, error) {
	if uid == nil {
		return nil, nil
	}

	return int64(uid.localID), nil
}

func (uid *UID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var i uint32

	switch v := value.(type) {
	case int:
		i = uint32(v)
	case int8:
		i = uint32(v)
	case int16:
		i = uint32(v)
	case int32:
		i = uint32(v)
	case int64:
		i = uint32(v)
	case uint:
		i = uint32(v)
	case uint8:
		i = uint32(v)
	case uint16:
		i = uint32(v)
	case uint32:
		i = v
	case uint64:
		i = uint32(v)
	case []byte:
		// Convert bytes to string and parse
		str := string(v)
		decodedUID, err := FromBase58(str)
		if err != nil {
			return err
		}
		uid.localID = decodedUID.localID
		uid.objectType = decodedUID.objectType
		uid.shardID = decodedUID.shardID
		return nil
	case string:
		// Parse string directly
		decodedUID, err := FromBase58(v)
		if err != nil {
			return err
		}
		uid.localID = decodedUID.localID
		uid.objectType = decodedUID.objectType
		uid.shardID = decodedUID.shardID
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into UID", value)
	}

	// For integer types, we assume it represents the localID
	uid.localID = i
	// Default values for other fields if not specified
	uid.objectType = 0
	uid.shardID = 0

	return nil
}
