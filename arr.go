package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type NullArr struct {
	Arr   []interface{}
	Valid bool
}

// Value method
func (na NullArr) Value() (driver.Value, error) {

	if !na.Valid {
		return nil, nil
	}

	return json.Marshal(na.Arr)
}

// Scan method
func (na *NullArr) Scan(value interface{}) error {

	if value == nil {
		na.Arr, na.Valid = []interface{}{}, false
		return nil
	}

	b, ok := value.([]interface{})
	if !ok {
		return errors.New("type assertion to array is failed")
	}

	na.Arr, na.Valid = b, true

	return nil
}

// MarshalJSON method
func (na NullArr) MarshalJSON() ([]byte, error) {

	if !na.Valid {
		return NullType, nil
	}

	return json.Marshal(na.Arr)
}

// UnmarshalJSON method
func (na *NullArr) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*na = NullArr{}
		return nil
	}

	var res []interface{}

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*na = NullArr{Arr: res, Valid: true}

	return nil
}
