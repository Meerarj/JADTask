package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	expec, err := Kafkafunc()
	assert.Nil(t, err)
	assert.NotNil(t, expec)
	assert.Equal(t, expec.Address, "ISED")
}
func TestFail(t *testing.T) {
	expec, err := Kafkafunc()
	assert.Nil(t, err)
	assert.NotNil(t, expec)
}
