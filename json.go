package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type NullJSON struct {
	Json  json.RawMessage
	Valid bool
}

// Value method
func (nj NullJSON) Value() (driver.Value, error) {

	if !nj.Valid {
		return nil, nil
	}

	return json.Marshal(nj.Json)
}

// Scan method
func (nj *NullJSON) Scan(value interface{}) error {

	if value == nil {
		nj.Json, nj.Valid = json.RawMessage{}, false
		return nil
	}

	var res []byte

	switch t := value.(type) {

	case string:
		res = []byte(t)

	case []byte:
		if len(t) == 0 {
			res = NullType
		} else {
			res = []byte(string(t))
		}

	default:
		return errors.New("invalid type json")

	}

	nj.Json, nj.Valid = res, true

	return nil
}

// MarshalJSON method
func (nj NullJSON) MarshalJSON() ([]byte, error) {

	if !nj.Valid {
		return NullType, nil
	}

	return json.Marshal(nj.Json)
}

// UnmarshalJSON method
func (nj *NullJSON) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*nj = NullJSON{}
		return nil
	}

	*nj = NullJSON{Json: data, Valid: true}

	return nil
}
