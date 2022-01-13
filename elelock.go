package main

import "github.com/Verilog2go-TDD-example/src/variable"

type Elelock struct {
	key, lock *variable.BitArray
}

func NewElelock() Elelock {
	args := &Elelock{variable.NewBitArray(1), variable.NewBitArray(1)}
	return *args
}

func NewGoroutineElelock(in []chan int, out []chan int) *Elelock {
	elelock := &Elelock{variable.NewBitArray(1), variable.NewBitArray(1)}
	go elelock.start(in, out)
	return elelock
}

func (elelock *Elelock) Exec() {
	elelock.lock.Assign(elelock.key.Not())
}

func (elelock *Elelock) start(in []chan int, out []chan int) {
	defer close(out[0])
	for {
		select {
		case v, ok := <-in[0]:
			if ok {
				elelock.key.Set(v)
				elelock.Exec()
				out[0] <- elelock.lock.ToInt()
			} else {
				return
			}
		}
	}
}
