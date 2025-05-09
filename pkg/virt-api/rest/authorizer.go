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

package rest

//go:generate mockgen -source $GOFILE -package=$GOPACKAGE -destination=generated_mock_$GOFILE -imports restful=github.com/emicklei/go-restful/v3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/emicklei/go-restful/v3"
	authv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	authclientv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/client-go/util/flowcontrol"

	"kubevirt.io/client-go/kubecli"
)

const (
	userHeader            = "X-Remote-User"
	groupHeader           = "X-Remote-Group"
	userExtraHeaderPrefix = "X-Remote-Extra-"

	namespacedResourceAttributesMinParts  = 9
	namespacedResourceBaseAttributesParts = 7
)

var noAuthEndpoints = map[string]struct{}{
	"/":           {},
	"/apis":       {},
	"/healthz":    {},
	"/openapi/v2": {},
	// Although KubeVirt does not publish v3, Kubernetes aggregator controller will
	// handle v2 to v3 (lossy) conversion if KubeVirt returns 404 on this endpoint
	"/openapi/v3": {},
	// The endpoints with just the version are needed for api aggregation discovery
	// Test with e.g. kubectl get --raw /apis/subresources.kubevirt.io/v1
	"/apis/subresources.kubevirt.io/v1":               {},
	"/apis/subresources.kubevirt.io/v1/version":       {},
	"/apis/subresources.kubevirt.io/v1/guestfs":       {},
	"/apis/subresources.kubevirt.io/v1/healthz":       {},
	"/apis/subresources.kubevirt.io/v1alpha3":         {},
	"/apis/subresources.kubevirt.io/v1alpha3/version": {},
	"/apis/subresources.kubevirt.io/v1alpha3/guestfs": {},
	"/apis/subresources.kubevirt.io/v1alpha3/healthz": {},
	// the profiler endpoints are blocked by a feature gate
	// to restrict the usage to development environments
	"/start-profiler": {},
	"/stop-profiler":  {},
	"/dump-profiler":  {},
	"/apis/subresources.kubevirt.io/v1/start-cluster-profiler":       {},
	"/apis/subresources.kubevirt.io/v1/stop-cluster-profiler":        {},
	"/apis/subresources.kubevirt.io/v1/dump-cluster-profiler":        {},
	"/apis/subresources.kubevirt.io/v1alpha3/start-cluster-profiler": {},
	"/apis/subresources.kubevirt.io/v1alpha3/stop-cluster-profiler":  {},
	"/apis/subresources.kubevirt.io/v1alpha3/dump-cluster-profiler":  {},
}

type VirtApiAuthorizor interface {
	Authorize(req *restful.Request) (bool, string, error)
	AddUserHeaders(header []string)
	GetUserHeaders() []string
	AddGroupHeaders(header []string)
	GetGroupHeaders() []string
	AddExtraPrefixHeaders(header []string)
	GetExtraPrefixHeaders() []string
}

type authorizor struct {
	userHeaders             []string
	groupHeaders            []string
	userExtraHeaderPrefixes []string
	client                  authclientv1.SubjectAccessReviewInterface
}

func (a *authorizor) getUserGroups(header http.Header) ([]string, error) {
	for _, key := range a.groupHeaders {
		groups, ok := header[key]
		if ok {
			return groups, nil
		}
	}

	return nil, fmt.Errorf("a valid group header is required for authorization")
}

func (a *authorizor) getUserName(header http.Header) (string, error) {
	for _, key := range a.userHeaders {
		user, ok := header[key]
		if ok {
			return user[0], nil
		}
	}

	return "", fmt.Errorf("a valid user header is required for authorization")
}

func hasPrefixIgnoreCase(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[:len(prefix)], prefix)
}

func unescapeExtraKey(encodedKey string) string {
	key, err := url.PathUnescape(encodedKey) // Decode %-encoded bytes.
	if err != nil {
		return encodedKey // Always record extra strings, even if malformed/unencoded.
	}
	return key
}

