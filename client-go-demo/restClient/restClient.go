package main

import (
	"client-go-demo/common"
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/deprecated/scheme"
	"k8s.io/client-go/rest"
)

//restClient查找指定命名空间中的所有pods
func main() {

	//初始化配置
	config := common.InitConfig()

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)

	if err != nil {
		panic(err)
	}

	resault := &corev1.PodList{}

	err = restClient.Get().
		Namespace("redis").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500},
			scheme.ParameterCodec).
		Do(context.TODO()).
		Into(resault)
	if err != nil {
		panic(err)
	}

	for _, d := range resault.Items {
		fmt.Printf("NAMESPACE:%v \t  STATUS:%v \t NAME:%v \t \n", d.Namespace, d.Status.Phase, d.Name)

	}

}
