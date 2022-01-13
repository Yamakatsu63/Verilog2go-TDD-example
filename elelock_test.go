package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElelock01(t *testing.T) {
	elelock := NewElelock()

	elelock.key.Set(1)
	elelock.Exec()
	assert.Equal(t, 0, elelock.lock.ToInt())
}

func TestElelock02(t *testing.T) {
	elelock := NewElelock()

	elelock.close.Set(1)
	elelock.Exec()
	assert.Equal(t, 1, elelock.lock.ToInt())
}

func TestElelock03(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.key.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 0, elelock.lock.ToInt())
}
