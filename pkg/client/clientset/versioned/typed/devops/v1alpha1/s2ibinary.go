/*
Copyright 2020 The KubeSphere Authors.

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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/devops/v1alpha1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// S2iBinariesGetter has a method to return a S2iBinaryInterface.
// A group's client should implement this interface.
type S2iBinariesGetter interface {
	S2iBinaries(namespace string) S2iBinaryInterface
}

// S2iBinaryInterface has methods to work with S2iBinary resources.
type S2iBinaryInterface interface {
	Create(*v1alpha1.S2iBinary) (*v1alpha1.S2iBinary, error)
	Update(*v1alpha1.S2iBinary) (*v1alpha1.S2iBinary, error)
	UpdateStatus(*v1alpha1.S2iBinary) (*v1alpha1.S2iBinary, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.S2iBinary, error)
	List(opts v1.ListOptions) (*v1alpha1.S2iBinaryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S2iBinary, err error)
	S2iBinaryExpansion
}

// s2iBinaries implements S2iBinaryInterface
type s2iBinaries struct {
	client rest.Interface
	ns     string
}

// newS2iBinaries returns a S2iBinaries
func newS2iBinaries(c *DevopsV1alpha1Client, namespace string) *s2iBinaries {
	return &s2iBinaries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s2iBinary, and returns the corresponding s2iBinary object, and an error if there is any.
func (c *s2iBinaries) Get(name string, options v1.GetOptions) (result *v1alpha1.S2iBinary, err error) {
	result = &v1alpha1.S2iBinary{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s2ibinaries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S2iBinaries that match those selectors.
func (c *s2iBinaries) List(opts v1.ListOptions) (result *v1alpha1.S2iBinaryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S2iBinaryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s2ibinaries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s2iBinaries.
func (c *s2iBinaries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s2ibinaries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a s2iBinary and creates it.  Returns the server's representation of the s2iBinary, and an error, if there is any.
func (c *s2iBinaries) Create(s2iBinary *v1alpha1.S2iBinary) (result *v1alpha1.S2iBinary, err error) {
	result = &v1alpha1.S2iBinary{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s2ibinaries").
		Body(s2iBinary).
		Do().
		Into(result)
	return
}

// Update takes the representation of a s2iBinary and updates it. Returns the server's representation of the s2iBinary, and an error, if there is any.
func (c *s2iBinaries) Update(s2iBinary *v1alpha1.S2iBinary) (result *v1alpha1.S2iBinary, err error) {
	result = &v1alpha1.S2iBinary{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s2ibinaries").
		Name(s2iBinary.Name).
		Body(s2iBinary).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *s2iBinaries) UpdateStatus(s2iBinary *v1alpha1.S2iBinary) (result *v1alpha1.S2iBinary, err error) {
	result = &v1alpha1.S2iBinary{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s2ibinaries").
		Name(s2iBinary.Name).
		SubResource("status").
		Body(s2iBinary).
		Do().
		Into(result)
	return
}

// Delete takes name of the s2iBinary and deletes it. Returns an error if one occurs.
func (c *s2iBinaries) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s2ibinaries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s2iBinaries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s2ibinaries").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched s2iBinary.
func (c *s2iBinaries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S2iBinary, err error) {
	result = &v1alpha1.S2iBinary{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s2ibinaries").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
