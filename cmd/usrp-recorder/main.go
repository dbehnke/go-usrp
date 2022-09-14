package main

import (
	"github.com/dbehnke/go-usrp.git"
)

func main() {
	c := usrp.Config{
		Group:  "testing",
		RXPort: ":57771",
		TXPort: ":57772",
	}
	usrp.RxUSRP(&c)
}
