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
 * Copyright 2018 Red Hat, Inc.
 *
 */

package utils

import (
	"fmt"
	"os"
	"strings"

	resourcev1beta1 "k8s.io/api/resource/v1beta1"

	"k8s.io/apimachinery/pkg/util/rand"
	"kubevirt.io/api/migrations/v1alpha1"
	"kubevirt.io/client-go/kubecli"

	k8sv1 "k8s.io/api/core/v1"
	schedulingv1 "k8s.io/api/scheduling/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	v1 "kubevirt.io/api/core/v1"
	instancetypev1beta1 "kubevirt.io/api/instancetype/v1beta1"
	poolv1 "kubevirt.io/api/pool/v1alpha1"
	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"

	"kubevirt.io/kubevirt/pkg/pointer"
)

const (
	strFmt                     = "%s/%s:%s"
	kubevirtIoVM               = "kubevirt.io/vm"
	vmName                     = "vm-${NAME}"
	kubevirtVM                 = "kubevirt-vm"
	githubKubevirtIsVMTemplate = "miq.github.io/kubevirt-is-vm-template"
	rhel74                     = "rhel-7.4"
)

const (
	VmiEphemeral                = "vmi-ephemeral"
	VmiMigratable               = "vmi-migratable"
	VmiInstancetypeSmall        = "vmi-instancetype-small"
	VmiSata                     = "vmi-sata"
	VmiFedora                   = "vmi-fedora"
	VmiFedoraIsolated           = "vmi-fedora-isolated"
	VmiSecureBoot               = "vmi-secureboot"
	VmiAlpineEFI                = "vmi-alpine-efi"
	VmiNoCloud                  = "vmi-nocloud"
	VmiPVC                      = "vmi-pvc"
	VmiWindows                  = "vmi-windows"
	VmiKernelBoot               = "vmi-kernel-boot"
	VmiMasquerade               = "vmi-masquerade"
	VmiSRIOV                    = "vmi-sriov"
	VmiWithHookSidecar          = "vmi-with-sidecar-hook"
	VmiWithHookSidecarConfigMap = "vmi-with-sidecar-hook-configmap"
	VmiMultusPtp                = "vmi-multus-ptp"
	VmiMultusMultipleNet        = "vmi-multus-multiple-net"
	VmiHostDisk                 = "vmi-host-disk"
	VmiGPU                      = "vmi-gpu"
	VmiARM                      = "vmi-arm"
	VmiUSB                      = "vmi-usb"
	VmiDRAGPU                   = "vmi-dra-pgpu"
	VmTemplateFedora            = "vm-template-fedora"
	VmTemplateRHEL7             = "vm-template-rhel7"
	VmTemplateWindows           = "vm-template-windows2012r2"
)

const (
	Preemtible    = "preemtible"
	NonPreemtible = "non-preemtible"
)

const (
	VirtualMachineInstancetypeComputeSmall              = "csmall"
	VirtualMachineClusterInstancetypeComputeSmall       = "cluster-csmall"
	VirtualMachineInstancetypeComputeLarge              = "clarge"
	VirtualMachinePreferenceVirtio                      = v1.VirtIO
	VirtualMachinePreferenceWindows                     = "windows"
	VmCirrosInstancetypeComputeSmall                    = "vm-cirros-csmall"
	VmCirrosClusterInstancetypeComputeSmall             = "vm-cirros-cluster-csmall"
	VmCirrosInstancetypeComputeLarge                    = "vm-cirros-clarge"
	VmCirrosInstancetypeComputeLargePreferncesVirtio    = "vm-cirros-clarge-virtio"
	VmCirrosInstancetypeComputeLargePreferencesWindows  = "vm-cirros-clarge-windows"
	VmWindowsInstancetypeComputeLargePreferencesWindows = "vm-windows-clarge-windows"
)

const (
	VmCirros                         = "vm-cirros"
	VmAlpineMultiPvc                 = "vm-alpine-multipvc"
	VmAlpineDataVolume               = "vm-alpine-datavolume"
	VMPriorityClass                  = "vm-priorityclass"
	VmCirrosSata                     = "vm-cirros-sata"
	VmCirrosWithHookSidecarConfigMap = "vm-cirros-with-sidecar-hook-configmap"
)

const VmiReplicaSetCirros = "vmi-replicaset-cirros"

const VmPoolCirros = "vm-pool-cirros"

const VmiPresetSmall = "vmi-preset-small"

const VmiMigration = "migration-job"

const MigrationPolicyName = "example-migration-policy"

const ResourceClaimTemplatePGPU = "pgpu-resource-claim-tmpl"

const DRARequestName = "pgpu"

const DRAResourceClaimName = "pgpu-resource-claim"

const (
	imageAlpine     = "alpine-container-disk-demo"
	imageCirros     = "cirros-container-disk-demo"
	imageFedora     = "fedora-with-test-tooling-container-disk"
	imageKernelBoot = "alpine-ext-kernel-boot-demo"
)
const windowsFirmware = "5d307ca9-b3ef-428c-8861-06e72d69f223"
const EthernetAdaptorModelToEnableMultiqueue = v1.VirtIO

const (
	cloudConfigHeader = "#cloud-config"

	cloudConfigInstallAndStartService = `packages:
  - nginx
runcmd:
  - [ "systemctl", "enable", "--now", "nginx" ]`

	cloudConfigUserPassword = `password: fedora
chpasswd: { expire: False }`

	secondaryIfaceDhcpNetworkData = `version: 2
ethernets:
  eth1:
    dhcp4: true
`
)

var DockerPrefix = "registry:5000/kubevirt"
var DockerTag = "devel"

var gracePeriod = int64(0)

func getBaseVMISpec() *v1.VirtualMachineInstanceSpec {
	return &v1.VirtualMachineInstanceSpec{
		TerminationGracePeriodSeconds: &gracePeriod,
		Domain: v1.DomainSpec{
			Memory: &v1.Memory{
				Guest: pointer.P(resource.MustParse("128Mi")),
			},
		},
	}
}

