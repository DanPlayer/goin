package impl

import "fmt"

type HelloImpl struct{}

func (u *HelloImpl) Hello() {
	fmt.Println("hello world")
	return
}