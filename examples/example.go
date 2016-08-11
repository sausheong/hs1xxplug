package main

import (
	"fmt"
	"github.com/sausheong/hs1xxplug"
)

func main() {
	plug := hs1xxplug.Hs1xxPlug{IPAddress: "192.168.0.196"}
	results, err := plug.MeterInfo()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(results)

}
