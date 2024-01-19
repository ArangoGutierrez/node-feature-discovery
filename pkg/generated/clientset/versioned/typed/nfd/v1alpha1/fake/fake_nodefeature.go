/*
Copyright 2024 The Kubernetes Authors.

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/node-feature-discovery/pkg/apis/nfd/v1alpha1"
)

// FakeNodeFeatures implements NodeFeatureInterface
type FakeNodeFeatures struct {
	Fake *FakeNfdV1alpha1
	ns   string
}

var nodefeaturesResource = v1alpha1.SchemeGroupVersion.WithResource("nodefeatures")

var nodefeaturesKind = v1alpha1.SchemeGroupVersion.WithKind("NodeFeature")

// Get takes name of the nodeFeature, and returns the corresponding nodeFeature object, and an error if there is any.
func (c *FakeNodeFeatures) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodefeaturesResource, c.ns, name), &v1alpha1.NodeFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeFeature), err
}

// List takes label and field selectors, and returns the list of NodeFeatures that match those selectors.
func (c *FakeNodeFeatures) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeFeatureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodefeaturesResource, nodefeaturesKind, c.ns, opts), &v1alpha1.NodeFeatureList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeFeatureList{ListMeta: obj.(*v1alpha1.NodeFeatureList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeFeatureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeFeatures.
func (c *FakeNodeFeatures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodefeaturesResource, c.ns, opts))

}

// Create takes the representation of a nodeFeature and creates it.  Returns the server's representation of the nodeFeature, and an error, if there is any.
func (c *FakeNodeFeatures) Create(ctx context.Context, nodeFeature *v1alpha1.NodeFeature, opts v1.CreateOptions) (result *v1alpha1.NodeFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodefeaturesResource, c.ns, nodeFeature), &v1alpha1.NodeFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeFeature), err
}

// Update takes the representation of a nodeFeature and updates it. Returns the server's representation of the nodeFeature, and an error, if there is any.
func (c *FakeNodeFeatures) Update(ctx context.Context, nodeFeature *v1alpha1.NodeFeature, opts v1.UpdateOptions) (result *v1alpha1.NodeFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodefeaturesResource, c.ns, nodeFeature), &v1alpha1.NodeFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeFeature), err
}

// Delete takes name of the nodeFeature and deletes it. Returns an error if one occurs.
func (c *FakeNodeFeatures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(nodefeaturesResource, c.ns, name, opts), &v1alpha1.NodeFeature{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeFeatures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodefeaturesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeFeatureList{})
	return err
}

// Patch applies the patch and returns the patched nodeFeature.
func (c *FakeNodeFeatures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodefeaturesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NodeFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeFeature), err
}
