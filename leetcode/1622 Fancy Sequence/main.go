package main

import (
	"fmt"
	"math/rand"
)

const modulo = 1000000007

// returns x, y: a*x + b*y = gcd(a, b)
func egcd(a, b uint64) (int64, int64) {
	if a < b {
		return egcd(b, a)
	}
	if b == 0 {
		return 1, 0
	}

	c := a % b
	r1, r2 := egcd(b, c)
	// r1 * b + r2 * c = gcd(b, c)
	// r1 * b + r2 * (a % b) = gcd(a, b)
	// r1 * b + r2 * [(a % b) + a/b * b - a/b * b] = gcd(a, b)
	// r1 * b + r2 * [a - a/b * b] = gcd(a, b)
	// r2 * a + r1 * b - r2 * a/b * b = gcd(a, b)
	// r2 * a + [r1 - r2 * a/b] * b = gcd(a, b)

	return r2, r1 - r2 * int64(a/b)
}

type Fancy struct {
	buffer []int
	off int // offset
	mult int// multiplier
	i int // zeroified position
}


func Constructor() Fancy {
	return Fancy{mult : 1}
}


func (this *Fancy) Append(val int)  {
	offsetted := (modulo - uint64(this.off) + uint64(val)) % modulo
	_, multinv := egcd(modulo, uint64(this.mult))
	multinv = (modulo + multinv) % modulo
	this.buffer = append(this.buffer, int(offsetted * uint64(multinv) % modulo))
}


func (this *Fancy) AddAll(inc int)  {
	this.off = int((uint64(this.off) + uint64(inc)) % modulo)
}


func (this *Fancy) MultAll(m int)  {
	if (m == 0) {
		this.i = len(this.buffer)
		this.mult = 1
		this.off = 0
		return
	}
	this.mult = int(uint64(this.mult) * uint64(m) % modulo)
	this.off =  int(uint64(this.off) * uint64(m) % modulo)
}


func (this *Fancy) GetIndex(idx int) int {
	if idx < this.i {
		return this.off
	}
	if idx >= len(this.buffer) {
		return -1
	}
	multed := (uint64(this.buffer[idx]) * uint64(this.mult)) % modulo
	return int((multed + uint64(this.off)) % modulo)
}

func (this *Fancy) ToList(len int) []int {
	ret := make([]int, len)
	for i := range ret {
		ret[i] = this.GetIndex(i)
	}
	return ret
}


/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */

func test_session(opNumber int) {
	raw := make([]int, 0)
	iq := Constructor()
	ops := make([]string, 0)
	for iter := 0; iter < opNumber; iter++ {
		val := int(rand.Uint32() % 100)
		switch rand.Uint32() % 3 {
		case 0: // append
			raw = append(raw, val)
			iq.Append(val)
			ops = append(ops, fmt.Sprintf("Append %d", val))
		case 1: // AddAll
			for i := range raw {
				raw[i] = (raw[i] + val) % modulo
			}
			iq.AddAll(val)
			ops = append(ops, fmt.Sprintf("AddAll %d", val))
		case 2: // MultAll
			for i := range raw {
				raw[i] = int((int64(raw[i]) * int64(val)) % modulo)
			}
			iq.MultAll(val)
			ops = append(ops, fmt.Sprintf("MultAll %d", val))
		}

		actual := iq.ToList(len(raw))
		for i := range raw {
			if raw[i] != actual[i] {
				panic(fmt.Errorf("oh no"))
			}
		}
		if iq.GetIndex(len(raw)) != -1 {
			panic(fmt.Errorf("oh no no no no no"))
		}
	}
}

func main() {
	test_session(100000)

	//obj := Constructor()
	//obj.Append(3)
	//obj.AddAll(6)
	//obj.MultAll(2)
	//param4 := obj.GetIndex(0)
	//fmt.Printf("%v\n", param4)
}
