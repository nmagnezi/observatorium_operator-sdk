package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ObservatoriumSpec defines the desired state of Observatorium
type ObservatoriumSpec struct {
	// Thanos Spec
	Thanos ThanosSpec `json:"thanos"`
}

type ThanosSpec struct {
	// Thanos Image name
	Image string `json:"image"`
	// Version describes the version of Thanos to use.
	Version *string `json:"version,omitempty"`
	// Tag of Thanos sidecar container image to be deployed. Defaults to the value of `version`.
	// Version is ignored if Tag is set.
	Tag string `json:"tag"`

	// Number of instances to deploy for a Thanos querier.
	QuerierReplicas *int32 `json:"querierreplicas,omitempty"`
}

// ObservatoriumStatus defines the observed state of Observatorium
type ObservatoriumStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Observatorium is the Schema for the observatoria API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=observatoria,scope=Namespaced
type Observatorium struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObservatoriumSpec   `json:"spec,omitempty"`
	Status ObservatoriumStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ObservatoriumList contains a list of Observatorium
type ObservatoriumList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Observatorium `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Observatorium{}, &ObservatoriumList{})
}
