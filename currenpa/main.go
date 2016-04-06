package main


import (
	"fmt"
	//"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
func modifySlice(data []int)  {
	fmt.Printf("modifySlice: %p\n",&data)
	data = nil
}
func modifySliceData(data []int)  {
	//append(data,6666)
	fmt.Printf("modifySliceData: %p\n",&data)
	data[0] = 100
	//append(data,6)
	//return data
}

type RefIntArray2 *[2]int
func NewRefIntArray2() RefIntArray2 {
	return RefIntArray2(new([2]int))
}

func modifyRefArr2(arr RefIntArray2) {
	arr[0] = 980
}


func modify(a1 *int)  {
	fmt.Printf("modify: %p\n",&a1)
	v := 100
	a1 = &v
}

func sliceAppend(s []int,n int)  []int {
	newSlice := make([]int,n)
	copy(s, newSlice)
	return  s
}

func insertSliceIndex(sTo []int,sFrom []int,index int)  []int {
	//按index拆分sTo
	sStart := sTo[0:index]
	sEnd := make([]int,len(sStart)+len(sFrom))
	for i,v := range sTo{
		if i >= index {
			sEnd = append(sEnd,v)
		}
	}
	fmt.Printf("len1: %d cap: %d %v\n",len(sEnd),cap(sEnd),sEnd)
	for _,v := range sFrom{
		sStart = append(sStart,v)
	}
	fmt.Printf("len2: %d cap: %d %v\n",len(sStart),cap(sStart),sStart)

	for  _,v1 := range sEnd  {
		fmt.Println(v1)
		sStart = append(sEnd,v1)
	}
	fmt.Printf("len3: %d cap: %d %v\n",len(sStart),cap(sStart),sStart)
	return sStart
}

func print(msg string,ss []string){
	fmt.Printf("[ %20s ]\t:\tlength:%v\taddr:%p\tisnil:%v\tcontent:%v\n",msg,len(ss),ss, ss==nil,ss)
}


func RemoveSlice(slice []interface{}, start, end int) []interface{} {
	result := make([]interface{}, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}
//--------------------另一种更为简便的写法-----------------------
func RemoveSlice2(slice []interface{}, start, end int) []interface{} {
	return append(slice[:start], slice[end:]...)
}


func InsertSlice(slice, insertion []interface{}, index int) []interface{} {
	result := make([]interface{}, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func RemoveSliceString(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}
func InsertSliceString(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func InsertSliceString2(strs []string, nstrs []string,index int)  []string{
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear:= append([]string{},strs[index:]...)
	for i,v := range  nstrs{
		strs = append(strs[0:index+i],v)
	}
	strs = append(strs,rear...)
	return strs
}

func  RemoveStringSlice(strs []string,iStart int,iEnd int) []string {
	if iStart > iEnd {
		return strs
	}
	if iStart > len(strs) || iEnd > len(strs){
		return strs
	}
	//注意：保存后部剩余元素，必须新建一个临时切片
	//rear := append([]string{},strs[iEnd+1:]...)
	strs = append(strs[:iStart],strs[iEnd+1:]...)
	/*for i,v := range  rear{
		strs = append(strs[:iStart+i],v)
	}*/
	//strs = append(strs[:iStart],strs[iStart+1:iEnd])
	return strs
}

func explode( dist string,div string)  (string,string){
	index := 0
	byteDist := []byte(dist)
	byteDiv := []byte(div)
	for i, v := range byteDist  {
		if v == byteDiv[0] {
		//if v == div {
			index = i
			break
		}
	}
	if index <= 1{
		return string(dist),""
	}
	return string(byteDist[0:index]),string(byteDist[index+1:])

}

func ReverString(dist string)  string{
	byteDist := []byte(dist)
	byteDist1 := []byte{}
	iLen := len(byteDist)
	/*fmt.Printf("%d\n",iLen)
	fmt.Printf("%s\n",dist)
	fmt.Printf("%v\n",byteDist)*/
	for i := iLen-1; i >= 0; i-- {
		//fmt.Printf("%d\n",i)
		byteDist1 = append(byteDist1,byteDist[i])
	}
	return string(byteDist1)
}
type intFnType func(int) int
func f10(n int) int{
	return n*10
}
func mapFunc( f intFnType,iArr []int) []int  {
	for i,v := range iArr {
		iArr[i] =  f(v)
	}
	return iArr
}
//var s4 []interface{}
func main() {
	var ss []string;
	str100 := "a我是Si"
	for  i,s :=  range str100 {
		fmt.Println(i,"Unicode(",s,") string=",string(s))
	}
	//s10 := []string{"abc12|test100"}
	s10,s100 := explode("abc12kza1|test100","|")
	//str[len(str)/2:] + str[:len(str)/2]
	fmt.Printf("s10: %s %d, s100: %s\n",s10,len(s10),s100)
	fmt.Printf("%v\n",s10[len(s10)/2:] + s10[:len(s10)/2])
	fmt.Printf("%s\n",ReverString(s100))
	fmt.Printf("%v\n",mapFunc(f10,[]int{80,800,8000}))
	s1 := []string{"a1","b1","c1","d1","e1","f1","g1","h1"}
	s2 :=[]string{"bmw100","passte100","df100"}
	//s3 := InsertStringSlice(s1,s2,7)
	s3 := InsertSliceString(s1,s2,7)
	fmt.Printf("len: %d cap: %d\n slice1:%v\n",len(s3),cap(s3),s3)
	//s4 := s3
	s4 := append([]string{},s3[0:]...)
	s4 = RemoveSliceString(s4,2,5)
	fmt.Printf("len: %d cap: %d\n slice1:%v\n",len(s4),cap(s4),s4)
	fmt.Printf("len: %d cap: %d\n slice1:%v\n",len(s3),cap(s3),s3)
	return
	fmt.Printf("[ local print ]\t:\t length:%v\taddr:%p\tisnil:%v\n",len(ss),ss, ss==nil)
	print("func print",ss)
	//切片尾部追加元素append elemnt
	for i:=0;i<10;i++{
		ss=append(ss,fmt.Sprintf("s%d",i));
	}
	fmt.Printf("[ local print ]\t:\tlength:%v\taddr:%p\tisnil:%v\n",len(ss),ss, ss==nil)
	print("after append",ss)
	//删除切片元素remove element at index
	index:=5;
	ss=append(ss[:index],ss[index+1:]...)
	print("after delete",ss)
	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear:=append([]string{},ss[index:]...)
	iStrs := []string{"inserted","hiwolrd"}
	for i,v := range iStrs {
		ss = append(ss[0:index+i],v)
	}
	//ss=append(ss[0:index],iStrs)
	ss=append(ss,rear...)
	print("after insert",ss)
	slice1 := make([]int, 0, 10)
	for i := 0; i< cap(slice1); i++ {
		slice1 = slice1[0:i+1]
		slice1[i] = i
		fmt.Printf("len: %d cap: %d\n slice1:%v\n",len(slice1),cap(slice1),slice1)
	}
	fmt.Printf("slice1:%v\n",slice1)
	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Printf("ar:%v\n",ar)

	var aslice = ar[4:7] // reference to subarray {5,6} - len(a) is 2 and cap(a) is 5
	//var aslice1 = slice1[5:9]
	aslice1 := slice1[5:9]
	//cap(aslice) cap(oldSlice) -起始位置
	fmt.Printf("len: %d cap: %d\n",len(aslice),cap(aslice))
	fmt.Printf("len: %d cap: %d\n",len(aslice1),cap(aslice1))
	return
	//runtime.GOMAXPROCS(1) //单进程的话，会被for占满
	//go say("world") //开一个新的Goroutines执行
	//for {
	//}
	a := []int{1,2,3}
	fmt.Printf("modifySlice: %p\n",&a)
	fmt.Println(a)
	modifySlice(a)
	fmt.Println(a)

	b := []int{66,88,99}
	fmt.Printf("modifySliceData: %p\n",&b)
	fmt.Println(b)
	modifySliceData(b)
	fmt.Println(b)

	a1 := new(int)
	fmt.Printf("modify: %p\n",&a1)
	fmt.Println(a1)
	modify(a1)
	fmt.Println(a1)

	refArr2 := NewRefIntArray2()
	fmt.Println(refArr2)
	modifyRefArr2(refArr2)
	fmt.Println(refArr2)

}

/*
import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 24; i++ { //为了观察，跑多些
		//runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		fmt.Printf("%d\n", i)
	}
	quit <- 0
}

func main() {
	//runtime.GOMAXPROCS(2) // 最多使用2个核
	//runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("NumCPU: %d\n",runtime.NumCPU())
	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<- quit
	}

}
*/

/*
import "fmt"
import "time"

var quit chan int


//从上面的两个例子执行后的表现来看，多个goroutine跑loop函数会挨个goroutine去进行，而sleep则是一起执行的
//认地， Go所有的goroutines只能在一个线程里跑
//也就是说， 以上两个代码都不是并行的，但是都是是并发的。

//如果当前goroutine不发生阻塞，它是不会让出CPU给其他goroutine的, 所以例子一中的输出会是一个一个goroutine进行的，
//而sleep函数则阻塞掉了 当前goroutine, 当前goroutine主动让其他goroutine执行, 所以形成了逻辑上的并行, 也就是并发。

func foo(id int) {
	fmt.Println(id)
	time.Sleep(time.Second) // 停顿一秒
	quit <- 0 // 发消息：我执行完啦！
}


func main() {
	count := 50
	quit = make(chan int, count) // 缓冲50个数据

	for i := 0; i < count; i++ { //开50个goroutine
		go foo(i)
	}

	for i :=0 ; i < count; i++ { // 等待所有完成消息发送完毕。
		<- quit  //总计用时接近一秒。 貌似并行了，输出的是乱序
	}
}


*/



/*import "fmt"

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
	fmt.Printf("%d ", i)
	}
	quit <- 0
}


func main() {
	// 开两个goroutine跑函数loop, loop函数负责打印10个数
	go loop()
	go loop()

	for i := 0; i < 2; i++ {
	   <- quit //0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	}
	//以前我们用线程去做类似任务的时候，系统的线程会抢占式地输出， 表现出来的是乱序地输出。而goroutine为什么是这样输出的呢？
}*/
