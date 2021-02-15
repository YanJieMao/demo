package main

import (
	"client-go-demo/common"
	"log"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

//informer读取集群资源的变化
func main() {

	clientset, err := kubernetes.NewForConfig(common.InitConfig())
	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)
	informer := sharedInformers.Core().V1().Pods().Informer()
	//读取增加 更新 删除 资源的变化
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("new pod add to store : %s", mObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			log.Printf("%s pod update to %s", oObj.GetName(), nObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("pod delete from store : %s", mObj.GetName())

		},
	})
	informer.Run(stopCh)

}
