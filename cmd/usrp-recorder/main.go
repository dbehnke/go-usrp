package main

import (
	"github.com/dbehnke/go-usrp.git"
)

func main() {
	c := usrp.Config{
		RXPort: ":57771",
		TXPort: ":57772",
	}
	usrp.RxUSRP(&c)
}
