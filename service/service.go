package service

import (
	"goin/service/hello"
	helloImpl "goin/service/hello/impl"
)

var (
	HelloService hello.Service = new(helloImpl.HelloImpl)
)
