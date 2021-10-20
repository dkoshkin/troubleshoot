/*
Copyright 2019 Replicated, Inc..

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

	v1beta2 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRemoteCollectors implements RemoteCollectorInterface
type FakeRemoteCollectors struct {
	Fake *FakeTroubleshootV1beta2
	ns   string
}

var remotecollectorsResource = schema.GroupVersionResource{Group: "troubleshoot.sh", Version: "v1beta2", Resource: "remotecollectors"}

var remotecollectorsKind = schema.GroupVersionKind{Group: "troubleshoot.sh", Version: "v1beta2", Kind: "RemoteCollector"}

// Get takes name of the remoteCollector, and returns the corresponding remoteCollector object, and an error if there is any.
func (c *FakeRemoteCollectors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.RemoteCollector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(remotecollectorsResource, c.ns, name), &v1beta2.RemoteCollector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.RemoteCollector), err
}

// List takes label and field selectors, and returns the list of RemoteCollectors that match those selectors.
func (c *FakeRemoteCollectors) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.RemoteCollectorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(remotecollectorsResource, remotecollectorsKind, c.ns, opts), &v1beta2.RemoteCollectorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.RemoteCollectorList{ListMeta: obj.(*v1beta2.RemoteCollectorList).ListMeta}
	for _, item := range obj.(*v1beta2.RemoteCollectorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested remoteCollectors.
func (c *FakeRemoteCollectors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(remotecollectorsResource, c.ns, opts))

}

// Create takes the representation of a remoteCollector and creates it.  Returns the server's representation of the remoteCollector, and an error, if there is any.
func (c *FakeRemoteCollectors) Create(ctx context.Context, remoteCollector *v1beta2.RemoteCollector, opts v1.CreateOptions) (result *v1beta2.RemoteCollector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(remotecollectorsResource, c.ns, remoteCollector), &v1beta2.RemoteCollector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.RemoteCollector), err
}

// Update takes the representation of a remoteCollector and updates it. Returns the server's representation of the remoteCollector, and an error, if there is any.
func (c *FakeRemoteCollectors) Update(ctx context.Context, remoteCollector *v1beta2.RemoteCollector, opts v1.UpdateOptions) (result *v1beta2.RemoteCollector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(remotecollectorsResource, c.ns, remoteCollector), &v1beta2.RemoteCollector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.RemoteCollector), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRemoteCollectors) UpdateStatus(ctx context.Context, remoteCollector *v1beta2.RemoteCollector, opts v1.UpdateOptions) (*v1beta2.RemoteCollector, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(remotecollectorsResource, "status", c.ns, remoteCollector), &v1beta2.RemoteCollector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.RemoteCollector), err
}

// Delete takes name of the remoteCollector and deletes it. Returns an error if one occurs.
func (c *FakeRemoteCollectors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(remotecollectorsResource, c.ns, name), &v1beta2.RemoteCollector{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRemoteCollectors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(remotecollectorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.RemoteCollectorList{})
	return err
}

// Patch applies the patch and returns the patched remoteCollector.
func (c *FakeRemoteCollectors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.RemoteCollector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(remotecollectorsResource, c.ns, name, pt, data, subresources...), &v1beta2.RemoteCollector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.RemoteCollector), err
}
