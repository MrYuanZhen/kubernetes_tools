package main

import (
	"os"
)

func setEnv(ipStr string) {

	os.Setenv("MASTERIP", ipStr)

}

func operateConfd(ipStr string, confdArg string) {

}
