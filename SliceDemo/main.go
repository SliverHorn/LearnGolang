package main

import "fmt"

// 切片(slice)

func frist() {
	var a []string
	var b []int
	var c = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func second() {
	// 基于数组得到切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Printf("Type:%T, Value:%#v\n", b, b)
	// 切片再切片
	c := b[0:len(b)]
	fmt.Printf("Type:%T, Value:%#v\n", c, c)
}

func third() {
	// make函数构造切片,
	// make(切片类型, 切片长度, 切片容量), 不写容量则切片容量等于切片长度
	d := make([]int, 5, 10)
	fmt.Printf("Type:%T, Value:%#v, len=%v, cap=%v\n", d, d, len(d), cap(d))

	// 通过len()函数获取切片的长度
	fmt.Println(len(d))
	// 通过len()函数获取切片的容量
	fmt.Println(cap(d))
}

func fourth() {
	var a []int     // 声明int类型切片
	var b = []int{} // 声明并且初始化
	c := make([]int, 0)
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("a == nil")
	}
	if b == nil {
		fmt.Println("b == nil")
	}
	fmt.Println(b, len(b), cap(b))
	if c == nil {
		fmt.Println("c == nil")
	}
	fmt.Println(c, len(c), cap(c))
}

func fifth() {
	// 	切片的赋值拷贝
	a := make([]int, 3) // [0 0 0]
	b := a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}

func six() {
	//切片的遍历
	a := []int{1, 2, 3, 4, 5}
	// 根据索引来遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	// for range 遍历
	for index, value := range a {
		fmt.Println(index, ":", value)
	}
}

func main() {
	frist()
	second()
	third()
	fourth()
	fifth()
	six()
}
