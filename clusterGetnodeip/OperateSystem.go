package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func setEnv(ipStr string) {

	os.Setenv("ENDPOINTS", ipStr)

}

func operateConfd(ipStr string, confdArg string) {

	//get controller NoeIP env
	var MASTERIPS string
	MASTERIPS = os.Getenv("ENDPOINTS")
	fmt.Println("当前准备执行更新的master节点IP为：", MASTERIPS)

	// run confd process
	cmd := exec.Command("ENDPOINTS="+ipStr, "confd", confdArg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	fmt.Println("执行更新confd命令：", cmd)

}
