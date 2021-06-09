/*

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

package v1alpha2

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "kubesphere.io/fluentbit-operator/api/fluentbitoperator/v1alpha2"
	scheme "kubesphere.io/fluentbit-operator/api/generated/clientset/versioned/scheme"
)

// ParsersGetter has a method to return a ParserInterface.
// A group's client should implement this interface.
type ParsersGetter interface {
	Parsers(namespace string) ParserInterface
}

// ParserInterface has methods to work with Parser resources.
type ParserInterface interface {
	Create(*v1alpha2.Parser) (*v1alpha2.Parser, error)
	Update(*v1alpha2.Parser) (*v1alpha2.Parser, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.Parser, error)
	List(opts v1.ListOptions) (*v1alpha2.ParserList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Parser, err error)
	ParserExpansion
}

// parsers implements ParserInterface
type parsers struct {
	client rest.Interface
	ns     string
}

// newParsers returns a Parsers
func newParsers(c *FluentbitoperatorV1alpha2Client, namespace string) *parsers {
	return &parsers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the parser, and returns the corresponding parser object, and an error if there is any.
func (c *parsers) Get(name string, options v1.GetOptions) (result *v1alpha2.Parser, err error) {
	result = &v1alpha2.Parser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("parsers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Parsers that match those selectors.
func (c *parsers) List(opts v1.ListOptions) (result *v1alpha2.ParserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ParserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("parsers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested parsers.
func (c *parsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("parsers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a parser and creates it.  Returns the server's representation of the parser, and an error, if there is any.
func (c *parsers) Create(parser *v1alpha2.Parser) (result *v1alpha2.Parser, err error) {
	result = &v1alpha2.Parser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("parsers").
		Body(parser).
		Do().
		Into(result)
	return
}

// Update takes the representation of a parser and updates it. Returns the server's representation of the parser, and an error, if there is any.
func (c *parsers) Update(parser *v1alpha2.Parser) (result *v1alpha2.Parser, err error) {
	result = &v1alpha2.Parser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("parsers").
		Name(parser.Name).
		Body(parser).
		Do().
		Into(result)
	return
}

// Delete takes name of the parser and deletes it. Returns an error if one occurs.
func (c *parsers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("parsers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *parsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("parsers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched parser.
func (c *parsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Parser, err error) {
	result = &v1alpha2.Parser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("parsers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}