package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type NullInt struct {
	Int   int
	Valid bool
}

// Value method
func (ni NullInt) Value() (driver.Value, error) {

	if !ni.Valid {
		return nil, nil
	}

	return ni.Int, nil
}

// Scan method
func (ni *NullInt) Scan(value interface{}) error {

	if value == nil {
		ni.Int, ni.Valid = 0, false
		return nil
	}

	b, ok := value.(int)
	if !ok {
		return errors.New("type assertion to int is failed")
	}

	ni.Int, ni.Valid = b, true

	return nil
}

// MarshalJSON - marshaller for json
func (ni NullInt) MarshalJSON() ([]byte, error) {

	if !ni.Valid {
		return NullType, nil
	}

	return json.Marshal(ni.Int)
}

// UnmarshalJSON - unmarshaller for json
func (ni *NullInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*ni = NullInt{}
		return nil
	}

	var res int

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*ni = NullInt{Int: res, Valid: true}

	return nil
}