func (a *authorizor) getUserExtras(header http.Header) map[string]authv1.ExtraValue {
	extras := map[string]authv1.ExtraValue{}

	for _, prefix := range a.userExtraHeaderPrefixes {
		for k, v := range header {
			if hasPrefixIgnoreCase(k, prefix) {
				extraKey := unescapeExtraKey(strings.ToLower(k[len(prefix):]))
				extras[extraKey] = v
			}
		}
	}

	return extras
}

func (a *authorizor) AddUserHeaders(headers []string) {
	a.userHeaders = append(a.userHeaders, headers...)
}

func (a *authorizor) GetUserHeaders() []string {
	return a.userHeaders
}

func (a *authorizor) AddGroupHeaders(headers []string) {
	a.groupHeaders = append(a.groupHeaders, headers...)
}

func (a *authorizor) GetGroupHeaders() []string {
	return a.groupHeaders
}

func (a *authorizor) AddExtraPrefixHeaders(headers []string) {
	a.userExtraHeaderPrefixes = append(a.userExtraHeaderPrefixes, headers...)
}

func (a *authorizor) GetExtraPrefixHeaders() []string {
	return a.userExtraHeaderPrefixes
}

func (a *authorizor) generateAccessReview(req *restful.Request) (*authv1.SubjectAccessReview, error) {
	if req.Request == nil {
		return nil, fmt.Errorf("empty http request")
	}
	if req.Request.URL == nil {
		return nil, fmt.Errorf("no URL in http request")
	}

	userName, err := a.getUserName(req.Request.Header)
	if err != nil {
		return nil, err
	}

	userGroups, err := a.getUserGroups(req.Request.Header)
	if err != nil {
		return nil, err
	}

	r := &authv1.SubjectAccessReview{}
	r.Spec = authv1.SubjectAccessReviewSpec{
		User:   userName,
		Groups: userGroups,
		Extra:  a.getUserExtras(req.Request.Header),
	}

	// URL examples
	// /apis/subresources.kubevirt.io/v1alpha3/namespaces/default/virtualmachineinstances/testvmi/console
	// /apis/subresources.kubevirt.io/v1alpha3/namespaces/default/expand-vm-spec
	pathSplit := strings.Split(req.Request.URL.Path, "/")
	if len(pathSplit) >= namespacedResourceAttributesMinParts {
		if err := addNamespacedResourceAttributes(pathSplit, req.Request.Method, r); err != nil {
			return nil, err
		}
	} else if len(pathSplit) == namespacedResourceBaseAttributesParts {
		if err := addNamespacedResourceBaseAttributes(pathSplit, req.Request.Method, r); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unknown api endpoint: %s", req.Request.URL.Path)
	}

	return r, nil
}

func addNamespacedResourceAttributes(pathSplit []string, requestMethod string, r *authv1.SubjectAccessReview) error {
	// URL example
	// /apis/subresources.kubevirt.io/v1alpha3/namespaces/default/virtualmachineinstances/testvmi/console
	group := pathSplit[2]
	version := pathSplit[3]
	namespace := pathSplit[5]
	resource := pathSplit[6]
	resourceName := pathSplit[7]
	subresource := pathSplit[8]

	if resource != "virtualmachineinstances" && resource != "virtualmachines" {
		return fmt.Errorf("unknown resource type %s", resource)
	}

	verb, err := mapHttpVerbToRbacVerb(requestMethod, resourceName)
	if err != nil {
		return err
	}

	r.Spec.ResourceAttributes = &authv1.ResourceAttributes{
		Namespace:   namespace,
		Verb:        verb,
		Group:       group,
		Version:     version,
		Resource:    resource,
		Subresource: subresource,
		Name:        resourceName,
	}

	return nil
}

