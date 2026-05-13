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
 * Copyright The KubeVirt Authors.
 *
 */

package tests_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kubevirt.io/client-go/kubecli"

	"kubevirt.io/kubevirt/pkg/libvmi"
	"kubevirt.io/kubevirt/pkg/virt-config/featuregate"
	"kubevirt.io/kubevirt/tests/decorators"
	"kubevirt.io/kubevirt/tests/framework/kubevirt"
	kvconfig "kubevirt.io/kubevirt/tests/libkubevirt/config"
	"kubevirt.io/kubevirt/tests/libmigration"
	"kubevirt.io/kubevirt/tests/libnet"
	"kubevirt.io/kubevirt/tests/libvmifact"
	"kubevirt.io/kubevirt/tests/libvmops"
)

var _ = Describe("[sig-compute]VMI Memory Overhead Reporting", decorators.SigCompute, decorators.RequiresTwoSchedulableNodes, Serial, func() {
	var virtClient kubecli.KubevirtClient

	BeforeEach(func() {
		virtClient = kubevirt.Client()
		kvconfig.EnableFeatureGate(featuregate.VmiMemoryOverheadReport)
	})

	It("should report memory overhead in VMI status and propagate it through migration", func() {
		vmi := libvmifact.NewGuestless(libvmi.WithMemoryRequest("128Mi"), libnet.WithMasqueradeNetworking())
		vmi = libvmops.RunVMIAndExpectLaunch(vmi, libvmops.StartupTimeoutSecondsSmall)

		By("Verifying memory overhead is reported in VMI status after startup")
		vmi, err := virtClient.VirtualMachineInstance(vmi.Namespace).Get(context.Background(), vmi.Name, metav1.GetOptions{})
		Expect(err).ToNot(HaveOccurred())
		Expect(vmi.Status.Memory).ToNot(BeNil())
		Expect(vmi.Status.Memory.MemoryOverhead).ToNot(BeNil(),
			"status.memory.memoryOverhead should be populated")
		Expect(vmi.Status.Memory.MemoryOverhead.Value()).To(BeNumerically(">", 0),
			"memory overhead should be a positive value")

		By("Migrating the VMI")
		migration := libmigration.New(vmi.Name, vmi.Namespace)
		migration = libmigration.RunMigrationAndExpectToCompleteWithDefaultTimeout(virtClient, migration)
		vmi, err = virtClient.VirtualMachineInstance(vmi.Namespace).Get(context.Background(), vmi.Name, metav1.GetOptions{})
		Expect(err).ToNot(HaveOccurred())

		By("Verifying that status.memory.memoryOverhead equals migrationState.targetMemoryOverhead after migration")
		Expect(vmi.Status.MigrationState).ToNot(BeNil())
		Expect(vmi.Status.MigrationState.TargetMemoryOverhead).ToNot(BeNil(),
			"migrationState.targetMemoryOverhead should be set after migration")
		Expect(vmi.Status.Memory).ToNot(BeNil())
		Expect(vmi.Status.Memory.MemoryOverhead).ToNot(BeNil(),
			"status.memory.memoryOverhead should be set after migration")
		Expect(vmi.Status.Memory.MemoryOverhead.Value()).To(Equal(vmi.Status.MigrationState.TargetMemoryOverhead.Value()),
			"status.memory.memoryOverhead should equal migrationState.targetMemoryOverhead")
	})
})
