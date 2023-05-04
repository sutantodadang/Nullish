package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type NullString struct {
	String string
	Valid  bool
}

// Value method
func (ns NullString) Value() (driver.Value, error) {

	if !ns.Valid {
		return nil, nil
	}

	return ns.String, nil
}

// Scan method
func (ns *NullString) Scan(value interface{}) error {

	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}

	b, ok := value.(string)
	if !ok {
		return errors.New("type assertion to string is failed")
	}

	ns.String, ns.Valid = b, true

	return nil
}

// MarshalJSON - marshaller for json
func (ns NullString) MarshalJSON() ([]byte, error) {

	if !ns.Valid {
		return NullType, nil
	}

	return json.Marshal(ns.String)
}

// UnmarshalJSON - unmarshaller for json
func (ns *NullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*ns = NullString{}
		return nil
	}

	var res string

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*ns = NullString{String: res, Valid: true}

	return nil
}
