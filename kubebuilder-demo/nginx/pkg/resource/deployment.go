package resource

import (
	hjydevv1 "base/kubernetes-sigs/kubebuilder-demo/nginx/api/v1"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewDeploy(nginx *hjydevv1.Nginx) *appsv1.Deployment {

	fmt.Printf("创建deployment\n")
	labels := map[string]string{"nginx": nginx.Name}
	selector := &metav1.LabelSelector{MatchLabels: labels}
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "hjy-dev.my.domain/v1",
			Kind:       "Deployment",
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
		Spec: appsv1.DeploymentSpec{
			Replicas: nginx.Spec.Size,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: newContainers(nginx),
				},
			},
			Selector: selector,
		},
	}
}

func newContainers(nginx *hjydevv1.Nginx) []corev1.Container {
	fmt.Printf("创建containers\n")
	containerPorts := []corev1.ContainerPort{}
	for _, svcPort := range nginx.Spec.Ports {
		cport := corev1.ContainerPort{}
		cport.ContainerPort = svcPort.TargetPort.IntVal
		containerPorts = append(containerPorts, cport)
	}
	return []corev1.Container{
		{
			Name:            nginx.Name,
			Image:           nginx.Spec.Image,
			Resources:       nginx.Spec.Resources,
			Ports:           containerPorts,
			ImagePullPolicy: corev1.PullIfNotPresent,
			Env:             nginx.Spec.Envs,
		},
	}
}
