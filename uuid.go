package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}

// Value method
func (nu NullUUID) Value() (driver.Value, error) {

	if !nu.Valid {
		return nil, nil
	}

	return nu.UUID.Value()
}

// Scan method
func (nu *NullUUID) Scan(value interface{}) error {

	if value == nil {
		nu.UUID, nu.Valid = uuid.Nil, false
		return nil
	}

	err := nu.UUID.Scan(value)
	if err != nil {
		nu.UUID, nu.Valid = uuid.Nil, false
		return errors.New("scan uuid is failed")

	}

	nu.Valid = true

	return nil
}

// MarshalJSON method
func (nu NullUUID) MarshalJSON() ([]byte, error) {

	if !nu.Valid {
		return NullType, nil
	}

	return json.Marshal(nu.UUID)
}

// UnmarshalJSON method
func (nu *NullUUID) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*nu = NullUUID{}
		return nil
	}

	err := json.Unmarshal(data, &nu.UUID)
	if err != nil {
		return err
	}

	nu.Valid = true

	return nil
}
