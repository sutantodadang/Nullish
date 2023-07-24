package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/goccy/go-json"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

// Value method
func (nt NullTime) Value() (driver.Value, error) {

	if !nt.Valid {
		return nil, nil
	}

	return nt.Time, nil
}

// Scan method
func (nt *NullTime) Scan(value interface{}) error {

	if value == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return nil
	}

	b, ok := value.(time.Time)
	if !ok {
		return errors.New("type assertion to time is failed")
	}

	nt.Time, nt.Valid = b, true

	return nil
}

// MarshalJSON method
func (nt NullTime) MarshalJSON() ([]byte, error) {

	if !nt.Valid {
		return NullType, nil
	}

	return json.Marshal(nt.Time.Format(time.RFC3339Nano))
}

// UnmarshalJSON method
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*nt = NullTime{}
		return nil
	}

	var res string

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	pTime, err := time.Parse(time.RFC3339Nano, res)
	if err != nil {
		return err
	}

	*nt = NullTime{Time: pTime, Valid: true}

	return nil
}
