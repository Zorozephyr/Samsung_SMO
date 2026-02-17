/*
Copyright 2026.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default="postgres:14"
	Image string `json:"image"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	// +kubebuilder:default=1
	Replicas  *int32                      `json:"replicas"`
	Storage   StorageSpec                 `json:"storage"`
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// +kubebuilder:validation:Required
	DatabaseName string `json:"databaseName"`
	// +kubebuilder:validation:Required
	Username string `json:"username"`
}

type StorageSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^[0-9]+(Mi|Gi)$"
	Size             string `json:"size"`
	StorageClassName string `json:"storageClassName,omitempty"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	// +kubebuilder:validation:Enum=Pending;Creating;Ready;Failed
	Phase      string `json:"phase,omitempty"`
	Ready      bool   `json:"ready,omitempty"`
	Endpoint   string `json:"endpoint,omitempty"`
	SecretName string `json:"secretName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.replicas"
// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Database is the Schema for the databases API
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec,omitempty"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Database
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []Database `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