func getBaseVMI(name string) *v1.VirtualMachineInstance {
	baseVMISpec := getBaseVMISpec()

	return &v1.VirtualMachineInstance{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "VirtualMachineInstance",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: map[string]string{"special": name},
		},
		Spec: *baseVMISpec,
	}
}

func initFedora(spec *v1.VirtualMachineInstanceSpec) *v1.VirtualMachineInstanceSpec {
	addContainerDisk(spec, fmt.Sprintf(strFmt, DockerPrefix, imageFedora, DockerTag), v1.DiskBusVirtio)
	addRNG(spec) // without RNG, newer fedora images may hang waiting for entropy sources
	return spec
}

func initFedoraIsolated(spec *v1.VirtualMachineInstanceSpec) *v1.VirtualMachineInstanceSpec {
	addContainerDisk(spec, fmt.Sprintf(strFmt, DockerPrefix, imageFedora, DockerTag), v1.DiskBusVirtio)
	addRNG(spec) // without RNG, newer fedora images may hang waiting for entropy sources

	addDedicatedAndIsolatedCPU(spec)
	return spec
}

func addDedicatedAndIsolatedCPU(spec *v1.VirtualMachineInstanceSpec) *v1.VirtualMachineInstanceSpec {
	cpu := &v1.CPU{
		IsolateEmulatorThread: true,
		DedicatedCPUPlacement: true,
		Sockets:               1,
		Cores:                 1,
		Threads:               1,
	}
	spec.Domain.CPU = cpu
	return spec
}

func addRNG(spec *v1.VirtualMachineInstanceSpec) *v1.VirtualMachineInstanceSpec {
	spec.Domain.Devices.Rng = &v1.Rng{}
	return spec
}

func addContainerDisk(spec *v1.VirtualMachineInstanceSpec, image string, bus v1.DiskBus) *v1.VirtualMachineInstanceSpec {
	// Only add a reference to the disk if it isn't using the default v1.DiskBusSATA bus
	if bus != v1.DiskBusSATA {
		disk := &v1.Disk{
			Name: "containerdisk",
			DiskDevice: v1.DiskDevice{
				Disk: &v1.DiskTarget{
					Bus: bus,
				},
			},
		}
		spec.Domain.Devices.Disks = append(spec.Domain.Devices.Disks, *disk)
	}
	volume := &v1.Volume{
		Name: "containerdisk",
		VolumeSource: v1.VolumeSource{
			ContainerDisk: &v1.ContainerDiskSource{
				Image: image,
			},
		},
	}
	spec.Volumes = append(spec.Volumes, *volume)
	return spec
}

func addKernelBootContainer(spec *v1.VirtualMachineInstanceSpec, image, kernelArgs, kernelPath, initrdPath string) *v1.VirtualMachineInstanceSpec {
	if spec.Domain.Firmware == nil {
		spec.Domain.Firmware = &v1.Firmware{}
	}

	spec.Domain.Firmware.KernelBoot = &v1.KernelBoot{
		KernelArgs: kernelArgs,
		Container: &v1.KernelBootContainer{
			Image:      image,
			KernelPath: kernelPath,
			InitrdPath: initrdPath,
		},
	}

	return spec
}

func addNoCloudDisk(spec *v1.VirtualMachineInstanceSpec) *v1.VirtualMachineInstanceSpec {
	return addNoCloudDiskWitUserData(spec, "#!/bin/sh\n\necho 'printed from cloud-init userdata'\n")
}

func addNoCloudDiskWitUserData(spec *v1.VirtualMachineInstanceSpec, data string) *v1.VirtualMachineInstanceSpec {
	spec.Domain.Devices.Disks = append(spec.Domain.Devices.Disks, v1.Disk{
		Name: "cloudinitdisk",
		DiskDevice: v1.DiskDevice{
			Disk: &v1.DiskTarget{
				Bus: v1.DiskBusVirtio,
			},
		},
	})

	spec.Volumes = append(spec.Volumes, v1.Volume{
		Name: "cloudinitdisk",
		VolumeSource: v1.VolumeSource{
			CloudInitNoCloud: &v1.CloudInitNoCloudSource{
				UserData: data,
			},
		},
	})
	return spec
}

func addNoCloudDiskWitUserDataNetworkData(spec *v1.VirtualMachineInstanceSpec, userData string, networkData string) *v1.VirtualMachineInstanceSpec {
	spec.Domain.Devices.Disks = append(spec.Domain.Devices.Disks, v1.Disk{
		Name: "cloudinitdisk",
		DiskDevice: v1.DiskDevice{
			Disk: &v1.DiskTarget{
				Bus: v1.DiskBusVirtio,
			},
		},
	})

	spec.Volumes = append(spec.Volumes, v1.Volume{
		Name: "cloudinitdisk",
		VolumeSource: v1.VolumeSource{
			CloudInitNoCloud: &v1.CloudInitNoCloudSource{
				UserData:    userData,
				NetworkData: networkData,
			},
		},
	})
	return spec
}

func addEmptyDisk(spec *v1.VirtualMachineInstanceSpec, size string) *v1.VirtualMachineInstanceSpec {
	spec.Domain.Devices.Disks = append(spec.Domain.Devices.Disks, v1.Disk{
		Name: "emptydisk",
		DiskDevice: v1.DiskDevice{
			Disk: &v1.DiskTarget{
				Bus: v1.DiskBusVirtio,
			},
		},
	})

	spec.Volumes = append(spec.Volumes, v1.Volume{
		Name: "emptydisk",
		VolumeSource: v1.VolumeSource{
			EmptyDisk: &v1.EmptyDiskSource{
				Capacity: resource.MustParse(size),
			},
		},
	})
	return spec
}

