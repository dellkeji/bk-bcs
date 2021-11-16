/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-general-pod-autoscaler/pkg/apis/autoscaling/v1alpha1"
	scheme "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-general-pod-autoscaler/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GeneralPodAutoscalersGetter has a method to return a GeneralPodAutoscalerInterface.
// A group's client should implement this interface.
type GeneralPodAutoscalersGetter interface {
	GeneralPodAutoscalers(namespace string) GeneralPodAutoscalerInterface
}

// GeneralPodAutoscalerInterface has methods to work with GeneralPodAutoscaler resources.
type GeneralPodAutoscalerInterface interface {
	Create(*v1alpha1.GeneralPodAutoscaler) (*v1alpha1.GeneralPodAutoscaler, error)
	Update(*v1alpha1.GeneralPodAutoscaler) (*v1alpha1.GeneralPodAutoscaler, error)
	UpdateStatus(*v1alpha1.GeneralPodAutoscaler) (*v1alpha1.GeneralPodAutoscaler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GeneralPodAutoscaler, error)
	List(opts v1.ListOptions) (*v1alpha1.GeneralPodAutoscalerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GeneralPodAutoscaler, err error)
	GeneralPodAutoscalerExpansion
}

// generalPodAutoscalers implements GeneralPodAutoscalerInterface
type generalPodAutoscalers struct {
	client rest.Interface
	ns     string
}

// newGeneralPodAutoscalers returns a GeneralPodAutoscalers
func newGeneralPodAutoscalers(c *AutoscalingV1alpha1Client, namespace string) *generalPodAutoscalers {
	return &generalPodAutoscalers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the generalPodAutoscaler, and returns the corresponding generalPodAutoscaler object, and an error if there is any.
func (c *generalPodAutoscalers) Get(name string, options v1.GetOptions) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	result = &v1alpha1.GeneralPodAutoscaler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return result, err
}

// List takes label and field selectors, and returns the list of GeneralPodAutoscalers that match those selectors.
func (c *generalPodAutoscalers) List(opts v1.ListOptions) (result *v1alpha1.GeneralPodAutoscalerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GeneralPodAutoscalerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return result, err
}

// Watch returns a watch.Interface that watches the requested generalPodAutoscalers.
func (c *generalPodAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a generalPodAutoscaler and creates it.  Returns the server's representation of the generalPodAutoscaler, and an error, if there is any.
func (c *generalPodAutoscalers) Create(generalPodAutoscaler *v1alpha1.GeneralPodAutoscaler) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	result = &v1alpha1.GeneralPodAutoscaler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		Body(generalPodAutoscaler).
		Do().
		Into(result)
	return result, err
}

// Update takes the representation of a generalPodAutoscaler and updates it. Returns the server's representation of the generalPodAutoscaler, and an error, if there is any.
func (c *generalPodAutoscalers) Update(generalPodAutoscaler *v1alpha1.GeneralPodAutoscaler) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	result = &v1alpha1.GeneralPodAutoscaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		Name(generalPodAutoscaler.Name).
		Body(generalPodAutoscaler).
		Do().
		Into(result)
	return result, err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *generalPodAutoscalers) UpdateStatus(generalPodAutoscaler *v1alpha1.GeneralPodAutoscaler) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	result = &v1alpha1.GeneralPodAutoscaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		Name(generalPodAutoscaler.Name).
		SubResource("status").
		Body(generalPodAutoscaler).
		Do().
		Into(result)
	return result, err
}

// Delete takes name of the generalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *generalPodAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *generalPodAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched generalPodAutoscaler.
func (c *generalPodAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	result = &v1alpha1.GeneralPodAutoscaler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("generalpodautoscalers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return result, err
}