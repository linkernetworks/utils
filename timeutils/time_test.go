package timeutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeNesChange(t *testing.T) {
	t1, err := time.Parse(time.RFC3339Nano, "2017-12-18T11:23:10.244286455+08:00")
	assert.NoError(t, err)
	expect1, err := time.Parse(time.RFC3339Nano, "2017-12-18T11:23:10.244000000+08:00")
	assert.NoError(t, err)
	assert.Equal(t, TruncateTime(&t1, time.Millisecond), &expect1)
}
