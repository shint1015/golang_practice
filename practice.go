package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	//中身に変更を加えたい時は、アドレスを渡す必要がある。
	p.Name = "Mr." + p.Name
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("run")
	} else {
		fmt.Println("get out")
	}
}

func main() {
	//Humanというinterfaceの中に、Personを入れる
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"x"}
	//別のストラクトだとエラーが起きる
	var dog Dog = Dog{"Tom"}

	DriveCar(mike)
	DriveCar(x)
	fmt.Println(dog)
	//DriveCar(dog)

}
