package main

import "fmt"

type foo struct {
	name string
	age  int
}

func (f foo) setNameByValueReceiver(n string) {
	f.name = n
	fmt.Println("nothing happened")
}

func (p *foo) setNameByPointerReceiver(n string) {
	p.name = n
}

func main() {
	f := foo{
		name: "tony",
		age:  1,
	}
	f.setNameByValueReceiver("tommy")
	fmt.Println(f)
	f.setNameByPointerReceiver("jimmy")
	fmt.Println(f)
}
