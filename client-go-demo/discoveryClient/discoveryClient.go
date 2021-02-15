package main

import (
	"client-go-demo/common"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

//服务发现
func main() {

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(common.InitConfig())
	if err != nil {
		panic(err)
	}

	_, APIResourceList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}

	for _, list := range APIResourceList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err)
		}

		for _, resource := range list.APIResources {
			fmt.Printf("NAME:%v \t  GROUP:%v \t VERSION:%v \t \n", resource.Name, gv.Group, gv.Version)
		}

	}

}
