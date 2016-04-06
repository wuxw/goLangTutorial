package main

import "fmt"

func Imod(n int,m int) bool{
	return n%m == 0
}

func switchfn(max int)  {
	for i := 1; i <= max; i++ {
		switch {
		case Imod(i,3) :
			fmt.Println("Fizz")
		case Imod(i,5):
			fmt.Println("Buzz")
		case Imod(i,3) && Imod(i,5):
			fmt.Println("FizzBuzz")
		default:
			fmt.Println("%d\n",i)
		}
	}

}

func chars1(){
	for i := 0 ;i <= 25 ; i++  {
		for j := 1; j <= i ; j++  {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}
}

func chars2()  {
	gstrs := "GGGGGGGGGGGGGGGGGGGGGGGGG"
	for i := 0 ;i <= 25 ; i++  {
		fmt.Println( gstrs[25-i:len(gstrs)] )
	}
}

func gFn()  {
	ix := 0;
	COUNTLABEL:
	ix++;
	if 15 == ix {
		fmt.Printf("counter2: %d\n",ix)
		return //return必须
	}
	goto COUNTLABEL
}