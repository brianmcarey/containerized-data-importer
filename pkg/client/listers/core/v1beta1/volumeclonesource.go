/*
Copyright 2018 The CDI Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// VolumeCloneSourceLister helps list VolumeCloneSources.
// All objects returned here must be treated as read-only.
type VolumeCloneSourceLister interface {
	// List lists all VolumeCloneSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VolumeCloneSource, err error)
	// VolumeCloneSources returns an object that can list and get VolumeCloneSources.
	VolumeCloneSources(namespace string) VolumeCloneSourceNamespaceLister
	VolumeCloneSourceListerExpansion
}

// volumeCloneSourceLister implements the VolumeCloneSourceLister interface.
type volumeCloneSourceLister struct {
	indexer cache.Indexer
}

// NewVolumeCloneSourceLister returns a new VolumeCloneSourceLister.
func NewVolumeCloneSourceLister(indexer cache.Indexer) VolumeCloneSourceLister {
	return &volumeCloneSourceLister{indexer: indexer}
}

// List lists all VolumeCloneSources in the indexer.
func (s *volumeCloneSourceLister) List(selector labels.Selector) (ret []*v1beta1.VolumeCloneSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeCloneSource))
	})
	return ret, err
}

// VolumeCloneSources returns an object that can list and get VolumeCloneSources.
func (s *volumeCloneSourceLister) VolumeCloneSources(namespace string) VolumeCloneSourceNamespaceLister {
	return volumeCloneSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeCloneSourceNamespaceLister helps list and get VolumeCloneSources.
// All objects returned here must be treated as read-only.
type VolumeCloneSourceNamespaceLister interface {
	// List lists all VolumeCloneSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VolumeCloneSource, err error)
	// Get retrieves the VolumeCloneSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.VolumeCloneSource, error)
	VolumeCloneSourceNamespaceListerExpansion
}

// volumeCloneSourceNamespaceLister implements the VolumeCloneSourceNamespaceLister
// interface.
type volumeCloneSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VolumeCloneSources in the indexer for a given namespace.
func (s volumeCloneSourceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.VolumeCloneSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeCloneSource))
	})
	return ret, err
}

// Get retrieves the VolumeCloneSource from the indexer for a given namespace and name.
func (s volumeCloneSourceNamespaceLister) Get(name string) (*v1beta1.VolumeCloneSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("volumeclonesource"), name)
	}
	return obj.(*v1beta1.VolumeCloneSource), nil
}
