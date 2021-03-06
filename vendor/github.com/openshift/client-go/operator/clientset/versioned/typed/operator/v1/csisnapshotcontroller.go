// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/api/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CSISnapshotControllersGetter has a method to return a CSISnapshotControllerInterface.
// A group's client should implement this interface.
type CSISnapshotControllersGetter interface {
	CSISnapshotControllers() CSISnapshotControllerInterface
}

// CSISnapshotControllerInterface has methods to work with CSISnapshotController resources.
type CSISnapshotControllerInterface interface {
	Create(*v1.CSISnapshotController) (*v1.CSISnapshotController, error)
	Update(*v1.CSISnapshotController) (*v1.CSISnapshotController, error)
	UpdateStatus(*v1.CSISnapshotController) (*v1.CSISnapshotController, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.CSISnapshotController, error)
	List(opts metav1.ListOptions) (*v1.CSISnapshotControllerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CSISnapshotController, err error)
	CSISnapshotControllerExpansion
}

// cSISnapshotControllers implements CSISnapshotControllerInterface
type cSISnapshotControllers struct {
	client rest.Interface
}

// newCSISnapshotControllers returns a CSISnapshotControllers
func newCSISnapshotControllers(c *OperatorV1Client) *cSISnapshotControllers {
	return &cSISnapshotControllers{
		client: c.RESTClient(),
	}
}

// Get takes name of the cSISnapshotController, and returns the corresponding cSISnapshotController object, and an error if there is any.
func (c *cSISnapshotControllers) Get(name string, options metav1.GetOptions) (result *v1.CSISnapshotController, err error) {
	result = &v1.CSISnapshotController{}
	err = c.client.Get().
		Resource("csisnapshotcontrollers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CSISnapshotControllers that match those selectors.
func (c *cSISnapshotControllers) List(opts metav1.ListOptions) (result *v1.CSISnapshotControllerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CSISnapshotControllerList{}
	err = c.client.Get().
		Resource("csisnapshotcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cSISnapshotControllers.
func (c *cSISnapshotControllers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("csisnapshotcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cSISnapshotController and creates it.  Returns the server's representation of the cSISnapshotController, and an error, if there is any.
func (c *cSISnapshotControllers) Create(cSISnapshotController *v1.CSISnapshotController) (result *v1.CSISnapshotController, err error) {
	result = &v1.CSISnapshotController{}
	err = c.client.Post().
		Resource("csisnapshotcontrollers").
		Body(cSISnapshotController).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cSISnapshotController and updates it. Returns the server's representation of the cSISnapshotController, and an error, if there is any.
func (c *cSISnapshotControllers) Update(cSISnapshotController *v1.CSISnapshotController) (result *v1.CSISnapshotController, err error) {
	result = &v1.CSISnapshotController{}
	err = c.client.Put().
		Resource("csisnapshotcontrollers").
		Name(cSISnapshotController.Name).
		Body(cSISnapshotController).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cSISnapshotControllers) UpdateStatus(cSISnapshotController *v1.CSISnapshotController) (result *v1.CSISnapshotController, err error) {
	result = &v1.CSISnapshotController{}
	err = c.client.Put().
		Resource("csisnapshotcontrollers").
		Name(cSISnapshotController.Name).
		SubResource("status").
		Body(cSISnapshotController).
		Do().
		Into(result)
	return
}

// Delete takes name of the cSISnapshotController and deletes it. Returns an error if one occurs.
func (c *cSISnapshotControllers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("csisnapshotcontrollers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cSISnapshotControllers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("csisnapshotcontrollers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cSISnapshotController.
func (c *cSISnapshotControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CSISnapshotController, err error) {
	result = &v1.CSISnapshotController{}
	err = c.client.Patch(pt).
		Resource("csisnapshotcontrollers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
