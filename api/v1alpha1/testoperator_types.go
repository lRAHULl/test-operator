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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TestOperatorSpec defines the desired state of TestOperator
type TestOperatorSpec struct {
	Image string `json:"image"`
	Resources corev1.ResourceRequirements `json:"resources"`
}

// TestOperatorStatus defines the observed state of TestOperator
type TestOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TestOperator is the Schema for the testoperators API
type TestOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestOperatorSpec   `json:"spec,omitempty"`
	Status TestOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TestOperatorList contains a list of TestOperator
type TestOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestOperator{}, &TestOperatorList{})
}
