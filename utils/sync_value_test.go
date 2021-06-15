package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyncValueBool(t *testing.T) {
	flag := NewSyncValueBool()
	flag.Set(true)
	assert.Equal(t, true, flag.Get())
	flag.Set(false)
	assert.Equal(t, false, flag.Get())
}
