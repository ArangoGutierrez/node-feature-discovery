/*
Copyright 2019 The Kubernetes Authors.

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

package apihelper

//go:generate mockery --name=APIHelpers --inpackage

import (
	topologyclientset "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/clientset/versioned"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8sclient "k8s.io/client-go/kubernetes"
)

// APIHelpers represents a set of API helpers for Kubernetes
type APIHelpers interface {
	// GetClient returns a client
	GetClient() (*k8sclient.Clientset, error)

	// GetNode returns the Kubernetes node on which this container is running.
	GetNode(*k8sclient.Clientset, string) (*corev1.Node, error)

	// GetNodes returns all the nodes in the cluster
	GetNodes(*k8sclient.Clientset) (*corev1.NodeList, error)

	// GetDaemonSet returns the Kubernetes daemonset in a namepace with a name.
	GetDaemonSet(*k8sclient.Clientset, string) (*appsv1.DaemonSet, error)

	// UpdateDaemonSet updates the daemonset object via the API server using a client.
	UpdateDaemonSet(*k8sclient.Clientset, *appsv1.DaemonSet) error

	// GetDeployment returns the Kubernetes deployment in a namepace with a name.
	GetDeployment(*k8sclient.Clientset, string) (*appsv1.Deployment, error)

	// UpdateDeployment updates the deployment object via the API server using a client.
	UpdateDeployment(*k8sclient.Clientset, *appsv1.Deployment) error

	// UpdateNode updates the node via the API server using a client.
	UpdateNode(*k8sclient.Clientset, *corev1.Node) error

	// PatchNode updates the node object via the API server using a client.
	PatchNode(*k8sclient.Clientset, string, []JsonPatch) error

	// PatchNodeStatus updates the node status via the API server using a client.
	PatchNodeStatus(*k8sclient.Clientset, string, []JsonPatch) error

	// GetTopologyClient returns a topologyclientset
	GetTopologyClient() (*topologyclientset.Clientset, error)

	// GetPod returns the Kubernetes pod in a namepace with a name.
	GetPod(*k8sclient.Clientset, string, string) (*corev1.Pod, error)
}
