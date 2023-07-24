package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
	"github.com/oklog/ulid/v2"
)

type NullULID struct {
	ULID  ulid.ULID
	Valid bool
}

// Value method
func (nl NullULID) Value() (driver.Value, error) {

	if !nl.Valid {
		return nil, nil
	}

	return nl.ULID.Value()
}

// Scan method
func (nl *NullULID) Scan(value interface{}) error {

	if value == nil {
		nl.ULID, nl.Valid = ulid.ULID{}, false
		return nil
	}

	err := nl.ULID.Scan(value)
	if err != nil {
		nl.ULID, nl.Valid = ulid.ULID{}, false
		return errors.New("scan ulid is failed")

	}

	nl.Valid = true

	return nil
}

// MarshalJSON method
func (nl NullULID) MarshalJSON() ([]byte, error) {

	if !nl.Valid {
		return NullType, nil
	}

	return json.Marshal(nl.ULID)
}

// UnmarshalJSON method
func (nl *NullULID) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*nl = NullULID{}
		return nil
	}

	err := json.Unmarshal(data, &nl.ULID)
	if err != nil {
		return err
	}

	nl.Valid = true

	return nil
}
