package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type RawBytes []byte

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

	switch s := value.(type) {
	case *string:
		if s == nil {
			return errors.New("type assertion to *string is failed")
		}

		ns.String, ns.Valid = *s, true
		return nil

	case *[]byte:
		if s == nil {
			return errors.New("type assertion to *byte is failed")
		}

		ns.String, ns.Valid = string(*s), true

		return nil

	case *RawBytes:
		if s == nil {
			return errors.New("type assertion to *RawBytes is failed")
		}

		*s = append((*s)[:0], []byte(ns.String)...)

		ns.String, ns.Valid = string(*s), true

		return nil
	}

	return nil
}

// MarshalJSON method
func (ns NullString) MarshalJSON() ([]byte, error) {

	if !ns.Valid {
		return NullType, nil
	}

	return json.Marshal(ns.String)
}

// UnmarshalJSON method
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
