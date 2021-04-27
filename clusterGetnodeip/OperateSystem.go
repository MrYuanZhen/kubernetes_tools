package main

import (
	"fmt"
	"os"
	"os/exec"
)

func setEnv(ipStr string) {

	os.Setenv("MASTERIP", ipStr)

}

func operateConfd(ipStr string, confdArg string) {
	var MASTERIPS string
	MASTERIPS = os.Getenv("MASTERIP")

	fmt.Println("当前准备执行更新的master节点IP为：",MASTERIPS)

	cmd := exec.Command(MASTERIPS,"confd", confdArg)
	fmt.Println("执行更新confd命令：",cmd)

}
