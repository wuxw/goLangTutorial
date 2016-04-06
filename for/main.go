package main

import (
	"fmt"
	"container/ring"
	//"time"
	//"unsafe"
	"regexp"
	"strconv"
)


//container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
//通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。
//目标字符串
//searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" 变成 John: 5156.68 William: 9134.46 Steve: 11264.36

func main()  {
	iArr := []int{101,102,103}
	xr := ring.New( len(iArr) ) //添加多大的数组
	for i := 0 ; i< xr.Len() ; i++  {
		xr.Value = iArr[i] //
		xr = xr.Next()
	}
	xr.Do(func( x interface{}) {
		fmt.Println(x)
	})

	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"
	if ok, _ := regexp.Match(pat,[]byte(searchIn) ); ok {
		fmt.Println("Match Found!")
	}
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn,"##.#")
	fmt.Println(str)
	f := func(s string) string{
		v, _ := strconv.ParseFloat(s, 32) //字符串转float32
		return strconv.FormatFloat(v * 2, 'f', 2, 32)//f代表是小数点的2个数字，类型float32,float在前面
	}
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
	/*for _ = range time.Tick(time.Second * 1) {
		xr = xr.Next()
		fmt.Println(xr.Value)
	}*/

}
