package main

import (
	"fmt"
	"os"
	"os/exec"
)

func setEnv(ipStr string) {

	os.Setenv("ENDPOINTS", ipStr)

}

func operateConfd(ipStr string, confdArg string) {

	//get controller NoeIP env
	//var MASTERIPS string
	//MASTERIPS = os.Getenv("ENDPOINTS")
	//fmt.Println("当前准备执行更新的master节点IP为：", MASTERIPS)

	// run confd process
	cmd := "ENDPOINTS=" + ipStr + " confd " + confdArg
	out, err := exec.Command("bash", "-c", cmd).Output()
	//cmd := exec.Command("ENDPOINTS="+ipStr, "confd", confdArg)
	//out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))

}
