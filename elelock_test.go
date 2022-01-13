package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestElelock01(t *testing.T) {
// 	elelock := NewElelock()

// 	elelock.key.Set(1)
// 	elelock.Exec()
// 	assert.Equal(t, 0, elelock.lock.ToInt())
// }

// func TestElelock02(t *testing.T) {
// 	elelock := NewElelock()

// 	elelock.close.Set(1)
// 	elelock.Exec()
// 	assert.Equal(t, 1, elelock.lock.ToInt())
// }

func TestElelock03(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.key.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 0, elelock.lock.ToInt())
}

func TestElelock04(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
}

func TestElelock05(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
	elelock.clk.Set(0)
	elelock.close.Set(0)
	elelock.tenkey.Set(2)
	elelock.clk.Set(1)
	assert.Equal(t, 0, elelock.lock.ToInt())
	elelock.clk.Set(0)
	elelock.tenkey.Set(0)
	elelock.close.Set(1)
	elelock.clk.Set(1)
	// assert.Equal(t, 1, elelock.lock.ToInt())
}

// func TestElelock06(t *testing.T) {
// 	elelock := NewElelock()

// 	assert.Equal(t, 2, len(elelock.tenkey.GetBits()))
// }

func TestElelock07(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.close.Set(1)
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("2'b10")
	elelock.clk.Set(1)
	assert.Equal(t, 0, elelock.lock.ToInt())
}

func TestElelock08(t *testing.T) {
	elelock := NewElelock()

	assert.Equal(t, 10, len(elelock.tenkey.GetBits()))
}

func TestElelock09(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	assert.Equal(t, 7, elelock.key.ToInt())
}
