/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
	servicemeshv1alpha2 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/servicemesh/v1alpha2"
	tenantv1alpha1 "kubesphere.io/kubesphere/pkg/client/clientset/versioned/typed/tenant/v1alpha1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ServicemeshV1alpha2() servicemeshv1alpha2.ServicemeshV1alpha2Interface
	// Deprecated: please explicitly pick a version if possible.
	Servicemesh() servicemeshv1alpha2.ServicemeshV1alpha2Interface
	TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Tenant() tenantv1alpha1.TenantV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	servicemeshV1alpha2 *servicemeshv1alpha2.ServicemeshV1alpha2Client
	tenantV1alpha1      *tenantv1alpha1.TenantV1alpha1Client
}

// ServicemeshV1alpha2 retrieves the ServicemeshV1alpha2Client
func (c *Clientset) ServicemeshV1alpha2() servicemeshv1alpha2.ServicemeshV1alpha2Interface {
	return c.servicemeshV1alpha2
}

// Deprecated: Servicemesh retrieves the default version of ServicemeshClient.
// Please explicitly pick a version.
func (c *Clientset) Servicemesh() servicemeshv1alpha2.ServicemeshV1alpha2Interface {
	return c.servicemeshV1alpha2
}

// TenantV1alpha1 retrieves the TenantV1alpha1Client
func (c *Clientset) TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface {
	return c.tenantV1alpha1
}

// Deprecated: Tenant retrieves the default version of TenantClient.
// Please explicitly pick a version.
func (c *Clientset) Tenant() tenantv1alpha1.TenantV1alpha1Interface {
	return c.tenantV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.servicemeshV1alpha2, err = servicemeshv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.tenantV1alpha1, err = tenantv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.servicemeshV1alpha2 = servicemeshv1alpha2.NewForConfigOrDie(c)
	cs.tenantV1alpha1 = tenantv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.servicemeshV1alpha2 = servicemeshv1alpha2.New(c)
	cs.tenantV1alpha1 = tenantv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
