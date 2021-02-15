package main

import (
	"client-go-demo/common"
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getPodLog() string {
	var (
		clientset *kubernetes.Clientset
		//tailLines int64
		res  rest.Result
		logs []byte
		err  error
	)

	// 初始化k8s客户端
	clientset = common.InitClientset(common.InitConfig())

	// 获取POD日志
	res = clientset.CoreV1().Pods("redis").GetLogs("redis-55c57696d6-7n7wz",
		&corev1.PodLogOptions{}).
		Do(context.TODO())

	// 获取结果
	if logs, err = res.Raw(); err != nil {
		panic(err)
	}

	fmt.Println("输出log:\n", string(logs))
	str := string(logs)
	return str

}

func main() {

	getPodLog()

}
