package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	gowmi "github.com/limanmys/go-wmi"
)

func main() {
	//go run wmi-test.go host user password namespace query
	//go run wmi-test.go 10.0.2.0 user pass "root\\cimv2" "SELECT * FROM Win32_ComputerSystem"
	res, err := gowmi.Query(os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	if err != nil {
		log.Fatal(err.Error())
	}

	bt, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(bt))
}
