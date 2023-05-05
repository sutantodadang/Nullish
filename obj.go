package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type NullObj struct {
	Obj   map[string]interface{}
	Valid bool
}

// Value method
func (no NullObj) Value() (driver.Value, error) {

	if !no.Valid {
		return nil, nil
	}

	return json.Marshal(no.Obj)
}

// Scan method
func (no *NullObj) Scan(value interface{}) error {

	if value == nil {
		no.Obj, no.Valid = map[string]interface{}{}, false
		return nil
	}

	b, ok := value.(map[string]interface{})
	if !ok {
		return errors.New("type assertion to object is failed")
	}

	no.Obj, no.Valid = b, true

	return nil
}

// MarshalJSON method
func (no NullObj) MarshalJSON() ([]byte, error) {

	if !no.Valid {
		return NullType, nil
	}

	return json.Marshal(no.Obj)
}

// UnmarshalJSON method
func (no *NullObj) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*no = NullObj{}
		return nil
	}

	var res map[string]interface{}

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*no = NullObj{Obj: res, Valid: true}

	return nil
}
