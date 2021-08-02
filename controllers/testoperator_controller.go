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
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "k8s.io/api/core/v1"

	k8sv1alpha1 "test-operator/api/v1alpha1"
)

// TestOperatorReconciler reconciles a TestOperator object
type TestOperatorReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=k8s.rahulp.tech,resources=testoperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=k8s.rahulp.tech,resources=testoperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=k8s.rahulp.tech,resources=testoperators/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete

func (r *TestOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("Reconciling testoperator")

	// your logic here
	var testop k8sv1alpha1.TestOperator
	err := r.Get(ctx, req.NamespacedName, &testop)
	if err != nil {
		fmt.Println(err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	} else {
		fmt.Println("NOERROR!")
	}

	var pod corev1.Pod
	pod.Name = testop.Name
	pod.Namespace = testop.Namespace
	pod.Spec.Containers = []corev1.Container{
		{
			Name:      testop.Name,
			Image:     testop.Spec.Image,
			Resources: *testop.Spec.Resources.DeepCopy(),
		},
	}

	_, err1 := ctrl.CreateOrUpdate(ctx, r.Client, &pod, func() error {
		return ctrl.SetControllerReference(&testop, &pod, r.Scheme)
	})

	if err1 != nil {
		fmt.Println(err1.Error())
		return ctrl.Result{}, nil
	} else {
		fmt.Println("NOERROR2!")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
// This enables updation/deletion of the objects owned by operator
func (r *TestOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&k8sv1alpha1.TestOperator{}).
		Owns(&corev1.Pod{}).
		Complete(r)
}