func addDataVolumeDisk(spec *v1.VirtualMachineInstanceSpec, dataVolumeName string, bus v1.DiskBus, diskName string) *v1.VirtualMachineInstanceSpec {

	// Only add a reference to the disk if it isn't using the default v1.DiskBusSATA bus
	if bus != v1.DiskBusSATA {
		spec.Domain.Devices.Disks = append(spec.Domain.Devices.Disks, v1.Disk{
			Name: diskName,
			DiskDevice: v1.DiskDevice{
				Disk: &v1.DiskTarget{
					Bus: bus,
				},
			},
		})
	}

	spec.Volumes = append(spec.Volumes, v1.Volume{
		Name: diskName,
		VolumeSource: v1.VolumeSource{
			DataVolume: &v1.DataVolumeSource{
				Name: dataVolumeName,
			},
		},
	})
	return spec
}

func addPVCDisk(spec *v1.VirtualMachineInstanceSpec, claimName string, bus v1.DiskBus, diskName string) *v1.VirtualMachineInstanceSpec {

	// Only add a reference to the disk if it isn't using the default v1.DiskBusSATA bus
	if bus != v1.DiskBusSATA {
		spec.Domain.Devices.Disks = append(spec.Domain.Devices.Disks, v1.Disk{
			Name: diskName,
			DiskDevice: v1.DiskDevice{
				Disk: &v1.DiskTarget{
					Bus: bus,
				},
			},
		})
	}

	spec.Volumes = append(spec.Volumes, v1.Volume{
		Name: diskName,
		VolumeSource: v1.VolumeSource{
			PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
				PersistentVolumeClaimVolumeSource: k8sv1.PersistentVolumeClaimVolumeSource{ClaimName: claimName},
			},
		},
	})
	return spec
}

func addHostDisk(spec *v1.VirtualMachineInstanceSpec, path string, hostDiskType v1.HostDiskType, size string) *v1.VirtualMachineInstanceSpec {
	spec.Domain.Devices.Disks = append(spec.Domain.Devices.Disks, v1.Disk{
		Name: "host-disk",
		DiskDevice: v1.DiskDevice{
			Disk: &v1.DiskTarget{
				Bus: v1.DiskBusVirtio,
			},
		},
	})
	spec.Volumes = append(spec.Volumes, v1.Volume{
		Name: "host-disk",
		VolumeSource: v1.VolumeSource{
			HostDisk: &v1.HostDisk{
				Path:     path,
				Type:     hostDiskType,
				Capacity: resource.MustParse(size),
			},
		},
	})
	return spec
}

func GetVMIMigratable() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiMigratable)
	makeMigratable(vmi)

	addContainerDisk(&vmi.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageAlpine, DockerTag), v1.DiskBusVirtio)
	return vmi
}

func GetVMIEphemeral() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiEphemeral)

	addContainerDisk(&vmi.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), v1.DiskBusVirtio)
	return vmi
}

func GetVMISata() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiSata)

	addContainerDisk(&vmi.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), v1.DiskBusSATA)
	return vmi
}

func GetVMIEphemeralFedora() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiFedora)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	makeMigratable(vmi)
	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword))
	return vmi
}

func GetVMIEphemeralFedoraIsolated() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiFedora)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	vmi.Spec.Domain.Resources = v1.ResourceRequirements{
		Requests: k8sv1.ResourceList{
			k8sv1.ResourceMemory: resource.MustParse("1024M"),
		},
	}
	initFedoraIsolated(&vmi.Spec)
	addNoCloudDiskWitUserData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword))
	return vmi
}

func GetVMISecureBoot() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiSecureBoot)

	addContainerDisk(&vmi.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageFedora, DockerTag), v1.DiskBusVirtio)

	_true := true
	vmi.Spec.Domain.Features = &v1.Features{
		SMM: &v1.FeatureState{
			Enabled: &_true,
		},
	}
	vmi.Spec.Domain.Firmware = &v1.Firmware{
		Bootloader: &v1.Bootloader{
			EFI: &v1.EFI{
				SecureBoot: &_true,
			},
		},
	}

	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1Gi"))
	return vmi
}

func GetVMIAlpineEFI() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiAlpineEFI)

	_false := false
	addContainerDisk(&vmi.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageAlpine, DockerTag), v1.DiskBusVirtio)
	vmi.Spec.Domain.Firmware = &v1.Firmware{
		Bootloader: &v1.Bootloader{
			EFI: &v1.EFI{
				SecureBoot: &_false,
			},
		},
	}

	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1Gi"))
	return vmi
}

func GetVMIMasquerade() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiMasquerade)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	vmi.Spec.Networks = []v1.Network{{Name: "testmasquerade", NetworkSource: v1.NetworkSource{Pod: &v1.PodNetwork{}}}}
	initFedora(&vmi.Spec)
	networkData := "version: 2\nethernets:\n  eth0:\n    addresses: [ fd10:0:2::2/120 ]\n    dhcp4: true\n    gateway6: fd10:0:2::1\n"
	addNoCloudDiskWitUserDataNetworkData(
		&vmi.Spec,
		generateCloudConfigString(cloudConfigUserPassword, cloudConfigInstallAndStartService),
		networkData)

	masquerade := &v1.InterfaceMasquerade{}
	ports := []v1.Port{{Name: "http", Protocol: "TCP", Port: 80}}
	vmi.Spec.Domain.Devices.Interfaces = []v1.Interface{{Name: "testmasquerade", Ports: ports, InterfaceBindingMethod: v1.InterfaceBindingMethod{Masquerade: masquerade}}}

	return vmi
}

func GetVMISRIOV() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiSRIOV)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	vmi.Spec.Networks = []v1.Network{*v1.DefaultPodNetwork(), {Name: "sriov-net", NetworkSource: v1.NetworkSource{Multus: &v1.MultusNetwork{NetworkName: "sriov/sriov-network"}}}}
	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserDataNetworkData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword), secondaryIfaceDhcpNetworkData)

	vmi.Spec.Domain.Devices.Interfaces = []v1.Interface{{Name: "default", InterfaceBindingMethod: v1.InterfaceBindingMethod{Masquerade: &v1.InterfaceMasquerade{}}},
		{Name: "sriov-net", InterfaceBindingMethod: v1.InterfaceBindingMethod{SRIOV: &v1.InterfaceSRIOV{}}}}

	return vmi
}

