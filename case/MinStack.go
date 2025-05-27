package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	Data    []int
	MinData []int
}

func Constructor() MinStack {
	return MinStack{
		Data:    []int{},
		MinData: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Data = append(this.Data, val)
	if val < this.GetMin() {
		this.MinData = append(this.MinData, val)
	} else {
		this.MinData = append(this.MinData, this.GetMin())
	}
}

func (this *MinStack) Pop() {
	if len(this.Data) > 0 {
		this.Data = this.Data[:len(this.Data)-1]
		this.MinData = this.MinData[:len(this.MinData)-1]
	}
}

func (this *MinStack) Top() int {
	ret := math.MaxInt32
	if len(this.Data) > 0 {
		ret = this.Data[len(this.Data)-1]
	}
	return ret
}

func (this *MinStack) GetMin() int {
	ret := math.MaxInt32
	if len(this.MinData) > 0 {
		ret = this.MinData[len(this.MinData)-1]
	}
	return ret
}

func main() {
	obj := Constructor()

	for _, val := range []int{3, 1, 2} {
		obj.Push(val)
	}
	fmt.Println(obj)
}
