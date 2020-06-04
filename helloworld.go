package main

type Employee struct {
	name    string
	address string
	age     int16
}

func (e Employee) SayHello() {
	fmt.println("Xinchao %s", e.name)
	fmt.println("dia chi cua ban la: %s", e.address)
	fmt.println("ban %d tuoi", e.age)
}

func main() {
	employ := Employee{"Huong", "Phu Tho", 22}
	employ.SayHello()
}
