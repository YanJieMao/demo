package resource

import (
	hjydevv1 "base/kubernetes-sigs/kubebuilder-demo/nginx/api/v1"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewService(nginx *hjydevv1.Nginx) *corev1.Service {
	fmt.Printf("创建service\n")
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "hjy-dev.my.domain/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      nginx.Name,
			Namespace: nginx.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(nginx, schema.GroupVersionKind{
					Group:   hjydevv1.GroupVersion.Group,
					Version: hjydevv1.GroupVersion.Version,
					Kind:    "Nginx",
				}),
			},
		},
		Spec: corev1.ServiceSpec{
			Type:  corev1.ServiceTypeNodePort,
			Ports: nginx.Spec.Ports,
			Selector: map[string]string{
				"nginx": nginx.Name,
			},
		},
	}
}
