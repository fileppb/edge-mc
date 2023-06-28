//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	logicalclusterv1alpha1 "github.com/kubestellar/kubestellar/pkg/apis/logicalcluster/v1alpha1"
	logicalclusterv1alpha1client "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/typed/logicalcluster/v1alpha1"
)

// ClusterProviderDescsClusterGetter has a method to return a ClusterProviderDescClusterInterface.
// A group's cluster client should implement this interface.
type ClusterProviderDescsClusterGetter interface {
	ClusterProviderDescs() ClusterProviderDescClusterInterface
}

// ClusterProviderDescClusterInterface can operate on ClusterProviderDescs across all clusters,
// or scope down to one cluster and return a logicalclusterv1alpha1client.ClusterProviderDescInterface.
type ClusterProviderDescClusterInterface interface {
	Cluster(logicalcluster.Path) logicalclusterv1alpha1client.ClusterProviderDescInterface
	List(ctx context.Context, opts metav1.ListOptions) (*logicalclusterv1alpha1.ClusterProviderDescList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type clusterProviderDescsClusterInterface struct {
	clientCache kcpclient.Cache[*logicalclusterv1alpha1client.LogicalclusterV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterProviderDescsClusterInterface) Cluster(clusterPath logicalcluster.Path) logicalclusterv1alpha1client.ClusterProviderDescInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).ClusterProviderDescs()
}

// List returns the entire collection of all ClusterProviderDescs across all clusters.
func (c *clusterProviderDescsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*logicalclusterv1alpha1.ClusterProviderDescList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterProviderDescs().List(ctx, opts)
}

// Watch begins to watch all ClusterProviderDescs across all clusters.
func (c *clusterProviderDescsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterProviderDescs().Watch(ctx, opts)
}
