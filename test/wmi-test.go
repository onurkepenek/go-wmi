package main

import (
	"encoding/json"
	"fmt"
	"log"

	gowmi "github.com/limanmys/go-wmi"
)

func main() {

	res, err := gowmi.Query("10.0.2.10", "onur", "Passw0rd", "root\\cimv2", "SELECT * FROM Win32_ComputerSystem")
	if err != nil {
		log.Fatal(err.Error())
	}

	bt, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(bt))
}