func GetVMIMultusPtp() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiMultusPtp)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	vmi.Spec.Networks = []v1.Network{{Name: "ptp", NetworkSource: v1.NetworkSource{Multus: &v1.MultusNetwork{NetworkName: "ptp-conf"}}}}
	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword))

	vmi.Spec.Domain.Devices.Interfaces = []v1.Interface{{Name: "ptp", InterfaceBindingMethod: v1.InterfaceBindingMethod{Bridge: &v1.InterfaceBridge{}}}}

	return vmi
}

func GetVMIMultusMultipleNet() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiMultusMultipleNet)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	vmi.Spec.Networks = []v1.Network{*v1.DefaultPodNetwork(), {Name: "ptp", NetworkSource: v1.NetworkSource{Multus: &v1.MultusNetwork{NetworkName: "ptp-conf"}}}}
	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserDataNetworkData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword), secondaryIfaceDhcpNetworkData)

	vmi.Spec.Domain.Devices.Interfaces = []v1.Interface{{Name: "default", InterfaceBindingMethod: v1.InterfaceBindingMethod{Masquerade: &v1.InterfaceMasquerade{}}},
		{Name: "ptp", InterfaceBindingMethod: v1.InterfaceBindingMethod{Bridge: &v1.InterfaceBridge{}}}}

	return vmi
}

func GetVMINoCloud() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiNoCloud)

	addContainerDisk(&vmi.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), v1.DiskBusVirtio)
	addNoCloudDisk(&vmi.Spec)
	addEmptyDisk(&vmi.Spec, "2Gi")
	return vmi
}

func GetVMIPvc() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiPVC)

	addPVCDisk(&vmi.Spec, "disk-alpine", v1.DiskBusVirtio, "pvcdisk")
	return vmi
}

func GetVMIHostDisk() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiHostDisk)
	addHostDisk(&vmi.Spec, "/var/data/disk.img", v1.HostDiskExistsOrCreate, "1Gi")
	return vmi
}

func GetVMIWindows() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiWindows)

	gracePeriod := int64(0)
	spinlocks := uint32(8191)
	firmware := types.UID(windowsFirmware)
	_true := true
	_false := false
	vmi.Spec = v1.VirtualMachineInstanceSpec{
		TerminationGracePeriodSeconds: &gracePeriod,
		Domain: v1.DomainSpec{
			CPU: &v1.CPU{Cores: 2},
			Features: &v1.Features{
				ACPI: v1.FeatureState{},
				APIC: &v1.FeatureAPIC{},
				Hyperv: &v1.FeatureHyperv{
					Relaxed:   &v1.FeatureState{},
					VAPIC:     &v1.FeatureState{},
					Spinlocks: &v1.FeatureSpinlocks{Retries: &spinlocks},
				},
				SMM: &v1.FeatureState{},
			},
			Clock: &v1.Clock{
				ClockOffset: v1.ClockOffset{UTC: &v1.ClockOffsetUTC{}},
				Timer: &v1.Timer{
					HPET:   &v1.HPETTimer{Enabled: &_false},
					PIT:    &v1.PITTimer{TickPolicy: v1.PITTickPolicyDelay},
					RTC:    &v1.RTCTimer{TickPolicy: v1.RTCTickPolicyCatchup},
					Hyperv: &v1.HypervTimer{},
				},
			},
			Firmware: &v1.Firmware{
				UUID: firmware,
				Bootloader: &v1.Bootloader{
					EFI: &v1.EFI{SecureBoot: &_true},
				},
			},
			Memory: &v1.Memory{
				Guest: pointer.P(resource.MustParse("2048Mi")),
			},
			Devices: v1.Devices{
				Interfaces: []v1.Interface{*v1.DefaultMasqueradeNetworkInterface()},
				TPM:        &v1.TPMDevice{},
			},
		},
		Networks: []v1.Network{*v1.DefaultPodNetwork()},
	}

	// pick e1000 network model type for windows machines
	vmi.Spec.Domain.Devices.Interfaces[0].Model = "e1000"

	addPVCDisk(&vmi.Spec, "disk-windows", v1.DiskBusSATA, "pvcdisk")
	return vmi
}

func GetVMIKernelBoot() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiKernelBoot)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1Gi"))
	AddKernelBootToVMI(vmi)
	return vmi
}

func GetVMIKernelBootWithRandName() *v1.VirtualMachineInstance {
	vmi := GetVMIKernelBoot()
	vmi.Name += "-" + rand.String(5)

	return vmi
}

func AddKernelBootToVMI(vmi *v1.VirtualMachineInstance) {
	image := fmt.Sprintf(strFmt, DockerPrefix, imageKernelBoot, DockerTag)
	const KernelArgs = "console=ttyS0"
	const kernelPath = "/boot/vmlinuz-virt"
	const initrdPath = "/boot/initramfs-virt"

	addKernelBootContainer(&vmi.Spec, image, KernelArgs, kernelPath, initrdPath)
}

func getBaseVM(name string, labels map[string]string) *v1.VirtualMachine {
	baseVMISpec := getBaseVMISpec()

	return &v1.VirtualMachine{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "VirtualMachine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
		Spec: v1.VirtualMachineSpec{
			RunStrategy: pointer.P(v1.RunStrategyHalted),
			Template: &v1.VirtualMachineInstanceTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: *baseVMISpec,
			},
		},
	}
}

func GetPreemtible() *schedulingv1.PriorityClass {
	preemtionPolicy := k8sv1.PreemptLowerPriority
	pc := schedulingv1.PriorityClass{
		TypeMeta: metav1.TypeMeta{
			APIVersion: schedulingv1.SchemeGroupVersion.String(),
			Kind:       "PriorityClass",
		},
		GlobalDefault:    false,
		Description:      "Priority class for VMs which are allowed to be preemtited.",
		PreemptionPolicy: &preemtionPolicy,
		Value:            1000000,
	}
	pc.ObjectMeta.Name = "preemtible"
	return &pc
}

