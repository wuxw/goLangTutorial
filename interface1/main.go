package main

import "fmt"

type Maps struct {
	sessionid string
	cookie string
}

type Serv  struct {
	*BaseServ
	mode string
	args []Maps
}
type BaseServ  struct{
	port int16
	host string
	debug bool
}

/**
// NewBaseService is the constructor for BaseService
func NewBaseService() (service *BaseService) {
	service = new(BaseService)
	service.Methods = NewMethods()
	service.filters = make([]Filter, 0)
	return
}
 */
func (service *Serv) Test() int {
	return  1
}

func (baseService *BaseServ) Hi() string {
	return "hi"
}

func NewBase() (service *BaseServ){
	service = new(BaseServ)
	return
}

func NewServ() (service *Serv) {
	service = new(Serv)
	service.BaseServ = NewBase()
	return
}

func main()  {
	/*abc := new(map[string]string)
	abc["a"] = "a"
	abc["v"] = "v"*/
	abc := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	if v1,ok := abc["r"]; ok{
		fmt.Println("v1",v1)
		fmt.Println("ok",ok)
	}
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a)
	fmt.Println("&a = ",&a)
	fmt.Println("*&a = ",*&a)
	fmt.Println("b = ",b)
	fmt.Println("&b = ",&b)
	fmt.Println("*&b = ",*&b)
	fmt.Println("*b = ",*b)
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
	fmt.Println("x = ",x)
	serv := NewServ()
	fmt.Printf("%v %d %s\n", serv,serv.Test(),	serv.Hi())
	LABEL100:
		for i:= 0; i<=5; i++ {
			for j:=0; j <= 5; j++ {
				if j == 4 {
					continue LABEL100;
				}
				fmt.Printf("i is: %d and j is: %d\n",i,j)
			}
		}
}