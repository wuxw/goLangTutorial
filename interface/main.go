package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Shouter interface {
	Area() float64
	Act() string
}

type Square struct  {
	side float64
}
type Rectangle struct  {
	length,width float32
}

func (sq *Square) Area() float64{
	return sq.side*sq.side
}

func (sq *Square) Act() string {
	return "Act"
}

func (r Rectangle)Area() float32 {
	return r.length*r.width
}

//接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）
/*func (sq *Square) Area() float32  {
	return sq.side * sq.side
}*/

func main(){
	/*sq1 :=  new(Square)
	sq1.side = 5
	fmt.Printf("The square has area:\n %f", sq1.side);
	fmt.Printf("The square has area: %f\n", sq1.Area())*/
	r := Rectangle{5,3}
	q := &Square{8}
	//(*q).Act()
	shapes := []Shaper{r}
	shouts := []Shouter{q}

	for n,_ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
	for n,_ := range shouts {
		fmt.Println("Shape details: ", shouts[n])
		fmt.Println("Area of this shape is: ", shouts[n].Area())
		fmt.Println("Area of this shape is: ", shouts[n].Act())
	}
}
