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

package admitter

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	v1 "kubevirt.io/api/core/v1"
)

func validatePasstBinding(
	fieldPath *field.Path, idx int, iface v1.Interface, net v1.Network, config clusterConfigChecker,
) []metav1.StatusCause {
	var causes []metav1.StatusCause
	if iface.InterfaceBindingMethod.DeprecatedPasst != nil && !config.PasstEnabled() {
		causes = append(causes, metav1.StatusCause{
			Type:    metav1.CauseTypeFieldValueInvalid,
			Message: "Passt feature gate is not enabled",
			Field:   fieldPath.Child("domain", "devices", "interfaces").Index(idx).Child("name").String(),
		})
	}
	if iface.DeprecatedPasst != nil && net.Pod == nil {
		causes = append(causes, metav1.StatusCause{
			Type:    metav1.CauseTypeFieldValueInvalid,
			Message: "Passt interface only implemented with pod network",
			Field:   fieldPath.Child("domain", "devices", "interfaces").Index(idx).Child("name").String(),
		})
	}
	return causes
}
