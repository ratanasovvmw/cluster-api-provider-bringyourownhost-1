/*
Copyright the Cluster API Provider BYOH contributors.

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

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (h *BYOHost) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(h).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1alpha3-byohost,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=byohosts,versions=v1alpha3,name=validation.byohost.infrastructure.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

var _ webhook.Validator = &BYOHost{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (h *BYOHost) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (h *BYOHost) ValidateUpdate(old runtime.Object) error {
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (h *BYOHost) ValidateDelete() error {
	return nil
}
