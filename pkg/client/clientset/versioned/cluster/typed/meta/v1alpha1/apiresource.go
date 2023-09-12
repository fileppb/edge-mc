/*
Copyright The KubeStellar Authors.

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


//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	metav1alpha1 "github.com/kubestellar/kubestellar/pkg/apis/meta/v1alpha1"


	metav1alpha1client "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/typed/meta/v1alpha1"
)

// APIResourcesClusterGetter has a method to return a APIResourceClusterInterface.
// A group's cluster client should implement this interface.
type APIResourcesClusterGetter interface {
	APIResources() APIResourceClusterInterface
}

// APIResourceClusterInterface can operate on APIResources across all clusters,
// or scope down to one cluster and return a metav1alpha1client.APIResourceInterface.
type APIResourceClusterInterface interface {
	Cluster(logicalcluster.Path) metav1alpha1client.APIResourceInterface
	List(ctx context.Context, opts metav1.ListOptions) (*metav1alpha1.APIResourceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type aPIResourcesClusterInterface struct {
	clientCache kcpclient.Cache[*metav1alpha1client.MetaV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *aPIResourcesClusterInterface) Cluster(clusterPath logicalcluster.Path) metav1alpha1client.APIResourceInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).APIResources()
}


// List returns the entire collection of all APIResources across all clusters. 
func (c *aPIResourcesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*metav1alpha1.APIResourceList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).APIResources().List(ctx, opts)
}

// Watch begins to watch all APIResources across all clusters.
func (c *aPIResourcesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).APIResources().Watch(ctx, opts)
}
