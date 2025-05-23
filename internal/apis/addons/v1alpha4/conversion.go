/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	addonsv1 "sigs.k8s.io/cluster-api/api/addons/v1beta2"

	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
)

func (src *ClusterResourceSet) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*addonsv1.ClusterResourceSet)

	if err := Convert_v1alpha4_ClusterResourceSet_To_v1beta2_ClusterResourceSet(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &addonsv1.ClusterResourceSet{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}
	dst.Status.V1Beta2 = restored.Status.V1Beta2

	return nil
}

func (dst *ClusterResourceSet) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*addonsv1.ClusterResourceSet)

	if err := Convert_v1beta2_ClusterResourceSet_To_v1alpha4_ClusterResourceSet(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

func (src *ClusterResourceSetList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*addonsv1.ClusterResourceSetList)

	return Convert_v1alpha4_ClusterResourceSetList_To_v1beta2_ClusterResourceSetList(src, dst, nil)
}

func (dst *ClusterResourceSetList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*addonsv1.ClusterResourceSetList)

	return Convert_v1beta2_ClusterResourceSetList_To_v1alpha4_ClusterResourceSetList(src, dst, nil)
}

func (src *ClusterResourceSetBinding) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*addonsv1.ClusterResourceSetBinding)

	if err := Convert_v1alpha4_ClusterResourceSetBinding_To_v1beta2_ClusterResourceSetBinding(src, dst, nil); err != nil {
		return err
	}
	// Manually restore data.
	restored := &addonsv1.ClusterResourceSetBinding{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}
	dst.Spec.ClusterName = restored.Spec.ClusterName
	return nil
}

func (dst *ClusterResourceSetBinding) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*addonsv1.ClusterResourceSetBinding)

	if err := Convert_v1beta2_ClusterResourceSetBinding_To_v1alpha4_ClusterResourceSetBinding(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

func (src *ClusterResourceSetBindingList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*addonsv1.ClusterResourceSetBindingList)

	return Convert_v1alpha4_ClusterResourceSetBindingList_To_v1beta2_ClusterResourceSetBindingList(src, dst, nil)
}

func (dst *ClusterResourceSetBindingList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*addonsv1.ClusterResourceSetBindingList)

	return Convert_v1beta2_ClusterResourceSetBindingList_To_v1alpha4_ClusterResourceSetBindingList(src, dst, nil)
}

// Convert_v1beta2_ClusterResourceSetBindingSpec_To_v1alpha4_ClusterResourceSetBindingSpec is a conversion function.
func Convert_v1beta2_ClusterResourceSetBindingSpec_To_v1alpha4_ClusterResourceSetBindingSpec(in *addonsv1.ClusterResourceSetBindingSpec, out *ClusterResourceSetBindingSpec, s apiconversion.Scope) error {
	// Spec.ClusterName does not exist in ClusterResourceSetBinding v1alpha4 API.
	return autoConvert_v1beta2_ClusterResourceSetBindingSpec_To_v1alpha4_ClusterResourceSetBindingSpec(in, out, s)
}

func Convert_v1beta2_ClusterResourceSetStatus_To_v1alpha4_ClusterResourceSetStatus(in *addonsv1.ClusterResourceSetStatus, out *ClusterResourceSetStatus, s apiconversion.Scope) error {
	// V1Beta2 was added in v1beta1
	return autoConvert_v1beta2_ClusterResourceSetStatus_To_v1alpha4_ClusterResourceSetStatus(in, out, s)
}