func GetNonPreemtible() *schedulingv1.PriorityClass {
	preemtionPolicy := k8sv1.PreemptNever
	pc := schedulingv1.PriorityClass{
		TypeMeta: metav1.TypeMeta{
			APIVersion: schedulingv1.SchemeGroupVersion.String(),
			Kind:       "PriorityClass",
		},
		GlobalDefault:    false,
		Description:      "Priority class for VMs which should not be preemtited.",
		PreemptionPolicy: &preemtionPolicy,
		Value:            999999999,
	}
	pc.ObjectMeta.Name = NonPreemtible
	return &pc
}

func GetVMPriorityClass() *v1.VirtualMachine {
	vm := GetVMCirros()
	vm.Spec.Template.Spec.PriorityClassName = NonPreemtible
	vm.ObjectMeta.Name = "vm-non-preemtible"
	return vm
}

func GetVMCirros() *v1.VirtualMachine {
	vm := getBaseVM(VmCirros, map[string]string{
		kubevirtIoVM: VmCirros,
	})

	addContainerDisk(&vm.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), v1.DiskBusVirtio)
	addNoCloudDisk(&vm.Spec.Template.Spec)
	return vm
}

func GetVMCirrosWithHookSidecarConfigMap() *v1.VirtualMachine {
	vm := getBaseVM(VmCirrosWithHookSidecarConfigMap, map[string]string{
		kubevirtIoVM: VmCirrosWithHookSidecarConfigMap,
	})

	addContainerDisk(&vm.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), v1.DiskBusVirtio)
	addNoCloudDisk(&vm.Spec.Template.Spec)

	if vm.Spec.Template.ObjectMeta.Annotations == nil {
		vm.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	vm.Spec.Template.ObjectMeta.Annotations["hooks.kubevirt.io/hookSidecars"] = `[{"args": ["--version", "v1alpha2"], "configMap": {"name": "my-config-map",` +
		`"key": "my_script.sh", "hookPath": "/usr/bin/onDefineDomain"}}]`

	return vm
}

func GetVMCirrosSata() *v1.VirtualMachine {
	vm := getBaseVM(VmCirrosSata, map[string]string{
		kubevirtIoVM: VmCirrosSata,
	})

	addContainerDisk(&vm.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), v1.DiskBusSATA)
	addNoCloudDisk(&vm.Spec.Template.Spec)
	vm.Spec.Template.Spec.Domain.Devices = v1.Devices{}
	return vm
}

func GetVMDataVolume() *v1.VirtualMachine {
	vm := getBaseVM(VmAlpineDataVolume, map[string]string{
		kubevirtIoVM: VmAlpineDataVolume,
	})

	quantity, err := resource.ParseQuantity("2Gi")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		panic(err)
	}
	storageClassName := "local"
	url := fmt.Sprintf("docker://%s/%s:%s", DockerPrefix, imageAlpine, DockerTag)
	dataVolumeSpec := v1.DataVolumeTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name: "alpine-dv",
		},
		Spec: cdiv1.DataVolumeSpec{
			Source: &cdiv1.DataVolumeSource{
				Registry: &cdiv1.DataVolumeSourceRegistry{
					URL: &url,
				},
			},
			PVC: &k8sv1.PersistentVolumeClaimSpec{
				AccessModes: []k8sv1.PersistentVolumeAccessMode{k8sv1.ReadWriteOnce},
				Resources: k8sv1.VolumeResourceRequirements{
					Requests: k8sv1.ResourceList{
						"storage": quantity,
					},
				},
				StorageClassName: &storageClassName,
			},
		},
	}

	vm.Spec.DataVolumeTemplates = append(vm.Spec.DataVolumeTemplates, dataVolumeSpec)
	addDataVolumeDisk(&vm.Spec.Template.Spec, "alpine-dv", v1.DiskBusVirtio, "datavolumedisk1")

	return vm
}

func GetVMMultiPvc() *v1.VirtualMachine {
	vm := getBaseVM(VmAlpineMultiPvc, map[string]string{
		kubevirtIoVM: VmAlpineMultiPvc,
	})

	addPVCDisk(&vm.Spec.Template.Spec, "disk-alpine", v1.DiskBusVirtio, "pvcdisk1")
	addPVCDisk(&vm.Spec.Template.Spec, "disk-custom", v1.DiskBusVirtio, "pvcdisk2")

	return vm
}

func getBaseVMPool(name string, replicas int, selectorLabels map[string]string) *poolv1.VirtualMachinePool {
	baseVMISpec := getBaseVMISpec()
	replicasInt32 := int32(replicas)

	return &poolv1.VirtualMachinePool{
		TypeMeta: metav1.TypeMeta{
			APIVersion: poolv1.SchemeGroupVersion.String(),
			Kind:       "VirtualMachinePool",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: poolv1.VirtualMachinePoolSpec{
			Replicas: &replicasInt32,
			Selector: &metav1.LabelSelector{
				MatchLabels: selectorLabels,
			},
			VirtualMachineTemplate: &poolv1.VirtualMachineTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: selectorLabels,
				},
				Spec: v1.VirtualMachineSpec{
					RunStrategy: pointer.P(v1.RunStrategyAlways),
					Template: &v1.VirtualMachineInstanceTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: selectorLabels,
						},
						Spec: *baseVMISpec,
					},
				},
			},
		},
	}
}

func getBaseVMIReplicaSet(name string, replicas int, selectorLabels map[string]string) *v1.VirtualMachineInstanceReplicaSet {
	baseVMISpec := getBaseVMISpec()
	replicasInt32 := int32(replicas)

	return &v1.VirtualMachineInstanceReplicaSet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "VirtualMachineInstanceReplicaSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.VirtualMachineInstanceReplicaSetSpec{
			Replicas: &replicasInt32,
			Selector: &metav1.LabelSelector{
				MatchLabels: selectorLabels,
			},
			Template: &v1.VirtualMachineInstanceTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: selectorLabels,
				},
				Spec: *baseVMISpec,
			},
		},
	}
}

