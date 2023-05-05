package nullish

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func Test_NullJSON(t *testing.T) {

	testData := []map[string]interface{}{
		{
			"foo": "bar",
		},
	}

	b, err := json.Marshal(&testData)
	assert.NoError(t, err)

	jsonNull := NewNullJSON(b, true)

	assert.True(t, jsonNull.Valid)

	var newTestData []map[string]interface{}

	nb, err := jsonNull.MarshalJSON()
	assert.NoError(t, err)

	err = json.Unmarshal(nb, &newTestData)
	assert.NoError(t, err)

	assert.Equal(t, testData, newTestData)

}
