// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/fajran/kuncen/pkg/apis/kuncen/v1alpha1"
	scheme "github.com/fajran/kuncen/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EncryptedSecretsGetter has a method to return a EncryptedSecretInterface.
// A group's client should implement this interface.
type EncryptedSecretsGetter interface {
	EncryptedSecrets(namespace string) EncryptedSecretInterface
}

// EncryptedSecretInterface has methods to work with EncryptedSecret resources.
type EncryptedSecretInterface interface {
	Create(*v1alpha1.EncryptedSecret) (*v1alpha1.EncryptedSecret, error)
	Update(*v1alpha1.EncryptedSecret) (*v1alpha1.EncryptedSecret, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EncryptedSecret, error)
	List(opts v1.ListOptions) (*v1alpha1.EncryptedSecretList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EncryptedSecret, err error)
	EncryptedSecretExpansion
}

// encryptedSecrets implements EncryptedSecretInterface
type encryptedSecrets struct {
	client rest.Interface
	ns     string
}

// newEncryptedSecrets returns a EncryptedSecrets
func newEncryptedSecrets(c *KuncenV1alpha1Client, namespace string) *encryptedSecrets {
	return &encryptedSecrets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the encryptedSecret, and returns the corresponding encryptedSecret object, and an error if there is any.
func (c *encryptedSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.EncryptedSecret, err error) {
	result = &v1alpha1.EncryptedSecret{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("encryptedsecrets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EncryptedSecrets that match those selectors.
func (c *encryptedSecrets) List(opts v1.ListOptions) (result *v1alpha1.EncryptedSecretList, err error) {
	result = &v1alpha1.EncryptedSecretList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("encryptedsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested encryptedSecrets.
func (c *encryptedSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("encryptedsecrets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a encryptedSecret and creates it.  Returns the server's representation of the encryptedSecret, and an error, if there is any.
func (c *encryptedSecrets) Create(encryptedSecret *v1alpha1.EncryptedSecret) (result *v1alpha1.EncryptedSecret, err error) {
	result = &v1alpha1.EncryptedSecret{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("encryptedsecrets").
		Body(encryptedSecret).
		Do().
		Into(result)
	return
}

// Update takes the representation of a encryptedSecret and updates it. Returns the server's representation of the encryptedSecret, and an error, if there is any.
func (c *encryptedSecrets) Update(encryptedSecret *v1alpha1.EncryptedSecret) (result *v1alpha1.EncryptedSecret, err error) {
	result = &v1alpha1.EncryptedSecret{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("encryptedsecrets").
		Name(encryptedSecret.Name).
		Body(encryptedSecret).
		Do().
		Into(result)
	return
}

// Delete takes name of the encryptedSecret and deletes it. Returns an error if one occurs.
func (c *encryptedSecrets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("encryptedsecrets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *encryptedSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("encryptedsecrets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched encryptedSecret.
func (c *encryptedSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EncryptedSecret, err error) {
	result = &v1alpha1.EncryptedSecret{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("encryptedsecrets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
