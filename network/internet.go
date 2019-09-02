package main

import (
	"fmt"
	"net"
)

func main() {
	res, err := net.LookupHost("beibei.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)

	ports, err := net.LookupPort("http", "beibei.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ports)
}