func GetVMPoolCirros() *poolv1.VirtualMachinePool {

	vmPool := getBaseVMPool(VmPoolCirros, 3, map[string]string{
		"kubevirt.io/vmpool": VmPoolCirros,
	})

	addContainerDisk(&vmPool.Spec.VirtualMachineTemplate.Spec.Template.Spec, fmt.Sprintf("%s/%s:%s", DockerPrefix, imageCirros, DockerTag), v1.DiskBusVirtio)
	return vmPool
}

func GetVMIReplicaSetCirros() *v1.VirtualMachineInstanceReplicaSet {
	vmReplicaSet := getBaseVMIReplicaSet(VmiReplicaSetCirros, 3, map[string]string{
		"kubevirt.io/vmReplicaSet": VmiReplicaSetCirros,
	})

	addContainerDisk(&vmReplicaSet.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), v1.DiskBusVirtio)
	return vmReplicaSet
}

func getBaseVMIPreset(name string, selectorLabels map[string]string) *v1.VirtualMachineInstancePreset {
	return &v1.VirtualMachineInstancePreset{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "VirtualMachineInstancePreset",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.VirtualMachineInstancePresetSpec{
			Selector: metav1.LabelSelector{
				MatchLabels: selectorLabels,
			},
		},
	}
}

func GetVMIMigration() *v1.VirtualMachineInstanceMigration {
	return &v1.VirtualMachineInstanceMigration{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "VirtualMachineInstanceMigration",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: VmiMigration,
		},
		Spec: v1.VirtualMachineInstanceMigrationSpec{
			VMIName: VmiMigratable,
		},
	}
}

func GetMigrationPolicy() *v1alpha1.MigrationPolicy {
	policy := kubecli.NewMinimalMigrationPolicy(MigrationPolicyName)
	policy.Spec = v1alpha1.MigrationPolicySpec{
		AllowAutoConverge:       pointer.P(false),
		BandwidthPerMigration:   pointer.P(resource.MustParse("2000Mi")),
		CompletionTimeoutPerGiB: pointer.P(int64(123456789)),
		AllowPostCopy:           pointer.P(false),
		Selectors: &v1alpha1.Selectors{
			NamespaceSelector:              map[string]string{"namespace-key": "namespace-value"},
			VirtualMachineInstanceSelector: map[string]string{"vmi-key": "vmi-value"},
		},
	}

	return policy
}

func GetVMIPresetSmall() *v1.VirtualMachineInstancePreset {
	vmPreset := getBaseVMIPreset(VmiPresetSmall, map[string]string{
		"kubevirt.io/vmPreset": VmiPresetSmall,
	})

	vmPreset.Spec.Domain = &v1.DomainSpec{
		Memory: &v1.Memory{
			Guest: pointer.P(resource.MustParse("128Mi")),
		},
	}
	return vmPreset
}

func GetVMIWithHookSidecar() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiWithHookSidecar)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))

	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword))

	vmi.ObjectMeta.Annotations = map[string]string{
		"hooks.kubevirt.io/hookSidecars":              fmt.Sprintf("[{\"args\": [\"--version\", \"v1alpha2\"], \"image\": \"%s/example-hook-sidecar:%s\"}]", DockerPrefix, DockerTag),
		"smbios.vm.kubevirt.io/baseBoardManufacturer": "Radical Edward",
	}
	return vmi
}

func GetVmiWithHookSidecarConfigMap() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiWithHookSidecarConfigMap)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))

	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword))

	annotation := `[{"args": ["--version", "v1alpha2"], "configMap": {"name": "my-config-map",` +
		`"key": "my_script.sh", "hookPath": "/usr/bin/onDefineDomain"}}]`

	vmi.ObjectMeta.Annotations = map[string]string{
		"hooks.kubevirt.io/hookSidecars": annotation,
	}
	// TODO: also add the ConfigMap in generated example. Refer https://github.com/kubevirt/kubevirt/pull/10479#discussion_r1362021721
	return vmi
}

func GetVMIGPU() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiGPU)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	GPUs := []v1.GPU{
		{
			Name:       "gpu1",
			DeviceName: "nvidia.com/GP102GL_Tesla_P40",
		},
	}
	vmi.Spec.Domain.Devices.GPUs = GPUs
	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword))
	return vmi
}

func GetResourceClaimTemplatePGPU() *resourcev1beta1.ResourceClaimTemplate {
	return &resourcev1beta1.ResourceClaimTemplate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: resourcev1beta1.SchemeGroupVersion.String(),
			Kind:       "ResourceClaimTemplate",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: ResourceClaimTemplatePGPU,
		},
		Spec: resourcev1beta1.ResourceClaimTemplateSpec{
			Spec: resourcev1beta1.ResourceClaimSpec{
				Devices: resourcev1beta1.DeviceClaim{
					Requests: []resourcev1beta1.DeviceRequest{
						{
							Name:            DRARequestName,
							DeviceClassName: "gpu.example.com",
						},
					},
				},
			},
		},
	}
}

func getDRAGPUDevice(claimName string) []v1.GPU {
	return []v1.GPU{
		{
			Name: "example-gpu",
			ClaimRequest: &v1.ClaimRequest{
				ClaimName:   &claimName,
				RequestName: pointer.P(DRARequestName),
			},
		},
	}
}

func getDRAGPUPodResourceClaims() []k8sv1.PodResourceClaim {
	return []k8sv1.PodResourceClaim{
		{
			Name:                      DRAResourceClaimName,
			ResourceClaimTemplateName: pointer.P(ResourceClaimTemplatePGPU),
		},
	}
}

