package main

import (
	"github.com/aditya37/modulation/transport"
	tv "github.com/aditya37/modulation/transport/v2"
	tv3 "github.com/aditya37/modulation/transport/v3"
)

func main() {
	transport.Print()
	tv.Transport()
	tv3.Print("s")
}
