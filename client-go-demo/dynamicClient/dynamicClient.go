package main

import (
	"client-go-demo/common"
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

//dynamicClient获取指定命名空间的pods
func main() {

	dynamicClient, err := dynamic.NewForConfig(common.InitConfig())

	if err != nil {
		panic(err)
	}

	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	unstructObj, err := dynamicClient.Resource(gvr).Namespace("redis").
		List(context.TODO(), metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}

	podList := &corev1.PodList{}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podList)

	for _, d := range podList.Items {
		fmt.Printf("NAMESPACE:%v \t  STATUS:%v \t NAME:%v \t \n", d.Namespace, d.Status.Phase, d.Name)

	}
}
