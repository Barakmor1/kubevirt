/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2020 Red Hat, Inc.
 *
 */

package libvmi

import (
	. "github.com/onsi/gomega"
	k8sv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/utils/pointer"

	v1 "kubevirt.io/api/core/v1"
)

// Option represents an action that enables an option.
type Option func(vmi *v1.VirtualMachineInstance)

// DiskOption represents an action that enables an option to disk and volume in Compile time .
type DiskOption struct {
	withVolume func(volume v1.Volume) Option
}

func (d DiskOption) WithVolume(volume v1.Volume) Option {
	return d.withVolume(volume)
}

// New instantiates a new VMI configuration,
// building its properties based on the specified With* options.
func New(opts ...Option) *v1.VirtualMachineInstance {
	vmi := baseVmi(randName())

	WithTerminationGracePeriod(0)(vmi)
	for _, f := range opts {
		f(vmi)
	}

	return vmi
}

// randName returns a random name for a virtual machine
func randName() string {
	const randomPostfixLen = 5
	return "testvmi" + "-" + rand.String(randomPostfixLen)
}

// WithLabel sets a label with specified value
func WithLabel(key, value string) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Labels == nil {
			vmi.Labels = map[string]string{}
		}
		vmi.Labels[key] = value
	}
}

// WithAnnotation adds an annotation with specified value
func WithAnnotation(key, value string) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Annotations == nil {
			vmi.Annotations = map[string]string{}
		}
		vmi.Annotations[key] = value
	}
}

// WithTerminationGracePeriod specifies the termination grace period in seconds.
func WithTerminationGracePeriod(seconds int64) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		vmi.Spec.TerminationGracePeriodSeconds = &seconds
	}
}

// WithRng adds `rng` to the the vmi devices.
func WithRng() Option {
	return func(vmi *v1.VirtualMachineInstance) {
		vmi.Spec.Domain.Devices.Rng = &v1.Rng{}
	}
}

// WithResourceMemory specifies the vmi memory resource.
func WithResourceMemory(value string) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Spec.Domain.Resources.Requests == nil {
			vmi.Spec.Domain.Resources.Requests = k8sv1.ResourceList{}
		}
		vmi.Spec.Domain.Resources.Requests[k8sv1.ResourceMemory] = resource.MustParse(value)
	}
}

// WithResourceCpu specifies the vmi cpu resource.
func WithResourceCpu(value string) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Spec.Domain.Resources.Requests == nil {
			vmi.Spec.Domain.Resources.Requests = k8sv1.ResourceList{}
		}
		vmi.Spec.Domain.Resources.Requests[k8sv1.ResourceCPU] = resource.MustParse(value)
	}
}

// WithNodeSelectorFor ensures that the VMI gets scheduled on the specified node
func WithNodeSelectorFor(node *k8sv1.Node) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Spec.NodeSelector == nil {
			vmi.Spec.NodeSelector = map[string]string{}
		}
		vmi.Spec.NodeSelector["kubernetes.io/hostname"] = node.Name
	}
}

// WithSecureBoot configures EFI bootloader and SecureBoot.
func WithSecureBoot(secureBoot bool) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Spec.Domain.Firmware == nil {
			vmi.Spec.Domain.Firmware = &v1.Firmware{}
		}
		if vmi.Spec.Domain.Firmware.Bootloader == nil {
			vmi.Spec.Domain.Firmware.Bootloader = &v1.Bootloader{}
		}
		if vmi.Spec.Domain.Firmware.Bootloader.EFI == nil {
			vmi.Spec.Domain.Firmware.Bootloader.EFI = &v1.EFI{}
		}
		if secureBoot {
			// secureBoot Requires SMM to be enabled
			vmi.Spec.Domain.Firmware.Bootloader.EFI.SecureBoot = pointer.Bool(secureBoot)
			vmi.Spec.Domain.Features.SMM.Enabled = pointer.Bool(secureBoot)
		} else {
			vmi.Spec.Domain.Firmware.Bootloader.EFI.SecureBoot = pointer.Bool(secureBoot)
		}

	}
}

// WithSerialBios if set the BIOS output will be transmitted over serial.
func WithSerialBios(secureBoot bool) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Spec.Domain.Firmware == nil {
			vmi.Spec.Domain.Firmware = &v1.Firmware{}
		}
		if vmi.Spec.Domain.Firmware.Bootloader == nil {
			vmi.Spec.Domain.Firmware.Bootloader = &v1.Bootloader{}
		}
		if vmi.Spec.Domain.Firmware.Bootloader.BIOS == nil {
			vmi.Spec.Domain.Firmware.Bootloader.BIOS = &v1.BIOS{}
		}
		vmi.Spec.Domain.Firmware.Bootloader.BIOS.UseSerial = pointer.BoolPtr(secureBoot)

	}
}

// WithSEV adds `launchSecurity` with `sev`.
func WithSEV() Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Spec.Domain.LaunchSecurity == nil {
			vmi.Spec.Domain.LaunchSecurity = &v1.LaunchSecurity{}
		}
		vmi.Spec.Domain.LaunchSecurity.SEV = &v1.SEV{}
	}
}

func baseVmi(name string) *v1.VirtualMachineInstance {
	vmi := v1.NewVMIReferenceFromNameWithNS("", name)
	vmi.Spec = v1.VirtualMachineInstanceSpec{Domain: v1.DomainSpec{}}
	vmi.TypeMeta = k8smetav1.TypeMeta{
		APIVersion: v1.GroupVersion.String(),
		Kind:       "VirtualMachineInstance",
	}
	return vmi
}

// WithLabel sets labels with specified values
func WithLabels(labels map[string]string) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Labels == nil {
			vmi.Labels = labels
			return
		}
		for label, value := range labels {
			vmi.Labels[label] = value
		}
	}
}

// WithAnnotation adds annotations with specified values
func WithAnnotations(annotations map[string]string) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		if vmi.Annotations == nil {
			vmi.Annotations = annotations
			return
		}
		for annotation, value := range annotations {
			vmi.Annotations[annotation] = value
		}
	}
}

func WithClientPassthrough() Option {
	return func(vmi *v1.VirtualMachineInstance) {
		vmi.Spec.Domain.Devices.ClientPassthrough = &v1.ClientPassthroughDevices{}
	}
}

// WithResourceMemory specifies the vmi memory resource.
func With1MiResourceMemory() Option {
	return WithResourceMemory("1Mi")
}

// WithMachineType specifies the vmi machine type.
func WithMachineType(value string) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		vmi.Spec.Domain.Machine = &v1.Machine{Type: value}
	}
}

// WithMachineType specifies the vmi Scheduler.
func WithScheduler(name string) Option {
	return func(vmi *v1.VirtualMachineInstance) {
		vmi.Spec.SchedulerName = name
	}
}

// WithSEV adds `launchSecurity` with `sev`.
func WithDisk(disk v1.Disk) DiskOption {
	attachDiskAndVolume := func(volume v1.Volume) Option {
		return func(vmi *v1.VirtualMachineInstance) {
			vmi.Spec.Volumes = append(vmi.Spec.Volumes, volume)
			vmi.Spec.Domain.Devices.Disks = append(vmi.Spec.Domain.Devices.Disks, disk)
		}
	}

	return DiskOption{
		func(volume v1.Volume) Option {
			Expect(disk.Name).To(Equal(volume.Name),
				"The disk name must match the volume name")
			return attachDiskAndVolume(volume)
		},
	}

}
