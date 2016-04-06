package main

import (
	"fmt"
	"gopl.io100/ch2/echo3/cl"
)

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func main() {
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
	cl.Ttest();
}