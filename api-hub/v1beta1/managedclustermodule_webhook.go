/*
Copyright 2022.

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

package v1beta1

import (
	kmmv1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var modulelog = logf.Log.WithName("managed-cluster-module-resource")

func (mcm *ManagedClusterModule) SetupWebhookWithManager(mgr ctrl.Manager) error {
	// controller-runtime will set the path to `validate-<group>-<version>-<resource> so we
	// need to make sure it is set correctly in the +kubebuilder annotation below.
	return ctrl.NewWebhookManagedBy(mgr).
		For(mcm).
		Complete()
}

//+kubebuilder:webhook:path=/validate-hub-kmm-sigs-x-k8s-io-v1beta1-managedclustermodule,mutating=false,failurePolicy=fail,sideEffects=None,groups=hub.kmm.sigs.x-k8s.io,resources=managedclustermodules,verbs=create;update,versions=v1beta1,name=vmanagedclustermodule.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ManagedClusterModule{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (mcm *ManagedClusterModule) ValidateCreate() (admission.Warnings, error) {
	modulelog.Info("Validating ManagedClusterModule creation", "name", mcm.Name, "namespace", mcm.Namespace)

	module := &kmmv1beta1.Module{
		Spec: mcm.Spec.ModuleSpec,
	}
	return module.ValidateCreate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (mcm *ManagedClusterModule) ValidateUpdate(obj runtime.Object) (admission.Warnings, error) {
	modulelog.Info("Validating ManagedClusterModule update", "name", mcm.Name, "namespace", mcm.Namespace)

	module := &kmmv1beta1.Module{
		Spec: mcm.Spec.ModuleSpec,
	}
	return module.ValidateUpdate(obj)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (mcm *ManagedClusterModule) ValidateDelete() (admission.Warnings, error) {
	modulelog.Info("Validating ManagedClusterModule delete", "name", mcm.Name, "namespace", mcm.Namespace)

	module := &kmmv1beta1.Module{
		Spec: mcm.Spec.ModuleSpec,
	}
	return module.ValidateDelete()
}