func GetVMIDRAGPU() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiDRAGPU)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	vmi.Spec.ResourceClaims = getDRAGPUPodResourceClaims()
	vmi.Spec.Domain.Devices.GPUs = getDRAGPUDevice(DRAResourceClaimName)

	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword))
	return vmi
}

// The minimum memory for UEFI boot on Arm64 is 256Mi
func GetVMIARM() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiARM)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("256Mi"))
	addContainerDisk(&vmi.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), v1.DiskBusVirtio)
	addNoCloudDisk(&vmi.Spec)
	addEmptyDisk(&vmi.Spec, "2Gi")
	return vmi
}

func GetVMIUSB() *v1.VirtualMachineInstance {
	vmi := getBaseVMI(VmiUSB)
	vmi.Spec.Domain.Memory.Guest = pointer.P(resource.MustParse("1024M"))
	initFedora(&vmi.Spec)
	addNoCloudDiskWitUserData(&vmi.Spec, generateCloudConfigString(cloudConfigUserPassword, cloudConfigInstallAndStartService))

	vmi.Spec.Domain.Devices.HostDevices = append(vmi.Spec.Domain.Devices.HostDevices,
		v1.HostDevice{
			Name:       "node-usb-to-vmi-storage",
			DeviceName: "kubevirt.io/storage",
		})
	return vmi
}

func generateCloudConfigString(cloudConfigElement ...string) string {
	return strings.Join(
		append([]string{cloudConfigHeader}, cloudConfigElement...), "\n")
}

func GetComputeSmallInstancetypeSpec() instancetypev1beta1.VirtualMachineInstancetypeSpec {
	return instancetypev1beta1.VirtualMachineInstancetypeSpec{
		CPU: instancetypev1beta1.CPUInstancetype{
			Guest: uint32(1),
		},
		Memory: instancetypev1beta1.MemoryInstancetype{
			Guest: resource.MustParse("128Mi"),
		},
	}
}

func GetVirtualMachineInstancetypeComputeSmall() *instancetypev1beta1.VirtualMachineInstancetype {
	return &instancetypev1beta1.VirtualMachineInstancetype{
		TypeMeta: metav1.TypeMeta{
			APIVersion: instancetypev1beta1.SchemeGroupVersion.String(),
			Kind:       "VirtualMachineInstancetype",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: VirtualMachineInstancetypeComputeSmall,
		},
		Spec: GetComputeSmallInstancetypeSpec(),
	}
}

func GetVirtualMachineClusterInstancetypeComputeSmall() *instancetypev1beta1.VirtualMachineClusterInstancetype {
	return &instancetypev1beta1.VirtualMachineClusterInstancetype{
		TypeMeta: metav1.TypeMeta{
			APIVersion: instancetypev1beta1.SchemeGroupVersion.String(),
			Kind:       "VirtualMachineClusterInstancetype",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: VirtualMachineClusterInstancetypeComputeSmall,
		},
		Spec: GetComputeSmallInstancetypeSpec(),
	}
}

func GetVirtualMachineInstancetypeComputeLarge() *instancetypev1beta1.VirtualMachineInstancetype {
	return &instancetypev1beta1.VirtualMachineInstancetype{
		TypeMeta: metav1.TypeMeta{
			APIVersion: instancetypev1beta1.SchemeGroupVersion.String(),
			Kind:       "VirtualMachineInstancetype",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: VirtualMachineInstancetypeComputeLarge,
		},
		Spec: instancetypev1beta1.VirtualMachineInstancetypeSpec{
			CPU: instancetypev1beta1.CPUInstancetype{
				Guest: uint32(4),
			},
			Memory: instancetypev1beta1.MemoryInstancetype{
				Guest: resource.MustParse("2048Mi"),
			},
		},
	}
}

func GetVmCirrosInstancetypeComputeSmall() *v1.VirtualMachine {
	vm := getBaseVM(VmCirrosInstancetypeComputeSmall, map[string]string{
		kubevirtIoVM: VmCirrosInstancetypeComputeSmall,
	})
	vm.Spec.Instancetype = &v1.InstancetypeMatcher{
		Name: VirtualMachineInstancetypeComputeSmall,
		Kind: "VirtualMachineInstancetype",
	}
	addContainerDisk(&vm.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), "")
	addNoCloudDisk(&vm.Spec.Template.Spec)
	return vm
}

func GetVmCirrosClusterInstancetypeComputeSmall() *v1.VirtualMachine {
	vm := getBaseVM(VmCirrosClusterInstancetypeComputeSmall, map[string]string{
		kubevirtIoVM: VmCirrosClusterInstancetypeComputeSmall,
	})

	vm.Spec.Instancetype = &v1.InstancetypeMatcher{
		Name: VirtualMachineClusterInstancetypeComputeSmall,
		Kind: "VirtualMachineClusterInstancetype",
	}

	addContainerDisk(&vm.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), "")
	addNoCloudDisk(&vm.Spec.Template.Spec)
	return vm
}

func GetVmCirrosInstancetypeComputeLarge() *v1.VirtualMachine {
	vm := getBaseVM(VmCirrosInstancetypeComputeLarge, map[string]string{
		kubevirtIoVM: VmCirrosInstancetypeComputeLarge,
	})
	vm.Spec.Instancetype = &v1.InstancetypeMatcher{
		Name: VirtualMachineInstancetypeComputeLarge,
		Kind: "VirtualMachineInstancetype",
	}
	addContainerDisk(&vm.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), "")
	addNoCloudDisk(&vm.Spec.Template.Spec)
	return vm
}

func GetVirtualMachinePreferenceVirtio() *instancetypev1beta1.VirtualMachinePreference {
	return &instancetypev1beta1.VirtualMachinePreference{
		TypeMeta: metav1.TypeMeta{
			APIVersion: instancetypev1beta1.SchemeGroupVersion.String(),
			Kind:       "VirtualMachinePreference",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: VirtualMachinePreferenceVirtio,
		},
		Spec: instancetypev1beta1.VirtualMachinePreferenceSpec{
			Devices: &instancetypev1beta1.DevicePreferences{
				PreferredDiskBus:        v1.VirtIO,
				PreferredInterfaceModel: v1.VirtIO,
			},
		},
	}
}

