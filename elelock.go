package main

import "github.com/Verilog2go-TDD-example/src/variable"

type Elelock struct {
	clk, key, close, lock *variable.BitArray
}

func NewElelock() Elelock {
	args := &Elelock{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1)}
	args.clk.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
	return *args
}

func NewGoroutineElelock(in []chan int, out []chan int) *Elelock {
	elelock := &Elelock{variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1), variable.NewBitArray(1)}
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
				elelock.key.Set(v)
				bitArrays1 := elelock.PreAlways1()
				elelock.Always1(bitArrays1)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		case v, ok := <-in[2]:
			if ok {
				elelock.close.Set(v)
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
	var2 := *variable.CreateBitArray(1, Elelock.key.ToInt())
	var3 := *variable.CreateBitArray(1, Elelock.close.ToInt())
	var4 := *variable.CreateBitArray(8, 0)
	var5 := *variable.CreateBitArray(8, 0)
	if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.key)) {
		var4.Assign(*variable.CreateBits("1'b0"))
	} else {
		if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.close)) {
			var5.Assign(*variable.CreateBits("1'b1"))
		}
	}
	return []variable.BitArray{var1, var2, var3, var4, var5}
}

func (Elelock *Elelock) Always1(vars []variable.BitArray) {
	if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.key)) {
		Elelock.lock.Assign(vars[3])
	} else {
		if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.close)) {
			Elelock.lock.Assign(vars[4])
		}
	}
}
