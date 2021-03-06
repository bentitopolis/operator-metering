// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/bentitopolis/operator-metering/pkg/apis/metering/v1"
	scheme "github.com/bentitopolis/operator-metering/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MeteringConfigsGetter has a method to return a MeteringConfigInterface.
// A group's client should implement this interface.
type MeteringConfigsGetter interface {
	MeteringConfigs(namespace string) MeteringConfigInterface
}

// MeteringConfigInterface has methods to work with MeteringConfig resources.
type MeteringConfigInterface interface {
	Create(*v1.MeteringConfig) (*v1.MeteringConfig, error)
	Update(*v1.MeteringConfig) (*v1.MeteringConfig, error)
	UpdateStatus(*v1.MeteringConfig) (*v1.MeteringConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MeteringConfig, error)
	List(opts metav1.ListOptions) (*v1.MeteringConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MeteringConfig, err error)
	MeteringConfigExpansion
}

// meteringConfigs implements MeteringConfigInterface
type meteringConfigs struct {
	client rest.Interface
	ns     string
}

// newMeteringConfigs returns a MeteringConfigs
func newMeteringConfigs(c *MeteringV1Client, namespace string) *meteringConfigs {
	return &meteringConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the meteringConfig, and returns the corresponding meteringConfig object, and an error if there is any.
func (c *meteringConfigs) Get(name string, options metav1.GetOptions) (result *v1.MeteringConfig, err error) {
	result = &v1.MeteringConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("meteringconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MeteringConfigs that match those selectors.
func (c *meteringConfigs) List(opts metav1.ListOptions) (result *v1.MeteringConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MeteringConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("meteringconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested meteringConfigs.
func (c *meteringConfigs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("meteringconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a meteringConfig and creates it.  Returns the server's representation of the meteringConfig, and an error, if there is any.
func (c *meteringConfigs) Create(meteringConfig *v1.MeteringConfig) (result *v1.MeteringConfig, err error) {
	result = &v1.MeteringConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("meteringconfigs").
		Body(meteringConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a meteringConfig and updates it. Returns the server's representation of the meteringConfig, and an error, if there is any.
func (c *meteringConfigs) Update(meteringConfig *v1.MeteringConfig) (result *v1.MeteringConfig, err error) {
	result = &v1.MeteringConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("meteringconfigs").
		Name(meteringConfig.Name).
		Body(meteringConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *meteringConfigs) UpdateStatus(meteringConfig *v1.MeteringConfig) (result *v1.MeteringConfig, err error) {
	result = &v1.MeteringConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("meteringconfigs").
		Name(meteringConfig.Name).
		SubResource("status").
		Body(meteringConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the meteringConfig and deletes it. Returns an error if one occurs.
func (c *meteringConfigs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("meteringconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *meteringConfigs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("meteringconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched meteringConfig.
func (c *meteringConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MeteringConfig, err error) {
	result = &v1.MeteringConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("meteringconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
