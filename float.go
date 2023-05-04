package nullish

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/goccy/go-json"
)

type NullFloat struct {
	Float float64
	Valid bool
}

// Value method
func (nf NullFloat) Value() (driver.Value, error) {

	if !nf.Valid {
		return nil, nil
	}

	return nf.Float, nil
}

// Scan method
func (nf *NullFloat) Scan(value interface{}) error {

	if value == nil {
		nf.Float, nf.Valid = 0, false
		return nil
	}

	b, ok := value.(float64)
	if !ok {
		return errors.New("type assertion to float64 is failed")
	}

	nf.Float, nf.Valid = b, true

	return nil
}

// MarshalJSON - marshaller for json
func (nf NullFloat) MarshalJSON() ([]byte, error) {

	if !nf.Valid {
		return NullType, nil
	}

	return json.Marshal(nf.Float)
}

// UnmarshalJSON - unmarshaller for json
func (nf *NullFloat) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullType) {
		*nf = NullFloat{}
		return nil
	}

	var res float64

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*nf = NullFloat{Float: res, Valid: true}

	return nil
}
