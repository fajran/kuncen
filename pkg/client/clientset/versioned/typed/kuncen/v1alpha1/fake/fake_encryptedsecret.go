// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/fajran/kuncen/pkg/apis/kuncen/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEncryptedSecrets implements EncryptedSecretInterface
type FakeEncryptedSecrets struct {
	Fake *FakeKuncenV1alpha1
	ns   string
}

var encryptedsecretsResource = schema.GroupVersionResource{Group: "kuncen.fajran.github.io", Version: "v1alpha1", Resource: "encryptedsecrets"}

var encryptedsecretsKind = schema.GroupVersionKind{Group: "kuncen.fajran.github.io", Version: "v1alpha1", Kind: "EncryptedSecret"}

// Get takes name of the encryptedSecret, and returns the corresponding encryptedSecret object, and an error if there is any.
func (c *FakeEncryptedSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.EncryptedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(encryptedsecretsResource, c.ns, name), &v1alpha1.EncryptedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptedSecret), err
}

// List takes label and field selectors, and returns the list of EncryptedSecrets that match those selectors.
func (c *FakeEncryptedSecrets) List(opts v1.ListOptions) (result *v1alpha1.EncryptedSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(encryptedsecretsResource, encryptedsecretsKind, c.ns, opts), &v1alpha1.EncryptedSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EncryptedSecretList{ListMeta: obj.(*v1alpha1.EncryptedSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.EncryptedSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested encryptedSecrets.
func (c *FakeEncryptedSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(encryptedsecretsResource, c.ns, opts))

}

// Create takes the representation of a encryptedSecret and creates it.  Returns the server's representation of the encryptedSecret, and an error, if there is any.
func (c *FakeEncryptedSecrets) Create(encryptedSecret *v1alpha1.EncryptedSecret) (result *v1alpha1.EncryptedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(encryptedsecretsResource, c.ns, encryptedSecret), &v1alpha1.EncryptedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptedSecret), err
}

// Update takes the representation of a encryptedSecret and updates it. Returns the server's representation of the encryptedSecret, and an error, if there is any.
func (c *FakeEncryptedSecrets) Update(encryptedSecret *v1alpha1.EncryptedSecret) (result *v1alpha1.EncryptedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(encryptedsecretsResource, c.ns, encryptedSecret), &v1alpha1.EncryptedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptedSecret), err
}

// Delete takes name of the encryptedSecret and deletes it. Returns an error if one occurs.
func (c *FakeEncryptedSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(encryptedsecretsResource, c.ns, name), &v1alpha1.EncryptedSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEncryptedSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(encryptedsecretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EncryptedSecretList{})
	return err
}

// Patch applies the patch and returns the patched encryptedSecret.
func (c *FakeEncryptedSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EncryptedSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(encryptedsecretsResource, c.ns, name, data, subresources...), &v1alpha1.EncryptedSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptedSecret), err
}
