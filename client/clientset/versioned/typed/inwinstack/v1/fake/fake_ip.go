/*
Copyright © 2018 Kyle Bai(kyle.b@inwinstack.com)

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
	inwinstackv1 "github.com/inwinstack/ipam/apis/inwinstack/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIPs implements IPInterface
type FakeIPs struct {
	Fake *FakeInwinstackV1
	ns   string
}

var ipsResource = schema.GroupVersionResource{Group: "inwinstack.com", Version: "v1", Resource: "ips"}

var ipsKind = schema.GroupVersionKind{Group: "inwinstack.com", Version: "v1", Kind: "IP"}

// Get takes name of the iP, and returns the corresponding iP object, and an error if there is any.
func (c *FakeIPs) Get(name string, options v1.GetOptions) (result *inwinstackv1.IP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipsResource, c.ns, name), &inwinstackv1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inwinstackv1.IP), err
}

// List takes label and field selectors, and returns the list of IPs that match those selectors.
func (c *FakeIPs) List(opts v1.ListOptions) (result *inwinstackv1.IPList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipsResource, ipsKind, c.ns, opts), &inwinstackv1.IPList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &inwinstackv1.IPList{ListMeta: obj.(*inwinstackv1.IPList).ListMeta}
	for _, item := range obj.(*inwinstackv1.IPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPs.
func (c *FakeIPs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipsResource, c.ns, opts))

}

// Create takes the representation of a iP and creates it.  Returns the server's representation of the iP, and an error, if there is any.
func (c *FakeIPs) Create(iP *inwinstackv1.IP) (result *inwinstackv1.IP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipsResource, c.ns, iP), &inwinstackv1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inwinstackv1.IP), err
}

// Update takes the representation of a iP and updates it. Returns the server's representation of the iP, and an error, if there is any.
func (c *FakeIPs) Update(iP *inwinstackv1.IP) (result *inwinstackv1.IP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipsResource, c.ns, iP), &inwinstackv1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inwinstackv1.IP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIPs) UpdateStatus(iP *inwinstackv1.IP) (*inwinstackv1.IP, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipsResource, "status", c.ns, iP), &inwinstackv1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inwinstackv1.IP), err
}

// Delete takes name of the iP and deletes it. Returns an error if one occurs.
func (c *FakeIPs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ipsResource, c.ns, name), &inwinstackv1.IP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &inwinstackv1.IPList{})
	return err
}

// Patch applies the patch and returns the patched iP.
func (c *FakeIPs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *inwinstackv1.IP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipsResource, c.ns, name, data, subresources...), &inwinstackv1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*inwinstackv1.IP), err
}
