//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSource) DeepCopyInto(out *CatalogSource) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSource.
func (in *CatalogSource) DeepCopy() *CatalogSource {
	if in == nil {
		return nil
	}
	out := new(CatalogSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCatalog) DeepCopyInto(out *ClusterCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCatalog.
func (in *ClusterCatalog) DeepCopy() *ClusterCatalog {
	if in == nil {
		return nil
	}
	out := new(ClusterCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterCatalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCatalogList) DeepCopyInto(out *ClusterCatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterCatalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCatalogList.
func (in *ClusterCatalogList) DeepCopy() *ClusterCatalogList {
	if in == nil {
		return nil
	}
	out := new(ClusterCatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterCatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCatalogSpec) DeepCopyInto(out *ClusterCatalogSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCatalogSpec.
func (in *ClusterCatalogSpec) DeepCopy() *ClusterCatalogSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterCatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCatalogStatus) DeepCopyInto(out *ClusterCatalogStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResolvedSource != nil {
		in, out := &in.ResolvedSource, &out.ResolvedSource
		*out = new(ResolvedCatalogSource)
		(*in).DeepCopyInto(*out)
	}
	if in.URLs != nil {
		in, out := &in.URLs, &out.URLs
		*out = new(ClusterCatalogURLs)
		**out = **in
	}
	if in.LastUnpacked != nil {
		in, out := &in.LastUnpacked, &out.LastUnpacked
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCatalogStatus.
func (in *ClusterCatalogStatus) DeepCopy() *ClusterCatalogStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterCatalogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCatalogURLs) DeepCopyInto(out *ClusterCatalogURLs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCatalogURLs.
func (in *ClusterCatalogURLs) DeepCopy() *ClusterCatalogURLs {
	if in == nil {
		return nil
	}
	out := new(ClusterCatalogURLs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSource) DeepCopyInto(out *ImageSource) {
	*out = *in
	if in.PollIntervalMinutes != nil {
		in, out := &in.PollIntervalMinutes, &out.PollIntervalMinutes
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSource.
func (in *ImageSource) DeepCopy() *ImageSource {
	if in == nil {
		return nil
	}
	out := new(ImageSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedCatalogSource) DeepCopyInto(out *ResolvedCatalogSource) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ResolvedImageSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedCatalogSource.
func (in *ResolvedCatalogSource) DeepCopy() *ResolvedCatalogSource {
	if in == nil {
		return nil
	}
	out := new(ResolvedCatalogSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedImageSource) DeepCopyInto(out *ResolvedImageSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedImageSource.
func (in *ResolvedImageSource) DeepCopy() *ResolvedImageSource {
	if in == nil {
		return nil
	}
	out := new(ResolvedImageSource)
	in.DeepCopyInto(out)
	return out
}
