package main
import "github.com/Verilog2go-TDD-example/src/variable"
type Elelock struct{
clk,close,tenkey,lock,match,SECRET *variable.BitArray
key []*variable.BitArray
}

func NewElelock() Elelock{
args := &Elelock{variable.NewBitArray(1),variable.NewBitArray(1),variable.NewBitArray(10),variable.NewBitArray(1),variable.NewBitArray(1),variable.NewBitArray(1),make([]*variable.BitArray, 2)}
args.key[0] = variable.NewBitArray(4)
args.key[1] = variable.NewBitArray(4)
args.clk.AddPosedgeObserver(args.PreAlways1, args.Always1, args.Exec)
args.clk.AddPosedgeObserver(args.PreAlways2, args.Always2, args.Exec)
args.SECRET = variable.CreateBits("4'h7")
return *args
}

func NewGoroutineElelock (in []chan int, out []chan int) *Elelock{
elelock := &Elelock{variable.NewBitArray(1),variable.NewBitArray(1),variable.NewBitArray(10),variable.NewBitArray(1),variable.NewBitArray(1),variable.NewBitArray(1),make([]*variable.BitArray, 2)}
elelock.key[0] = variable.NewBitArray(4)
elelock.key[1] = variable.NewBitArray(4)
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
bitArrays2 := elelock.PreAlways2()
elelock.Always1(bitArrays1)
elelock.Always2(bitArrays2)
elelock.Exec()
out[0] <- elelock.lock.ToInt()
} else {
return 
}
case v, ok := <-in[1]:
if ok {
elelock.close.Set(v)
bitArrays1 := elelock.PreAlways1()
bitArrays2 := elelock.PreAlways2()
elelock.Always1(bitArrays1)
elelock.Always2(bitArrays2)
elelock.Exec()
out[0] <- elelock.lock.ToInt()
} else {
return 
}
case v, ok := <-in[2]:
if ok {
elelock.tenkey.Set(v)
bitArrays1 := elelock.PreAlways1()
bitArrays2 := elelock.PreAlways2()
elelock.Always1(bitArrays1)
elelock.Always2(bitArrays2)
elelock.Exec()
out[0] <- elelock.lock.ToInt()
} else {
return 
}
}
}
}

func (Elelock *Elelock) PreAlways1() []variable.BitArray{
var1 := *variable.CreateBitArray(1, Elelock.clk.ToInt())
var2 := *variable.CreateBitArray(1, Elelock.close.ToInt())
var3 := *variable.CreateBitArray(10, Elelock.tenkey.ToInt())
var4 := *variable.CreateBitArray(1, Elelock.match.ToInt())
var5 := *variable.CreateBitArray(8, 0)
var6 := *variable.CreateBitArray(8, 0)
var5.Assign(*Elelock.key[0])
var6.Assign(Elelock.keyenc(*Elelock.tenkey))
return []variable.BitArray{var1, var2, var3, var4, var5, var6}
}

func (Elelock *Elelock) PreAlways2() []variable.BitArray{
var1 := *variable.CreateBitArray(1, Elelock.clk.ToInt())
var2 := *variable.CreateBitArray(1, Elelock.close.ToInt())
var3 := *variable.CreateBitArray(10, Elelock.tenkey.ToInt())
var4 := *variable.CreateBitArray(1, Elelock.match.ToInt())
var5 := *variable.CreateBitArray(8, 0)
var6 := *variable.CreateBitArray(8, 0)
var7 := *variable.CreateBitArray(8, 0)
if variable.CheckBit(variable.CreateBits("4'h3").Equal(*Elelock.key[1]).And(Elelock.SECRET.Equal(*Elelock.key[0]))) {
var5.Assign(*variable.CreateBits("1'b0"))} else{
if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.close)) {
var6.Assign(*variable.CreateBits("1'b1"))
var7.Assign(*variable.CreateBits("4'b1111"))
}
}
return []variable.BitArray{var1, var2, var3, var4, var5, var6, var7}
}

func (Elelock *Elelock) Always1(vars []variable.BitArray){
Elelock.key[1].Assign(vars[4])
Elelock.key[0].Assign(vars[5])
}

func (Elelock *Elelock) Always2(vars []variable.BitArray){
if variable.CheckBit(variable.CreateBits("4'h3").Equal(*Elelock.key[1]).And(Elelock.SECRET.Equal(*Elelock.key[0]))) {
Elelock.lock.Assign(vars[4])} else{
if variable.CheckBit(variable.CreateBits("1'b1").Equal(*Elelock.close)) {
Elelock.lock.Assign(vars[5])
Elelock.key[0].Assign(vars[6])
}
}
}


func (Elelock *Elelock) keyenc(sw variable.BitArray) variable.BitArray {
keyenc := *variable.CreateBitArray(3, 0)
switch sw.ToInt(){
case 1:
keyenc.Set(0)
case 2:
keyenc.Set(1)
case 4:
keyenc.Set(2)
case 8:
keyenc.Set(3)
case 16:
keyenc.Set(4)
case 32:
keyenc.Set(5)
case 64:
keyenc.Set(6)
case 128:
keyenc.Set(7)
case 256:
keyenc.Set(8)
case 512:
keyenc.Set(9)
}
return keyenc
}
