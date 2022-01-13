package main

import (
	"github.com/Verilog2go-TDD-example/src/variable"
)

type Elelock struct {
	clk, close, tenkey, key, lock *variable.BitArray
}

func NewElelock() Elelock {
	args := &Elelock{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(2), variable.NewBitArray(1), variable.NewBitArray(1)}
	args.clk.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	return *args
}

func NewGoroutineElelock(in []chan int, out []chan int) *Elelock {
	elelock := &Elelock{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(2), variable.NewBitArray(1), variable.NewBitArray(1)}
	go elelock.start(in, out)
	return elelock
}

func (elelock *Elelock) Exec() {
}

func (elelock *Elelock) start(in []chan int, out []chan int) {
	defer close(out[0])
	for {
		select {
		case v, ok := <-in[0]:
			if ok {
				elelock.clk.Set(v)
				bitArrays1 := elelock.PreAlways1()
				elelock.Always1(bitArrays1)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		case v, ok := <-in[1]:
			if ok {
				elelock.close.Set(v)
				bitArrays1 := elelock.PreAlways1()
				elelock.Always1(bitArrays1)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		case v, ok := <-in[2]:
			if ok {
				elelock.tenkey.Set(v)
				bitArrays1 := elelock.PreAlways1()
				elelock.Always1(bitArrays1)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		}
	}
}

func (Elelock *Elelock) PreAlways1() []variable.BitArray {
	var1 := *variable.CreateBitArray(1, Elelock.clk.ToInt())
	var2 := *variable.CreateBitArray(1, Elelock.close.ToInt())
	var3 := *variable.CreateBitArray(2, Elelock.tenkey.ToInt())
	var4 := *variable.CreateBitArray(8, 0)
	var5 := *variable.CreateBitArray(8, 0)
	var6 := *variable.CreateBitArray(8, 0)
	var4.Assign(Elelock.keyenc(*Elelock.tenkey))
	if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.key)) {
		var5.Assign(*variable.CreateBits("1'b0"))
	} else {
		if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.close)) {
			var6.Assign(*variable.CreateBits("1'b1"))
		}
	}
	return []variable.BitArray{var1, var2, var3, var4, var5, var6}
}

func (Elelock *Elelock) Always1(vars []variable.BitArray) {
	Elelock.key.Assign(vars[3])
	if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.key)) {
		Elelock.lock.Assign(vars[4])
	} else {
		if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.close)) {
			Elelock.lock.Assign(vars[5])
		}
	}
}

func (Elelock *Elelock) keyenc(sw variable.BitArray) variable.BitArray {
	keyenc := *variable.CreateBitArray(1, 0)
	switch sw.ToInt() {
	case 1:
		keyenc.Set(0)
	case 2:
		keyenc.Set(1)
	}
	return keyenc
}