func addNamespacedResourceBaseAttributes(pathSplit []string, requestMethod string, r *authv1.SubjectAccessReview) error {
	// URL example
	// /apis/subresources.kubevirt.io/v1alpha3/namespaces/default/expand-vm-spec
	group := pathSplit[2]
	version := pathSplit[3]
	namespace := pathSplit[5]
	resource := pathSplit[6]

	if resource != "expand-vm-spec" {
		return fmt.Errorf("unknown resource type %s", resource)
	}

	verb, err := mapHttpVerbToRbacVerb(requestMethod, "")
	if err != nil {
		return err
	}

	r.Spec.ResourceAttributes = &authv1.ResourceAttributes{
		Namespace: namespace,
		Verb:      verb,
		Group:     group,
		Version:   version,
		Resource:  resource,
	}

	return nil
}

func mapHttpVerbToRbacVerb(httpVerb string, name string) (string, error) {
	// see https://kubernetes.io/docs/reference/access-authn-authz/authorization/#determine-the-request-verb
	// if name is empty, we assume plural verbs
	switch strings.ToLower(httpVerb) {
	case strings.ToLower(http.MethodPost):
		return "create", nil
	case strings.ToLower(http.MethodGet), strings.ToLower(http.MethodHead):
		if name != "" {
			return "get", nil
		} else {
			return "list", nil
		}
	case strings.ToLower(http.MethodPut):
		return "update", nil
	case strings.ToLower(http.MethodPatch):
		return "patch", nil
	case strings.ToLower(http.MethodDelete):
		if name != "" {
			return "delete", nil
		} else {
			return "deletecollection", nil
		}
	default:
		return "", fmt.Errorf("unknown http verb in request: %v", httpVerb)
	}
}

func isNoAuthEndpoint(req *restful.Request) bool {
	if req.Request == nil || req.Request.URL == nil {
		return false
	}

	_, noAuth := noAuthEndpoints[req.Request.URL.Path]
	return noAuth
}

func isAuthenticated(req *restful.Request) bool {
	// Peer cert is required for authentication.
	// If the peer's cert is provided, we are guaranteed
	// it has been validated against our CA pool containing the requestheader CA
	if req.Request == nil || req.Request.TLS == nil || len(req.Request.TLS.PeerCertificates) == 0 {
		return false
	}
	return true
}

func (a *authorizor) Authorize(req *restful.Request) (bool, string, error) {
	// Endpoints related to getting information about
	// what apis our server provides are authorized to
	// all users.
	if isNoAuthEndpoint(req) {
		return true, "", nil
	}

	if !isAuthenticated(req) {
		return false, "request is not authenticated", nil
	}

	r, err := a.generateAccessReview(req)
	if err != nil {
		// only internal service errors are returned
		// as an error.
		// A failure to generate the access review request
		// means the client did not properly format the request.
		// Return that error as the "Reason" for the authorization failure.
		return false, fmt.Sprintf("%v", err), nil
	}

	result, err := a.client.Create(context.Background(), r, metav1.CreateOptions{})
	if err != nil {
		return false, "internal server error", err
	}

	if result.Status.Allowed {
		return true, "", nil
	}

	return false, result.Status.Reason, nil
}

func NewAuthorizorFromClient(client authclientv1.SubjectAccessReviewInterface) VirtApiAuthorizor {
	return &authorizor{
		userHeaders:             []string{userHeader},
		groupHeaders:            []string{groupHeader},
		userExtraHeaderPrefixes: []string{userExtraHeaderPrefix},
		client:                  client,
	}
}

func NewAuthorizor(rateLimiter flowcontrol.RateLimiter) (VirtApiAuthorizor, error) {
	config, err := kubecli.GetKubevirtClientConfig()
	if err != nil {
		return nil, err
	}
	config.RateLimiter = rateLimiter

	client, err := authclientv1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return NewAuthorizorFromClient(client.SubjectAccessReviews()), nil
}
