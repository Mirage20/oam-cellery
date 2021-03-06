/*
 * Copyright (c) 2020 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "cellery.io/oam-cellery/pkg/apis/core/v1alpha1"
	scheme "cellery.io/oam-cellery/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComponentSchematicsGetter has a method to return a ComponentSchematicInterface.
// A group's client should implement this interface.
type ComponentSchematicsGetter interface {
	ComponentSchematics(namespace string) ComponentSchematicInterface
}

// ComponentSchematicInterface has methods to work with ComponentSchematic resources.
type ComponentSchematicInterface interface {
	Create(*v1alpha1.ComponentSchematic) (*v1alpha1.ComponentSchematic, error)
	Update(*v1alpha1.ComponentSchematic) (*v1alpha1.ComponentSchematic, error)
	UpdateStatus(*v1alpha1.ComponentSchematic) (*v1alpha1.ComponentSchematic, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComponentSchematic, error)
	List(opts v1.ListOptions) (*v1alpha1.ComponentSchematicList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComponentSchematic, err error)
	ComponentSchematicExpansion
}

// componentSchematics implements ComponentSchematicInterface
type componentSchematics struct {
	client rest.Interface
	ns     string
}

// newComponentSchematics returns a ComponentSchematics
func newComponentSchematics(c *CoreV1alpha1Client, namespace string) *componentSchematics {
	return &componentSchematics{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the componentSchematic, and returns the corresponding componentSchematic object, and an error if there is any.
func (c *componentSchematics) Get(name string, options v1.GetOptions) (result *v1alpha1.ComponentSchematic, err error) {
	result = &v1alpha1.ComponentSchematic{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("componentschematics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComponentSchematics that match those selectors.
func (c *componentSchematics) List(opts v1.ListOptions) (result *v1alpha1.ComponentSchematicList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComponentSchematicList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("componentschematics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested componentSchematics.
func (c *componentSchematics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("componentschematics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a componentSchematic and creates it.  Returns the server's representation of the componentSchematic, and an error, if there is any.
func (c *componentSchematics) Create(componentSchematic *v1alpha1.ComponentSchematic) (result *v1alpha1.ComponentSchematic, err error) {
	result = &v1alpha1.ComponentSchematic{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("componentschematics").
		Body(componentSchematic).
		Do().
		Into(result)
	return
}

// Update takes the representation of a componentSchematic and updates it. Returns the server's representation of the componentSchematic, and an error, if there is any.
func (c *componentSchematics) Update(componentSchematic *v1alpha1.ComponentSchematic) (result *v1alpha1.ComponentSchematic, err error) {
	result = &v1alpha1.ComponentSchematic{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("componentschematics").
		Name(componentSchematic.Name).
		Body(componentSchematic).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *componentSchematics) UpdateStatus(componentSchematic *v1alpha1.ComponentSchematic) (result *v1alpha1.ComponentSchematic, err error) {
	result = &v1alpha1.ComponentSchematic{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("componentschematics").
		Name(componentSchematic.Name).
		SubResource("status").
		Body(componentSchematic).
		Do().
		Into(result)
	return
}

// Delete takes name of the componentSchematic and deletes it. Returns an error if one occurs.
func (c *componentSchematics) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("componentschematics").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *componentSchematics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("componentschematics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched componentSchematic.
func (c *componentSchematics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComponentSchematic, err error) {
	result = &v1alpha1.ComponentSchematic{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("componentschematics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
