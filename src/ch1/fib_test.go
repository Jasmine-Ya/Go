package main

import (
	"fmt"
	"testing"
)

//单元测试 方法名要大写Test
func TestFibonacciCreation(t *testing.T) {
	//变量声明 自动推断 没有使用的变量会报错
	var (
		a = 1
		b = 1
	)
	t.Log(a)
	//for循环
	for i := 1; i < 5; i++ {
		t.Log(b)
		temp := a
		a = b
		b = temp + a
	}
	fmt.Println()
}

func TestChange(t *testing.T) {
	a := 1
	b := 2
	t.Log(a, b)
	//temp:=a
	//a=b
	//b=temp
	a, b = b, a
	t.Log(a, b)
}

//快速设置连续的值
func TestConst(t *testing.T) {

	const (
		Monday = iota + 1
		TuesDay
		Wednesday
	)
	t.Log(Monday, TuesDay, Wednesday)
}
func TestConst1(t *testing.T) {
	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	//0111
	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
