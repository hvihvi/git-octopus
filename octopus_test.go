package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlags(t *testing.T) {
	assert.Equal(t, "", *pattern)
	assert.Equal(t, ".", *repository)
}
