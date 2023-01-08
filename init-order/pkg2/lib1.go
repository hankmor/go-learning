package pkg2

import "fmt"

var (
	_  = constInitCheck()
	v1 = varInitCheck("v1", 1)
	v2 = varInitCheck("v2", 2)
)

const (
	c1 = 1
	c2 = 2
)

func init() {
	fmt.Println("pkg2: invoke init method")
}

func constInitCheck() int {
	fmt.Println("pkg2: invoke constInitCheck...")
	if c1 != 1 {
		fmt.Println("pkg2: const c1 init")
	}
	if c2 != 2 {
		fmt.Println("pkg2: const c2 init")
	}
	return 1
}

func varInitCheck(name string, v int) int {
	fmt.Printf("pkg2: var %s init with value %d \n", name, v)
	return v
}

func init() {
	fmt.Println("pkg2: invoke init method 2")
}
