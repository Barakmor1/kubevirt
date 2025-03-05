//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2025 The KubeVirt Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Error) DeepCopyInto(out *Error) {
	*out = *in
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = (*in).DeepCopy()
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Error.
func (in *Error) DeepCopy() *Error {
	if in == nil {
		return nil
	}
	out := new(Error)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotVolumesLists) DeepCopyInto(out *SnapshotVolumesLists) {
	*out = *in
	if in.IncludedVolumes != nil {
		in, out := &in.IncludedVolumes, &out.IncludedVolumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedVolumes != nil {
		in, out := &in.ExcludedVolumes, &out.ExcludedVolumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotVolumesLists.
func (in *SnapshotVolumesLists) DeepCopy() *SnapshotVolumesLists {
	if in == nil {
		return nil
	}
	out := new(SnapshotVolumesLists)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceSpec) DeepCopyInto(out *SourceSpec) {
	*out = *in
	if in.VirtualMachine != nil {
		in, out := &in.VirtualMachine, &out.VirtualMachine
		*out = new(VirtualMachine)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceSpec.
func (in *SourceSpec) DeepCopy() *SourceSpec {
	if in == nil {
		return nil
	}
	out := new(SourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachine) DeepCopyInto(out *VirtualMachine) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachine.
func (in *VirtualMachine) DeepCopy() *VirtualMachine {
	if in == nil {
		return nil
	}
	out := new(VirtualMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineRestore) DeepCopyInto(out *VirtualMachineRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VirtualMachineRestoreStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineRestore.
func (in *VirtualMachineRestore) DeepCopy() *VirtualMachineRestore {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineRestoreList) DeepCopyInto(out *VirtualMachineRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineRestoreList.
func (in *VirtualMachineRestoreList) DeepCopy() *VirtualMachineRestoreList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineRestoreSpec) DeepCopyInto(out *VirtualMachineRestoreSpec) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineRestoreSpec.
func (in *VirtualMachineRestoreSpec) DeepCopy() *VirtualMachineRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineRestoreStatus) DeepCopyInto(out *VirtualMachineRestoreStatus) {
	*out = *in
	if in.Restores != nil {
		in, out := &in.Restores, &out.Restores
		*out = make([]VolumeRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestoreTime != nil {
		in, out := &in.RestoreTime, &out.RestoreTime
		*out = (*in).DeepCopy()
	}
	if in.DeletedDataVolumes != nil {
		in, out := &in.DeletedDataVolumes, &out.DeletedDataVolumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Complete != nil {
		in, out := &in.Complete, &out.Complete
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineRestoreStatus.
func (in *VirtualMachineRestoreStatus) DeepCopy() *VirtualMachineRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSnapshot) DeepCopyInto(out *VirtualMachineSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VirtualMachineSnapshotStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSnapshot.
func (in *VirtualMachineSnapshot) DeepCopy() *VirtualMachineSnapshot {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSnapshotContent) DeepCopyInto(out *VirtualMachineSnapshotContent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VirtualMachineSnapshotContentStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSnapshotContent.
func (in *VirtualMachineSnapshotContent) DeepCopy() *VirtualMachineSnapshotContent {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSnapshotContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineSnapshotContent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSnapshotContentList) DeepCopyInto(out *VirtualMachineSnapshotContentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineSnapshotContent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSnapshotContentList.
func (in *VirtualMachineSnapshotContentList) DeepCopy() *VirtualMachineSnapshotContentList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSnapshotContentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineSnapshotContentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSnapshotContentSpec) DeepCopyInto(out *VirtualMachineSnapshotContentSpec) {
	*out = *in
	if in.VirtualMachineSnapshotName != nil {
		in, out := &in.VirtualMachineSnapshotName, &out.VirtualMachineSnapshotName
		*out = new(string)
		**out = **in
	}
	in.Source.DeepCopyInto(&out.Source)
	if in.VolumeBackups != nil {
		in, out := &in.VolumeBackups, &out.VolumeBackups
		*out = make([]VolumeBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSnapshotContentSpec.
func (in *VirtualMachineSnapshotContentSpec) DeepCopy() *VirtualMachineSnapshotContentSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSnapshotContentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSnapshotContentStatus) DeepCopyInto(out *VirtualMachineSnapshotContentStatus) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.ReadyToUse != nil {
		in, out := &in.ReadyToUse, &out.ReadyToUse
		*out = new(bool)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(Error)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeSnapshotStatus != nil {
		in, out := &in.VolumeSnapshotStatus, &out.VolumeSnapshotStatus
		*out = make([]VolumeSnapshotStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSnapshotContentStatus.
func (in *VirtualMachineSnapshotContentStatus) DeepCopy() *VirtualMachineSnapshotContentStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSnapshotContentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSnapshotList) DeepCopyInto(out *VirtualMachineSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSnapshotList.
func (in *VirtualMachineSnapshotList) DeepCopy() *VirtualMachineSnapshotList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSnapshotSpec) DeepCopyInto(out *VirtualMachineSnapshotSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.DeletionPolicy != nil {
		in, out := &in.DeletionPolicy, &out.DeletionPolicy
		*out = new(DeletionPolicy)
		**out = **in
	}
	if in.FailureDeadline != nil {
		in, out := &in.FailureDeadline, &out.FailureDeadline
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSnapshotSpec.
func (in *VirtualMachineSnapshotSpec) DeepCopy() *VirtualMachineSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSnapshotStatus) DeepCopyInto(out *VirtualMachineSnapshotStatus) {
	*out = *in
	if in.SourceUID != nil {
		in, out := &in.SourceUID, &out.SourceUID
		*out = new(types.UID)
		**out = **in
	}
	if in.VirtualMachineSnapshotContentName != nil {
		in, out := &in.VirtualMachineSnapshotContentName, &out.VirtualMachineSnapshotContentName
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.ReadyToUse != nil {
		in, out := &in.ReadyToUse, &out.ReadyToUse
		*out = new(bool)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(Error)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Indications != nil {
		in, out := &in.Indications, &out.Indications
		*out = make([]Indication, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotVolumes != nil {
		in, out := &in.SnapshotVolumes, &out.SnapshotVolumes
		*out = new(SnapshotVolumesLists)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSnapshotStatus.
func (in *VirtualMachineSnapshotStatus) DeepCopy() *VirtualMachineSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackup) DeepCopyInto(out *VolumeBackup) {
	*out = *in
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
	if in.VolumeSnapshotName != nil {
		in, out := &in.VolumeSnapshotName, &out.VolumeSnapshotName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackup.
func (in *VolumeBackup) DeepCopy() *VolumeBackup {
	if in == nil {
		return nil
	}
	out := new(VolumeBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeRestore) DeepCopyInto(out *VolumeRestore) {
	*out = *in
	if in.DataVolumeName != nil {
		in, out := &in.DataVolumeName, &out.DataVolumeName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeRestore.
func (in *VolumeRestore) DeepCopy() *VolumeRestore {
	if in == nil {
		return nil
	}
	out := new(VolumeRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSnapshotStatus) DeepCopyInto(out *VolumeSnapshotStatus) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.ReadyToUse != nil {
		in, out := &in.ReadyToUse, &out.ReadyToUse
		*out = new(bool)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(Error)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSnapshotStatus.
func (in *VolumeSnapshotStatus) DeepCopy() *VolumeSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}
