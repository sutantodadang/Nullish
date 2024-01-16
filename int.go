package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"strconv"

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

	var res int

	switch b := value.(type) {
	case int8:

		res = int(b)

	case int16:

		res = int(b)

	case int32:

		res = int(b)

	case int64:

		res = int(b)

	case []byte:

		a, err := strconv.Atoi(string(b))
		if err != nil {
			return errors.New("type assertion to int is failed")
		}

		res = a

	}

	ni.Int, ni.Valid = res, true

	return nil
}

// MarshalJSON method
func (ni NullInt) MarshalJSON() ([]byte, error) {

	if !ni.Valid {
		return NullType, nil
	}

	return json.Marshal(ni.Int)
}

// UnmarshalJSON method
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
