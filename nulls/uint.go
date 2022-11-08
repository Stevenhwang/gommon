package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// UInt adds an implementation for int
// that supports proper JSON encoding/decoding.
type UInt struct {
	UInt  uint
	Valid bool // Valid is true if Int is not NULL
}

// Interface implements the nullable interface. It returns nil if
// the uint is not valid, otherwise it returns the uint value.
func (ns UInt) Interface() interface{} {
	if !ns.Valid {
		return nil
	}
	return ns.UInt
}

// NewUInt returns a new, properly instantiated
// Int object.
func NewUInt(i uint) UInt {
	return UInt{UInt: i, Valid: true}
}

// Scan implements the Scanner interface.
func (ns *UInt) Scan(value interface{}) error {
	n := sql.NullInt64{Int64: int64(ns.UInt)}
	err := n.Scan(value)
	ns.UInt, ns.Valid = uint(n.Int64), n.Valid
	return err
}

// Value implements the driver Valuer interface.
func (ns UInt) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return int64(ns.UInt), nil
}

// MarshalJSON marshals the underlying value to a
// proper JSON representation.
func (ns UInt) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.UInt)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON will unmarshal a JSON value into
// the propert representation of that value.
func (ns *UInt) UnmarshalJSON(text []byte) error {
	txt := string(text)
	ns.Valid = true
	if txt == "null" {
		ns.Valid = false
		return nil
	}
	i, err := strconv.ParseUint(txt, 10, 64)
	if err != nil {
		ns.Valid = false
		return err
	}
	j := uint(i)
	ns.UInt = j
	return nil
}

// UnmarshalText will unmarshal text value into
// the propert representation of that value.
func (ns *UInt) UnmarshalText(text []byte) error {
	return ns.UnmarshalJSON(text)
}
