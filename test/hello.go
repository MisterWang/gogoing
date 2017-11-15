package test

import (
	"fmt"
)

//接口
type man interface {
	say()
}
type aab struct {
	name string
}

func (a *aab) say() {
	fmt.Println("my name is:", a.name)
}

//成员方法
type linked struct {
	val  string
	next *linked
}

func (list *linked) create() {
	list.val = "header"
	list.next = nil
}

//全局变量
var (
	g uint32
)

func hello() {
	fmt.Println("hello world")

	//变量
	var a int32
	var b uint32
	var c float32
	var d *int32
	e := true
	a, b, c = 1, 2, 1.2
	d = &a

	fmt.Println(a, b, c, d, e)

	//数组切片
	var arr = []int{1, 3, 4, 5, 6}
	// var arr1 [10]int32

	swap(arr, 1, 3)
	fmt.Println(arr)

	//map 键值对
	var m map[string]string
	m = make(map[string]string)
	var m2 = make(map[string]string)

	m["key"] = "val"
	m["temp"] = "temp"
	key, ok := m["key"]
	if ok {
		fmt.Println(key)
	}
	delete(m, "temp")

	fmt.Println(m, m2)
	//map迭代
	for k, v := range m {
		fmt.Println(k, v)
	}

	//函数
	fmt.Println(max(3, 4))

	//方法 结构体与指针访问一致
	var list linked
	list.create()
	fmt.Println(list.val)

	//接口
	var ab man
	ab = new(aab)
	// ab.name = "aab"
	ab.say()
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func swap(arr []int, s int, e int) {
	var temp = arr[s]
	arr[s] = arr[e]
	arr[e] = temp
}
func print(arr []int) {
	var i, size int
	size = len(arr)

	for i = 0; i < size; i++ {
		fmt.Println(arr[i])
	}
}
func printr(arr []int) {
	//迭代
	for i, num := range arr {
		fmt.Println(i, num)
	}
}
