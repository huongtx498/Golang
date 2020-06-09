package helloworld

import "fmt"

type employee struct {
	name    string
	address string
	age     int16
}

func (e *employee) SayHello() {
	fmt.Printf("Xin chao %s \n", e.name)
	fmt.Printf("dia chi cua ban la: %s \n", e.address)
	fmt.Printf("ban %d tuoi \n", e.age)
}
