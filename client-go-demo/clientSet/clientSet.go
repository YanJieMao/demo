package main

import (
	"client-go-demo/common"
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//clientset获取指定命名空间的pods
func main() {

	//初始化客户端
	clientset := common.InitClientset(common.InitConfig())

	podclient := clientset.CoreV1().Pods("redis")
	list, err := podclient.List(context.TODO(), metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf("NAMESPACE:%v \t  STATUS:%v \t NAME:%v \t \n", d.Namespace, d.Status.Phase, d.Name)

	}
}
