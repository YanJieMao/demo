/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	hjydevv1 "base/kubernetes-sigs/kubebuilder-demo/nginx/api/v1"
	"base/kubernetes-sigs/kubebuilder-demo/nginx/pkg/resource"

	"github.com/go-logr/logr"
	"github.com/juju/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// NginxReconciler reconciles a Nginx object
type NginxReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=hjy-dev.my.domain,resources=nginxes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hjy-dev.my.domain,resources=nginxes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hjy-dev.my.domain,resources=nginxes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Nginx object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *NginxReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("nginx", req.NamespacedName)

	// your logic here
	// Fetch the nginx instance
	instance := &hjydevv1.Nginx{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		log.Log.Error(err, "clent 不能创建")
		if errors.IsNotFound(err) {
			log.Log.Error(err, "ns not found")
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if instance.DeletionTimestamp != nil {
		log.Log.Error(err, "deletionTimeStmp")
		return reconcile.Result{}, err
	}

	//deploy := &appsv1.Deployment{}

	deploy := resource.NewDeploy(instance)
	if err := r.Client.Create(context.TODO(), deploy); err != nil {
		log.Log.Error(err, "新建deploy出错")
		return reconcile.Result{}, err
	}
	service := resource.NewService(instance)
	if err := r.Client.Create(context.TODO(), service); err != nil {
		log.Log.Error(err, "新建service出错")
		return reconcile.Result{}, err
	}

	/* if err := r.Client.Get(context.TODO(), req.NamespacedName, deploy); err != nil && errors.IsNotFound(err) {
		log.Log.Error(err, "新建deploy出错")
		// 创建关联资源
		// 1. 创建 Deploy
		deploy := resource.NewDeploy(instance)
		if err := r.Client.Create(context.TODO(), deploy); err != nil {
			log.Log.Error(err, "新建deploy出错")
			return reconcile.Result{}, err
		}
		// 2. 创建 Service
		service := resource.NewService(instance)
		if err := r.Client.Create(context.TODO(), service); err != nil {
			log.Log.Error(err, "新建service出错")
			return reconcile.Result{}, err
		}

		// 3. 关联 Annotations
		data, _ := json.Marshal(instance.Spec)
		if instance.Annotations != nil {
			instance.Annotations["spec"] = string(data)
		} else {
			instance.Annotations = map[string]string{"spec": string(data)}
		}

		if err := r.Client.Update(context.TODO(), instance); err != nil {
			log.Log.Error(err, "update出错")
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, nil
	} */

	/* oldspec := hjydevv1.NginxSpec{}

	if !reflect.DeepEqual(instance.Spec, oldspec) {
		// 更新关联资源
		newDeploy := resource.NewDeploy(instance)
		oldDeploy := &appsv1.Deployment{}
		if err := r.Client.Get(context.TODO(), req.NamespacedName, oldDeploy); err != nil {
			log.Log.Error(err, "获取oldDeploy出错")
			return reconcile.Result{}, err
		}
		oldDeploy.Spec = newDeploy.Spec
		if err := r.Client.Update(context.TODO(), oldDeploy); err != nil {
			log.Log.Error(err, "newDeploy出错")
			return reconcile.Result{}, err
		}

		newService := resource.NewService(instance)
		oldService := &corev1.Service{}
		if err := r.Client.Get(context.TODO(), req.NamespacedName, oldService); err != nil {
			log.Log.Error(err, "获取oldService出错")
			return reconcile.Result{}, err
		}
		oldService.Spec = newService.Spec
		if err := r.Client.Update(context.TODO(), oldService); err != nil {
			log.Log.Error(err, "更新service出错")
			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil
	}
	*/
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NginxReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hjydevv1.Nginx{}).
		Complete(r)
}
