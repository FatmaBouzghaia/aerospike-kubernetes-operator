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

// import (
// 	"context"

// 	"github.com/go-logr/logr"
// 	"k8s.io/apimachinery/pkg/runtime"
// 	ctrl "sigs.k8s.io/controller-runtime"
// 	"sigs.k8s.io/controller-runtime/pkg/client"

// 	aerospikev1alpha1 "github.com/aerospike/aerospike-kubernetes-operator/api/v1alpha1"
// )

// // AerospikeClusterReconciler reconciles a AerospikeCluster object
// type AerospikeClusterReconciler struct {
// 	client.Client
// 	Log    logr.Logger
// 	Scheme *runtime.Scheme
// }

// // +kubebuilder:rbac:groups=aerospike.aerospike.com,resources=aerospikeclusters,verbs=get;list;watch;create;update;patch;delete
// // +kubebuilder:rbac:groups=aerospike.aerospike.com,resources=aerospikeclusters/status,verbs=get;update;patch
// // +kubebuilder:rbac:groups=aerospike.aerospike.com,resources=aerospikeclusters/finalizers,verbs=update

// // Reconcile is part of the main kubernetes reconciliation loop which aims to
// // move the current state of the cluster closer to the desired state.
// // TODO(user): Modify the Reconcile function to compare the state specified by
// // the AerospikeCluster object against the actual cluster state, and then
// // perform operations to make the cluster state reflect the state specified by
// // the user.
// //
// // For more details, check Reconcile and its Result here:
// // - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
// func (r *AerospikeClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
// 	_ = r.Log.WithValues("aerospikecluster", req.NamespacedName)

// 	// your logic here

// 	return ctrl.Result{}, nil
// }

// // SetupWithManager sets up the controller with the Manager.
// func (r *AerospikeClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
// 	return ctrl.NewControllerManagedBy(mgr).
// 		For(&aerospikev1alpha1.AerospikeCluster{}).
// 		Complete(r)
// }
