// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observatorium) DeepCopyInto(out *Observatorium) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observatorium.
func (in *Observatorium) DeepCopy() *Observatorium {
	if in == nil {
		return nil
	}
	out := new(Observatorium)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Observatorium) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumList) DeepCopyInto(out *ObservatoriumList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Observatorium, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumList.
func (in *ObservatoriumList) DeepCopy() *ObservatoriumList {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservatoriumList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumSpec) DeepCopyInto(out *ObservatoriumSpec) {
	*out = *in
	in.Thanos.DeepCopyInto(&out.Thanos)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumSpec.
func (in *ObservatoriumSpec) DeepCopy() *ObservatoriumSpec {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumStatus) DeepCopyInto(out *ObservatoriumStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumStatus.
func (in *ObservatoriumStatus) DeepCopy() *ObservatoriumStatus {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosSpec) DeepCopyInto(out *ThanosSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.QuerierReplicas != nil {
		in, out := &in.QuerierReplicas, &out.QuerierReplicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosSpec.
func (in *ThanosSpec) DeepCopy() *ThanosSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosSpec)
	in.DeepCopyInto(out)
	return out
}
