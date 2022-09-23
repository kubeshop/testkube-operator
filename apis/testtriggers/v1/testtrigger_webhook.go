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

package v1

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// +k8s:deepcopy-gen=false
type TestTriggerValidator interface {
	ValidateCreate(context.Context, *TestTrigger) error
	ValidateUpdate(ctx context.Context, old runtime.Object, new *TestTrigger) error
	ValidateDelete(context.Context, *TestTrigger) error
}

var ctx = context.Background()
var vldtr TestTriggerValidator

func (in *TestTrigger) ValidateCreate() error {
	return vldtr.ValidateCreate(ctx, in)
}

func (in *TestTrigger) ValidateUpdate(old runtime.Object) error {
	return vldtr.ValidateUpdate(ctx, old, in)
}

func (in *TestTrigger) ValidateDelete() error {
	return vldtr.ValidateDelete(ctx, in)
}

var _ webhook.Validator = &TestTrigger{}

func (in *TestTrigger) SetupWebhookWithManager(mgr ctrl.Manager, validator TestTriggerValidator) error {
	vldtr = validator
	return ctrl.NewWebhookManagedBy(mgr).
		For(in).
		Complete()
}
