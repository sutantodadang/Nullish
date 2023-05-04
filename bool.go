package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type NullBool struct {
	Bool  bool
	Valid bool
}

// Value method
func (nb NullBool) Value() (driver.Value, error) {

	if !nb.Valid {
		return nil, nil
	}

	return nb.Bool, nil
}

// Scan method
func (nb *NullBool) Scan(value interface{}) error {

	if value == nil {
		nb.Bool, nb.Valid = false, false
		return nil
	}

	b, ok := value.(bool)
	if !ok {
		return errors.New("type assertion to bool is failed")
	}

	nb.Bool, nb.Valid = b, true

	return nil
}

// MarshalJSON method
func (nb NullBool) MarshalJSON() ([]byte, error) {

	if !nb.Valid {
		return NullType, nil
	}

	return json.Marshal(nb.Bool)
}

// UnmarshalJSON method
func (nb *NullBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*nb = NullBool{}
		return nil
	}

	var res bool

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*nb = NullBool{Bool: res, Valid: true}

	return nil
}
