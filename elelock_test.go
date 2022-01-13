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

// func TestElelock03(t *testing.T) {
// 	elelock := NewElelock()

// 	elelock.clk.Set(0)
// 	elelock.key.Set(1)
// 	elelock.clk.Set(1)
// 	assert.Equal(t, 0, elelock.lock.ToInt())
// }

func TestElelock04(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
}

// func TestElelock05(t *testing.T) {
// 	elelock := NewElelock()

// 	elelock.clk.Set(0)
// 	elelock.close.Set(1)
// 	elelock.clk.Set(1)
// 	assert.Equal(t, 1, elelock.lock.ToInt())
// 	elelock.clk.Set(0)
// 	elelock.close.Set(0)
// 	elelock.tenkey.Set(2)
// 	elelock.clk.Set(1)
// 	assert.Equal(t, 0, elelock.lock.ToInt())
// 	elelock.clk.Set(0)
// 	elelock.tenkey.Set(0)
// 	elelock.close.Set(1)
// 	elelock.clk.Set(1)
// 	assert.Equal(t, 1, elelock.lock.ToInt())
// }

// func TestElelock06(t *testing.T) {
// 	elelock := NewElelock()

// 	assert.Equal(t, 2, len(elelock.tenkey.GetBits()))
// }

// func TestElelock07(t *testing.T) {
// 	elelock := NewElelock()

// 	elelock.clk.Set(0)
// 	elelock.close.Set(1)
// 	elelock.clk.Set(1)
// 	elelock.clk.Set(0)
// 	elelock.tenkey.SetBits("2'b10")
// 	elelock.clk.Set(1)
// 	assert.Equal(t, 0, elelock.lock.ToInt())
// }

func TestElelock08(t *testing.T) {
	elelock := NewElelock()

	assert.Equal(t, 10, len(elelock.tenkey.GetBits()))
}

func TestElelock09(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	assert.Equal(t, 7, elelock.key[0].ToInt())
}

func TestElelock10(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	assert.Equal(t, 0, elelock.lock.ToInt())
}

func TestElelock11(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
	assert.Equal(t, 15, elelock.key[0].ToInt())
}

func TestElelock12(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	assert.Equal(t, 0, elelock.lock.ToInt())
	elelock.clk.Set(0)
	elelock.tenkey.Set(0)
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
	assert.Equal(t, 15, elelock.key[0].ToInt())
}

// func TestElelock13(t *testing.T) {
// 	elelock := NewElelock()

// 	assert.Equal(t, 2, len(elelock.key))
// }

func TestElelock14(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	assert.Equal(t, 7, elelock.key[0].ToInt())
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0000001000")
	elelock.clk.Set(1)
	assert.Equal(t, 3, elelock.key[0].ToInt())
	assert.Equal(t, 7, elelock.key[1].ToInt())
}

func TestElelock15(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0000001000")
	elelock.clk.Set(1)
	assert.Equal(t, 0, elelock.lock.ToInt())
}

func TestElelock16(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
	assert.Equal(t, 15, elelock.key[1].ToInt())
	assert.Equal(t, 15, elelock.key[0].ToInt())
}

func TestElelock17(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	// Close the key
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
	assert.Equal(t, 15, elelock.key[1].ToInt())
	assert.Equal(t, 15, elelock.key[0].ToInt())
	elelock.clk.Set(0)
	elelock.close.Set(0)
	// input 7
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	// input 3
	elelock.tenkey.SetBits("10'b0000001000")
	elelock.clk.Set(1)
	// Make sure the lock is open
	assert.Equal(t, 0, elelock.lock.ToInt())
	elelock.clk.Set(0)
	elelock.tenkey.Set(0)
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	// Close the key again
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
	assert.Equal(t, 15, elelock.key[0].ToInt())
}

func TestElelock18(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.reset.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 15, elelock.key[1].ToInt())
	assert.Equal(t, 15, elelock.key[0].ToInt())
}

func TestElelock19(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.reset.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
}

func TestElelock20(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	// Close the key
	elelock.close.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
	assert.Equal(t, 15, elelock.key[1].ToInt())
	assert.Equal(t, 15, elelock.key[0].ToInt())
	elelock.clk.Set(0)
	elelock.close.Set(0)
	// input 7
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	// input 3
	elelock.tenkey.SetBits("10'b0000001000")
	elelock.clk.Set(1)
	// Make sure the lock is open
	assert.Equal(t, 0, elelock.lock.ToInt())
	elelock.clk.Set(0)
	elelock.tenkey.Set(0)
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	// Reset
	elelock.reset.Set(1)
	elelock.clk.Set(1)
	assert.Equal(t, 1, elelock.lock.ToInt())
	assert.Equal(t, 15, elelock.key[0].ToInt())
}

func TestElelock21(t *testing.T) {
	elelock := NewElelock()

	assert.Equal(t, 4, len(elelock.key))
}

func TestElelock22(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0010000000")
	elelock.clk.Set(1)
	assert.Equal(t, 7, elelock.key[0].ToInt())
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0000001000")
	elelock.clk.Set(1)
	assert.Equal(t, 3, elelock.key[0].ToInt())
	assert.Equal(t, 7, elelock.key[1].ToInt())
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0100000000")
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0000000001")
	elelock.clk.Set(1)
	assert.Equal(t, 7, elelock.key[3].ToInt())
	assert.Equal(t, 3, elelock.key[2].ToInt())
	assert.Equal(t, 8, elelock.key[1].ToInt())
	assert.Equal(t, 0, elelock.key[0].ToInt())
}

func TestElelock23(t *testing.T) {
	elelock := NewElelock()

	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0000100000")
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b1000000000")
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0001000000")
	elelock.clk.Set(1)
	elelock.clk.Set(0)
	elelock.tenkey.SetBits("10'b0000001000")
	elelock.clk.Set(1)
	assert.Equal(t, 0, elelock.lock.ToInt())
}
