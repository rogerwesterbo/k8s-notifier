/*
Copyright 2024.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	taskv1 "github.com/rogerwesterbo/k8s-notifier/api/v1"
)

// NotifyReconciler reconciles a Notify object
type NotifyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=task.hemmelig.io,resources=notifies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=task.hemmelig.io,resources=notifies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=task.hemmelig.io,resources=notifies/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Notify object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *NotifyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	notify := &taskv1.Notify{}
	err := r.Get(ctx, req.NamespacedName, notify)
	if err != nil {
		if errors.IsNotFound(err) {
			rlog.Info("Notify resource not found.")
			return ctrl.Result{}, nil
		}

		rlog.Error("Error getting notify", err)
		return ctrl.Result{}, err
	}

	rlog.Info("Reconciling Notify", rlog.String("name", notify.Name), rlog.String("namespace", notify.Namespace), rlog.String("type.apiversion", notify.Spec.Type.APIVersion), rlog.String("type.kind", notify.Spec.Type.Kind))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NotifyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&taskv1.Notify{}).
		Complete(r)
}
