package main

import (
	"client-go-demo/common"
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNode() {

	//初始化客户端
	clientset := common.InitClientset(common.InitConfig())

	//获取NODE
	fmt.Println("####### 获取node ######")
	ctx := context.Background()
	nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, nds := range nodes.Items {
		fmt.Printf("NodeName: %s\n", nds.Name)
	}

	//获取 指定NODE 的详细信息
	fmt.Println("\n ####### node详细信息 ######")
	nodeName := "172.27.139.103"
	nodeRel, err := clientset.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s \n", nodeRel.Name)
	fmt.Printf("CreateTime: %s \n", nodeRel.CreationTimestamp)
	fmt.Printf("NowTime: %s \n", nodeRel.Status.Conditions[0].LastHeartbeatTime)
	fmt.Printf("kernelVersion: %s \n", nodeRel.Status.NodeInfo.KernelVersion)
	fmt.Printf("SystemOs: %s \n", nodeRel.Status.NodeInfo.OSImage)
	fmt.Printf("Cpu: %s \n", nodeRel.Status.Capacity.Cpu())
	fmt.Printf("docker: %s \n", nodeRel.Status.NodeInfo.ContainerRuntimeVersion)
	// fmt.Printf("Status: %s \n", nodeRel.Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type)
	fmt.Printf("Status: %s \n", nodeRel.Status.Conditions[len(nodeRel.Status.Conditions)-1].Type)
	fmt.Printf("Mem: %s \n", nodeRel.Status.Allocatable.Memory().String())
}

func main() {

	GetNode()
}