func GetVirtualMachinePreferenceWindows() *instancetypev1beta1.VirtualMachinePreference {
	spinlocks := uint32(8191)
	preferredCPUTopology := instancetypev1beta1.Sockets
	return &instancetypev1beta1.VirtualMachinePreference{
		TypeMeta: metav1.TypeMeta{
			APIVersion: instancetypev1beta1.SchemeGroupVersion.String(),
			Kind:       "VirtualMachinePreference",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: VirtualMachinePreferenceWindows,
		},
		Spec: instancetypev1beta1.VirtualMachinePreferenceSpec{
			CPU: &instancetypev1beta1.CPUPreferences{
				PreferredCPUTopology: &preferredCPUTopology,
			},
			Clock: &instancetypev1beta1.ClockPreferences{
				PreferredClockOffset: &v1.ClockOffset{UTC: &v1.ClockOffsetUTC{}},
				PreferredTimer: &v1.Timer{
					HPET:   &v1.HPETTimer{Enabled: pointer.P(false)},
					PIT:    &v1.PITTimer{TickPolicy: v1.PITTickPolicyDelay},
					RTC:    &v1.RTCTimer{TickPolicy: v1.RTCTickPolicyCatchup},
					Hyperv: &v1.HypervTimer{},
				},
			},
			Devices: &instancetypev1beta1.DevicePreferences{
				PreferredDiskBus:        "sata",
				PreferredInterfaceModel: "e1000",
				PreferredTPM:            &v1.TPMDevice{},
			},
			Features: &instancetypev1beta1.FeaturePreferences{
				PreferredAcpi: &v1.FeatureState{},
				PreferredApic: &v1.FeatureAPIC{},
				PreferredHyperv: &v1.FeatureHyperv{
					Relaxed:   &v1.FeatureState{},
					VAPIC:     &v1.FeatureState{},
					Spinlocks: &v1.FeatureSpinlocks{Retries: &spinlocks},
				},
				PreferredSmm: &v1.FeatureState{},
			},
			Firmware: &instancetypev1beta1.FirmwarePreferences{
				PreferredEfi: &v1.EFI{
					SecureBoot: pointer.P(true),
				},
			},
		},
	}
}

func GetVmCirrosInstancetypeComputeLargePreferencesVirtio() *v1.VirtualMachine {
	vm := getBaseVM(VmCirrosInstancetypeComputeLargePreferncesVirtio, map[string]string{
		kubevirtIoVM: VmCirrosInstancetypeComputeLargePreferncesVirtio,
	})
	vm.Spec.Instancetype = &v1.InstancetypeMatcher{
		Name: VirtualMachineInstancetypeComputeLarge,
		Kind: "VirtualMachineInstancetype",
	}
	vm.Spec.Preference = &v1.PreferenceMatcher{
		Name: VirtualMachinePreferenceVirtio,
		Kind: "VirtualMachinePreference",
	}
	addContainerDisk(&vm.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), "")
	addNoCloudDisk(&vm.Spec.Template.Spec)
	vm.Spec.Template.Spec.Domain.Devices.Disks[1].DiskDevice.Disk.Bus = ""

	return vm
}

func GetVmCirrosInstancetypeComputeLargePreferencesWindows() *v1.VirtualMachine {
	vm := getBaseVM(VmCirrosInstancetypeComputeLargePreferencesWindows, map[string]string{
		kubevirtIoVM: VmCirrosInstancetypeComputeLargePreferencesWindows,
	})
	vm.Spec.Instancetype = &v1.InstancetypeMatcher{
		Name: VirtualMachineInstancetypeComputeLarge,
		Kind: "VirtualMachineInstancetype",
	}
	vm.Spec.Preference = &v1.PreferenceMatcher{
		Name: VirtualMachinePreferenceWindows,
		Kind: "VirtualMachinePreference",
	}
	addContainerDisk(&vm.Spec.Template.Spec, fmt.Sprintf(strFmt, DockerPrefix, imageCirros, DockerTag), "")
	addNoCloudDisk(&vm.Spec.Template.Spec)
	vm.Spec.Template.Spec.Domain.Devices.Disks[1].DiskDevice.Disk.Bus = ""

	return vm
}

func GetVmWindowsInstancetypeComputeLargePreferencesWindows() *v1.VirtualMachine {
	vm := getBaseVM(VmWindowsInstancetypeComputeLargePreferencesWindows, map[string]string{
		kubevirtIoVM: VmWindowsInstancetypeComputeLargePreferencesWindows,
	})
	vm.Spec.Instancetype = &v1.InstancetypeMatcher{
		Name: VirtualMachineInstancetypeComputeLarge,
		Kind: "VirtualMachineInstancetype",
	}
	vm.Spec.Preference = &v1.PreferenceMatcher{
		Name: VirtualMachinePreferenceWindows,
		Kind: "VirtualMachinePreference",
	}

	// Do not set a disk bus, let that come from preferences
	addPVCDisk(&vm.Spec.Template.Spec, "disk-windows", "", "pvcdisk")

	// Copy the same remaining defaults as the vmi-windows example
	vm.Spec.Template.Spec.TerminationGracePeriodSeconds = pointer.P(int64(0))
	vm.Spec.Template.Spec.Domain.Firmware = &v1.Firmware{
		UUID: types.UID(windowsFirmware),
	}
	return vm
}

func makeMigratable(vmi *v1.VirtualMachineInstance) {
	// having no network leads to adding a default interface that may be of type bridge on
	// the pod network and that would make the VMI non-migratable. Therefore, adding a network.
	vmi.Spec.Networks = []v1.Network{*v1.DefaultPodNetwork()}
	vmi.Spec.Domain.Devices.Interfaces = []v1.Interface{*v1.DefaultMasqueradeNetworkInterface()}
}
