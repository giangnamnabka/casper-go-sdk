package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/giangnamnabka/casper-go-sdk/types"
)

func Test_DurationUnmarshal_withSpace_shouldBeParsed(t *testing.T) {
	value := `"2h 46m 40s"`
	var result types.Duration
	err := json.Unmarshal([]byte(value), &result)
	require.NoError(t, err)
	data, err := result.MarshalJSON()
	require.NoError(t, err)
	assert.Equal(t, `"2h46m40s"`, string(data))
}
