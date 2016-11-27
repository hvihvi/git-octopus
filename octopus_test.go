package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInitFlags(t *testing.T) {
	pattern, repository := InitFlags()
	assert.Equal(t, "", *pattern)
	assert.Equal(t, ".", *repository)
}

