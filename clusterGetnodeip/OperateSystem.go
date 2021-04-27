package main

import (
	"fmt"
	"os"
)

func setEnv(ipStr string) {

	os.Setenv("MASTERIP", ipStr)
	for index, key := range os.Environ() {
		fmt.Println("index = ", index, " key = ", key)
	}
}

func operateConfd(ipStr string, confdArg string) {

}
