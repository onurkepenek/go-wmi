package main

import (
	"fmt"
	"log"

	gowmi "github.com/limanmys/go-wmi"
)

func main() {

	res, err := gowmi.Query("10.0.2.10", "onur", "Passw0rd", "root\\cimv2", "SELECT Model FROM Win32_ComputerSystem")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range res {

		fmt.Println(item)
	}
}
