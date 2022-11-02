package main

import "fmt"

type Age uint

// ==========
// 結構體
type address struct {
	country string
	city string
}

type person struct {
	name string
	age uint
}

type personWithAddress struct {
	name string
	age uint
	addr address
}

type dog struct{
	name string
	color string
}

// ==========
// 接口
type Stringer interface{
	String() string
}

// method
func (p person) String() string{
	return fmt.Sprintf("the name is %s, age is %d",p.name,p.age)
}

func printString(s fmt.Stringer){
	fmt.Println(s.String())
}

// 接口2
type Animal interface{
	Talk() string
}

func (p person) Talk() string{
	return fmt.Sprintf("Hello, my name is %s, age is %d", p.name, p.age)
}

func (d dog) Talk() string{
	return fmt.Sprintf("Hello, my name is %s, color is %s", d.name, d.color)
}

func printHello(s Animal){
	fmt.Println(s.Talk())
}

// ==========
// 工廠函數
func NewPerson(name string) *person{
	return &person{name:name}
}

func New(text string) error{
	return &errorString{text}
}

type errorString struct{
	s string
}

func (e *errorString) Error() string{
	return e.s
}

// ==========
// 繼承組合
type personWithAddressExtends struct { //外部類型
	name string
	age uint
	address //內部類型
}

func main(){
	var p person
	p.name = "ita"
	fmt.Println(p.name)

	p2 := person{"ita", 20} // 順序需一致
	fmt.Println(p2.name)

	p3 := person{age:20, name:"ita"} // 順序不需一致
	fmt.Println(p3.name)

	d := dog{name:"小黑", color:"black"}

	p4:= personWithAddress{
		age:20,
		name:"名稱",
		addr:address{
			country:"台灣",
			city:"台中",
		},
	}
	fmt.Println(p4.addr.city)

	printString(p2)
	printHello(p2)
	printHello(d)

	var p5 = NewPerson("ita工廠")
	fmt.Println(p5.name)

	p6:= personWithAddressExtends{
		age:20,
		name:"名稱",
		address:address{
			country:"台灣",
			city:"台中-組合",
		},
	}
	fmt.Println(p6.city)


	// ==========
	// 類型斷言
	// - 用來檢查接口類型變量
	// - 檢查的同時順便轉換
	// value, ok := p2.(*person)
	// fmt.Println(value, ok)

}

