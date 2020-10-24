package config

import (
	"testing"

	"github.com/blbgo/testing/assert"
)

func TestConfigReading(t *testing.T) {
	a := assert.New(t)

	c, err := New()

	a.NoError(err)

	v, err := c.Value("Section2", "S2N1")
	a.NoError(err)
	a.Equal("S2N1Value", v)
}
