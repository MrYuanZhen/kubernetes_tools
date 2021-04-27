package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
	"time"
)

func kubeCfg() *kubernetes.Clientset {
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join("/etc","kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	//// uses the current context in kubeconfig
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err.Error())
	//}
	//// creates the clientset
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

//var rise , confdArg ="",""

func main() {

	// get args
	var rise = flag.String("rise", "30", "执行间隔时间")
	var confdArg = flag.String("confd-arg", "-log-level=debug -onetime -backend env", "confd执行参数")
	flag.Parse()

	// read kubecfg
	clientset := kubeCfg()

	// Time Ticker
	t, _ := time.ParseDuration(*rise + "s")
	timeTicker := time.NewTicker(t)

	i := 0
	for {
		// test
		//if i > 5 {
		//	break
		//}

		// get node ip
		ipArr := getNodeIP(clientset)
		ipStr := strings.Join(ipArr, ",")

		// log
		now := time.Now()
		fmt.Println(now, "当前controller节点IP地址为："+ipStr)

		//set env
		setEnv(ipStr)

		// operate confd
		operateConfd(ipStr, *confdArg)

		i++
		<-timeTicker.C
	}
	// 清理计时器
	timeTicker.Stop()

}
