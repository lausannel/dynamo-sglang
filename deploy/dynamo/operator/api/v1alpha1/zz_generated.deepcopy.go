//go:build !ignore_autogenerated

/*
 * SPDX-FileCopyrightText: Copyright (c) 2022 Atalaya Tech. Inc
 * SPDX-FileCopyrightText: Copyright (c) 2024-2025 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * Modifications Copyright (c) 2025 NVIDIA CORPORATION & AFFILIATES
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/ai-dynamo/dynamo/deploy/dynamo/operator/api/dynamo/common"
	"github.com/ai-dynamo/dynamo/deploy/dynamo/operator/api/dynamo/schemas"
	"k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Autoscaling) DeepCopyInto(out *Autoscaling) {
	*out = *in
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(v2.HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Autoscaling.
func (in *Autoscaling) DeepCopy() *Autoscaling {
	if in == nil {
		return nil
	}
	out := new(Autoscaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseCRD) DeepCopyInto(out *BaseCRD) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(BaseStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseCRD.
func (in *BaseCRD) DeepCopy() *BaseCRD {
	if in == nil {
		return nil
	}
	out := new(BaseCRD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseStatus) DeepCopyInto(out *BaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseStatus.
func (in *BaseStatus) DeepCopy() *BaseStatus {
	if in == nil {
		return nil
	}
	out := new(BaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoContext) DeepCopyInto(out *BentoContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoContext.
func (in *BentoContext) DeepCopy() *BentoContext {
	if in == nil {
		return nil
	}
	out := new(BentoContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoModel) DeepCopyInto(out *BentoModel) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoModel.
func (in *BentoModel) DeepCopy() *BentoModel {
	if in == nil {
		return nil
	}
	out := new(BentoModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDeployment) DeepCopyInto(out *DynamoDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDeployment.
func (in *DynamoDeployment) DeepCopy() *DynamoDeployment {
	if in == nil {
		return nil
	}
	out := new(DynamoDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDeploymentList) DeepCopyInto(out *DynamoDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamoDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDeploymentList.
func (in *DynamoDeploymentList) DeepCopy() *DynamoDeploymentList {
	if in == nil {
		return nil
	}
	out := new(DynamoDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDeploymentSpec) DeepCopyInto(out *DynamoDeploymentSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make(map[string]*DynamoNimDeployment, len(*in))
		for key, val := range *in {
			var outVal *DynamoNimDeployment
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(DynamoNimDeployment)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDeploymentSpec.
func (in *DynamoDeploymentSpec) DeepCopy() *DynamoDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DynamoDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoDeploymentStatus) DeepCopyInto(out *DynamoDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoDeploymentStatus.
func (in *DynamoDeploymentStatus) DeepCopy() *DynamoDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DynamoDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNim) DeepCopyInto(out *DynamoNim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNim.
func (in *DynamoNim) DeepCopy() *DynamoNim {
	if in == nil {
		return nil
	}
	out := new(DynamoNim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoNim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimDeployment) DeepCopyInto(out *DynamoNimDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimDeployment.
func (in *DynamoNimDeployment) DeepCopy() *DynamoNimDeployment {
	if in == nil {
		return nil
	}
	out := new(DynamoNimDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoNimDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimDeploymentList) DeepCopyInto(out *DynamoNimDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamoNimDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimDeploymentList.
func (in *DynamoNimDeploymentList) DeepCopy() *DynamoNimDeploymentList {
	if in == nil {
		return nil
	}
	out := new(DynamoNimDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoNimDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimDeploymentSpec) DeepCopyInto(out *DynamoNimDeploymentSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(common.Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(Autoscaling)
		(*in).DeepCopyInto(*out)
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFromSecret != nil {
		in, out := &in.EnvFromSecret, &out.EnvFromSecret
		*out = new(string)
		**out = **in
	}
	if in.PVC != nil {
		in, out := &in.PVC, &out.PVC
		*out = new(PVC)
		(*in).DeepCopyInto(*out)
	}
	if in.RunMode != nil {
		in, out := &in.RunMode, &out.RunMode
		*out = new(RunMode)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalServices != nil {
		in, out := &in.ExternalServices, &out.ExternalServices
		*out = make(map[string]ExternalService, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Ingress.DeepCopyInto(&out.Ingress)
	if in.MonitorExporter != nil {
		in, out := &in.MonitorExporter, &out.MonitorExporter
		*out = new(common.MonitorExporterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraPodMetadata != nil {
		in, out := &in.ExtraPodMetadata, &out.ExtraPodMetadata
		*out = new(common.ExtraPodMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraPodSpec != nil {
		in, out := &in.ExtraPodSpec, &out.ExtraPodSpec
		*out = new(common.ExtraPodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimDeploymentSpec.
func (in *DynamoNimDeploymentSpec) DeepCopy() *DynamoNimDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DynamoNimDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimDeploymentStatus) DeepCopyInto(out *DynamoNimDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimDeploymentStatus.
func (in *DynamoNimDeploymentStatus) DeepCopy() *DynamoNimDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DynamoNimDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimList) DeepCopyInto(out *DynamoNimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamoNim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimList.
func (in *DynamoNimList) DeepCopy() *DynamoNimList {
	if in == nil {
		return nil
	}
	out := new(DynamoNimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoNimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimRequest) DeepCopyInto(out *DynamoNimRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimRequest.
func (in *DynamoNimRequest) DeepCopy() *DynamoNimRequest {
	if in == nil {
		return nil
	}
	out := new(DynamoNimRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoNimRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimRequestList) DeepCopyInto(out *DynamoNimRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamoNimRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimRequestList.
func (in *DynamoNimRequestList) DeepCopy() *DynamoNimRequestList {
	if in == nil {
		return nil
	}
	out := new(DynamoNimRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamoNimRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimRequestSpec) DeepCopyInto(out *DynamoNimRequestSpec) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(BentoContext)
		**out = **in
	}
	if in.Models != nil {
		in, out := &in.Models, &out.Models
		*out = make([]BentoModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageBuildTimeout != nil {
		in, out := &in.ImageBuildTimeout, &out.ImageBuildTimeout
		*out = new(schemas.Duration)
		**out = **in
	}
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImageBuilderExtraPodMetadata != nil {
		in, out := &in.ImageBuilderExtraPodMetadata, &out.ImageBuilderExtraPodMetadata
		*out = new(common.ExtraPodMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageBuilderExtraPodSpec != nil {
		in, out := &in.ImageBuilderExtraPodSpec, &out.ImageBuilderExtraPodSpec
		*out = new(common.ExtraPodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageBuilderExtraContainerEnv != nil {
		in, out := &in.ImageBuilderExtraContainerEnv, &out.ImageBuilderExtraContainerEnv
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageBuilderContainerResources != nil {
		in, out := &in.ImageBuilderContainerResources, &out.ImageBuilderContainerResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.OCIRegistryInsecure != nil {
		in, out := &in.OCIRegistryInsecure, &out.OCIRegistryInsecure
		*out = new(bool)
		**out = **in
	}
	if in.DownloaderContainerEnvFrom != nil {
		in, out := &in.DownloaderContainerEnvFrom, &out.DownloaderContainerEnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimRequestSpec.
func (in *DynamoNimRequestSpec) DeepCopy() *DynamoNimRequestSpec {
	if in == nil {
		return nil
	}
	out := new(DynamoNimRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimRequestStatus) DeepCopyInto(out *DynamoNimRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimRequestStatus.
func (in *DynamoNimRequestStatus) DeepCopy() *DynamoNimRequestStatus {
	if in == nil {
		return nil
	}
	out := new(DynamoNimRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimSpec) DeepCopyInto(out *DynamoNimSpec) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(BentoContext)
		**out = **in
	}
	if in.Models != nil {
		in, out := &in.Models, &out.Models
		*out = make([]BentoModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimSpec.
func (in *DynamoNimSpec) DeepCopy() *DynamoNimSpec {
	if in == nil {
		return nil
	}
	out := new(DynamoNimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamoNimStatus) DeepCopyInto(out *DynamoNimStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamoNimStatus.
func (in *DynamoNimStatus) DeepCopy() *DynamoNimStatus {
	if in == nil {
		return nil
	}
	out := new(DynamoNimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalService) DeepCopyInto(out *ExternalService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalService.
func (in *ExternalService) DeepCopy() *ExternalService {
	if in == nil {
		return nil
	}
	out := new(ExternalService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
	*out = *in
	if in.UseVirtualService != nil {
		in, out := &in.UseVirtualService, &out.UseVirtualService
		*out = new(bool)
		**out = **in
	}
	if in.HostPrefix != nil {
		in, out := &in.HostPrefix, &out.HostPrefix
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(IngressTLSSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSpec.
func (in *IngressSpec) DeepCopy() *IngressSpec {
	if in == nil {
		return nil
	}
	out := new(IngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressTLSSpec) DeepCopyInto(out *IngressTLSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressTLSSpec.
func (in *IngressTLSSpec) DeepCopy() *IngressTLSSpec {
	if in == nil {
		return nil
	}
	out := new(IngressTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVC) DeepCopyInto(out *PVC) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	out.Size = in.Size.DeepCopy()
	if in.MountPoint != nil {
		in, out := &in.MountPoint, &out.MountPoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVC.
func (in *PVC) DeepCopy() *PVC {
	if in == nil {
		return nil
	}
	out := new(PVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunMode) DeepCopyInto(out *RunMode) {
	*out = *in
	if in.Standalone != nil {
		in, out := &in.Standalone, &out.Standalone
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunMode.
func (in *RunMode) DeepCopy() *RunMode {
	if in == nil {
		return nil
	}
	out := new(RunMode)
	in.DeepCopyInto(out)
	return out
}
