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
	"log"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testhjyv1 "base/kubernetes-sigs/kubebuilder-demo/demo01/api/v1"
)

// Demo01Reconciler reconciles a Demo01 object
type Demo01Reconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=test-hjy.demo01,resources=demo01s,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=test-hjy.demo01,resources=demo01s/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=test-hjy.demo01,resources=demo01s/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Demo01 object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *Demo01Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("demo01", req.NamespacedName)

	// your logic here
	ctx = context.Background()
	_ = r.Log.WithValues("apiexamplea", req.NamespacedName)

	// 获取当前的 CR，并打印
	obj := &testhjyv1.Demo01{}
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		log.Println(err, "Unable to fetch object")
	} else {
		log.Println("Geeting from Kubebuilder to", obj.Spec.FirstName, obj.Spec.LastName)
	}

	// 初始化 CR 的 Status 为 Running
	obj.Status.Status = "Running"
	if err := r.Status().Update(ctx, obj); err != nil {
		log.Println(err, "unable to update status")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Demo01Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testhjyv1.Demo01{}).
		Complete(r)
}
