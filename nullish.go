package nullish

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

var NullType = []byte("null")

func NewNullFloat(float float64, valid bool) NullFloat {
	return NullFloat{
		Float: float,
		Valid: valid,
	}
}

func NewNullInt(integer int, valid bool) NullInt {
	return NullInt{
		Int:   integer,
		Valid: valid,
	}
}

func NewNullString(str string, valid bool) NullString {
	return NullString{
		String: str,
		Valid:  valid,
	}
}

func NewNullTime(time time.Time, valid bool) NullTime {
	return NullTime{
		Time:  time,
		Valid: valid,
	}
}

func NewNullBool(boolean bool, valid bool) NullBool {
	return NullBool{
		Bool:  boolean,
		Valid: valid,
	}
}

func NewNullObj(object map[string]interface{}, valid bool) NullObj {
	return NullObj{
		Obj:   object,
		Valid: valid,
	}
}

func NewNullArr(array []interface{}, valid bool) NullArr {
	return NullArr{
		Arr:   array,
		Valid: valid,
	}
}

func NewNullArrObj(arrayObject []map[string]interface{}, valid bool) NullArrObj {
	return NullArrObj{
		ArrObj: arrayObject,
		Valid:  valid,
	}
}

func NewNullJSON(json json.RawMessage, valid bool) NullJSON {
	return NullJSON{
		Json:  json,
		Valid: valid,
	}
}

func NewNullUUID(uuid uuid.UUID, valid bool) NullUUID {
	return NullUUID{
		UUID:  uuid,
		Valid: valid,
	}
}

func NewNullULID(ulid ulid.ULID, valid bool) NullULID {
	return NullULID{
		ULID:  ulid,
		Valid: valid,
	}
}
