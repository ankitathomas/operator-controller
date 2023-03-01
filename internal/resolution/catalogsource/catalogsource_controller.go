/*
Copyright 2023.

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

package catalogsource

import (
	"context"

	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// CatalogSourceReconciler reconciles a Operator object
type CatalogSourceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func NewCatalogSourceReconciler(client client.Client, scheme *runtime.Scheme) *CatalogSourceReconciler {
	return &CatalogSourceReconciler{
		Client: client,
		Scheme: scheme,
	}
}

//+kubebuilder:rbac:groups=operators.coreos.com,resources=catalogsources,verbs=get;list;watch
//+kubebuilder:rbac:groups=operators.coreos.com,resources=catalogsources/status,verbs=get

func (r *CatalogSourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx).WithName("CatalogSource Reconciler")
	l.V(1).Info("starting")
	defer l.V(1).Info("ending")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CatalogSourceReconciler) SetupWithManager(mgr ctrl.Manager) error {

	if err := v1alpha1.AddToScheme(r.Scheme); err != nil {
		return nil
	}

	err := ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.CatalogSource{}).
		Complete(r)

	if err != nil {
		return err
	}
	return nil
}
