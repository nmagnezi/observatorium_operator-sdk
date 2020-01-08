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

type ThanosReceiveController struct {
	// Thanos receive controller Image name
	Image *string `json:"image"`
	// Tag describes the tag of Thanos receive controller to use.
	Tag *string `json:"tag,omitempty"`
}

type ThanosSpec struct {
	// Thanos Image name
	Image *string `json:"image"`
	// Tag of Thanos sidecar container image to be deployed.
	Tag *string `json:"tag"`

	ThanosReceiveControllerSpec ThanosReceiveController `json:"thanosReceiveControllerSpec"` 
	// Number of instances to deploy for a Thanos querier.
	QuerierReplicas *int32 `json:"querierReplicas,omitempty"`
	// Number of instances to deploy for a Thanos Store.
	StoreReplicas *int32 `json:"storeReplicas,omitempty"`
	// Number of instances to deploy for a Thanos Compactor.
	CompactorReplicas *int32 `json:"compactorReplicas,omitempty"`
	// Number of instances to deploy for a Thanos Receive.
	ReceiveReplicas *int32 `json:"receiveReplicas,omitempty"`
    // Object Store Config Secret for Thanos 
    ObjectStoreConfigSecret *string `json:"objectStoreConfigSecret"`
    // TODO: AWS secrets?
	// TODO: Resource limits for: QUERIER, STORE, RECEIVE, COMPACTOR
	// TODO: handle with THANOS_QUERIER_SVC_URL
	// TODO: Do we need a THANOS_RULER?
	// TODO: JAEGER
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
