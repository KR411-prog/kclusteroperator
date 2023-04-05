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

package v1

import (
	"context"
	"time"

	v1 "github.com/KR411-prog/kluster/pkg/apis/radha/v1"
	scheme "github.com/KR411-prog/kluster/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KlustersGetter has a method to return a KlusterInterface.
// A group's client should implement this interface.
type KlustersGetter interface {
	Klusters(namespace string) KlusterInterface
}

// KlusterInterface has methods to work with Kluster resources.
type KlusterInterface interface {
	Create(ctx context.Context, kluster *v1.Kluster, opts v1.CreateOptions) (*v1.Kluster, error)
	Update(ctx context.Context, kluster *v1.Kluster, opts v1.UpdateOptions) (*v1.Kluster, error)
	UpdateStatus(ctx context.Context, kluster *v1.Kluster, opts v1.UpdateOptions) (*v1.Kluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1.Kluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1.KlusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1.Kluster, err error)
	KlusterExpansion
}

// klusters implements KlusterInterface
type klusters struct {
	client rest.Interface
	ns     string
}

// newKlusters returns a Klusters
func newKlusters(c *radhav1Client, namespace string) *klusters {
	return &klusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kluster, and returns the corresponding kluster object, and an error if there is any.
func (c *klusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("klusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Klusters that match those selectors.
func (c *klusters) List(ctx context.Context, opts v1.ListOptions) (result *v1.KlusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KlusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("klusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested klusters.
func (c *klusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("klusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kluster and creates it.  Returns the server's representation of the kluster, and an error, if there is any.
func (c *klusters) Create(ctx context.Context, kluster *v1.Kluster, opts v1.CreateOptions) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("klusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kluster and updates it. Returns the server's representation of the kluster, and an error, if there is any.
func (c *klusters) Update(ctx context.Context, kluster *v1.Kluster, opts v1.UpdateOptions) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("klusters").
		Name(kluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *klusters) UpdateStatus(ctx context.Context, kluster *v1.Kluster, opts v1.UpdateOptions) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("klusters").
		Name(kluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kluster and deletes it. Returns an error if one occurs.
func (c *klusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("klusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *klusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("klusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kluster.
func (c *klusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1.Kluster, err error) {
	result = &v1.Kluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("klusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
