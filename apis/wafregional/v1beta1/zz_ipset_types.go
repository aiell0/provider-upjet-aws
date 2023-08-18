/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IPSetDescriptorInitParameters struct {

	// The string like IPV4 or IPV6.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The CIDR notation.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type IPSetDescriptorObservation struct {

	// The string like IPV4 or IPV6.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The CIDR notation.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type IPSetDescriptorParameters struct {

	// The string like IPV4 or IPV6.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// The CIDR notation.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type IPSetInitParameters struct {

	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IPSetDescriptor []IPSetDescriptorInitParameters `json:"ipSetDescriptor,omitempty" tf:"ip_set_descriptor,omitempty"`

	// The name or description of the IPSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IPSetObservation struct {

	// The ARN of the WAF IPSet.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the WAF IPSet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	IPSetDescriptor []IPSetDescriptorObservation `json:"ipSetDescriptor,omitempty" tf:"ip_set_descriptor,omitempty"`

	// The name or description of the IPSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IPSetParameters struct {

	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR notation) from which web requests originate.
	// +kubebuilder:validation:Optional
	IPSetDescriptor []IPSetDescriptorParameters `json:"ipSetDescriptor,omitempty" tf:"ip_set_descriptor,omitempty"`

	// The name or description of the IPSet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// IPSetSpec defines the desired state of IPSet
type IPSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPSetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IPSetInitParameters `json:"initProvider,omitempty"`
}

// IPSetStatus defines the observed state of IPSet.
type IPSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPSet is the Schema for the IPSets API. Provides a AWS WAF Regional IPSet resource for use with ALB.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IPSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   IPSetSpec   `json:"spec"`
	Status IPSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPSetList contains a list of IPSets
type IPSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPSet `json:"items"`
}

// Repository type metadata.
var (
	IPSet_Kind             = "IPSet"
	IPSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPSet_Kind}.String()
	IPSet_KindAPIVersion   = IPSet_Kind + "." + CRDGroupVersion.String()
	IPSet_GroupVersionKind = CRDGroupVersion.WithKind(IPSet_Kind)
)

func init() {
	SchemeBuilder.Register(&IPSet{}, &IPSetList{})
}
