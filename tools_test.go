package schreibvogel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	assert := assert.New(t)
	tables := []struct {
		errors bool
		needs  []string
		there  []string
	}{
		{
			errors: false,
			needs:  []string{"foo"},
			there:  []string{"foo", "bar"},
		},
		{
			errors: true,
			needs:  []string{"zonk"},
			there:  []string{"foo", "bar"},
		},
	}

	for _, table := range tables {
		err := InSlice(table.needs, table.there)

		if table.errors {
			assert.NotNil(err)
			assert.EqualError(err, fmt.Sprintf("%q item not in slice", table.needs[0]))
		} else if !table.errors {
			assert.Nil(err)
		}
	}
}
