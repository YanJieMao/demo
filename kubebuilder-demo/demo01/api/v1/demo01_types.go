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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Demo01Spec defines the desired state of Demo01
type Demo01Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Demo01. Edit demo01_types.go to remove/update
	//Foo string `json:"foo,omitempty"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Demo01Status defines the observed state of Demo01
type Demo01Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"Status"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Demo01 is the Schema for the demo01s API
type Demo01 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Demo01Spec   `json:"spec,omitempty"`
	Status Demo01Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Demo01List contains a list of Demo01
type Demo01List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Demo01 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Demo01{}, &Demo01List{})
}
