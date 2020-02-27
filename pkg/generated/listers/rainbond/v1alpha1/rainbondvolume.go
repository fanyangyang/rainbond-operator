// RAINBOND, Application Management Platform
// Copyright (C) 2014-2020 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/goodrain/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RainbondVolumeLister helps list RainbondVolumes.
type RainbondVolumeLister interface {
	// List lists all RainbondVolumes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RainbondVolume, err error)
	// RainbondVolumes returns an object that can list and get RainbondVolumes.
	RainbondVolumes(namespace string) RainbondVolumeNamespaceLister
	RainbondVolumeListerExpansion
}

// rainbondVolumeLister implements the RainbondVolumeLister interface.
type rainbondVolumeLister struct {
	indexer cache.Indexer
}

// NewRainbondVolumeLister returns a new RainbondVolumeLister.
func NewRainbondVolumeLister(indexer cache.Indexer) RainbondVolumeLister {
	return &rainbondVolumeLister{indexer: indexer}
}

// List lists all RainbondVolumes in the indexer.
func (s *rainbondVolumeLister) List(selector labels.Selector) (ret []*v1alpha1.RainbondVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RainbondVolume))
	})
	return ret, err
}

// RainbondVolumes returns an object that can list and get RainbondVolumes.
func (s *rainbondVolumeLister) RainbondVolumes(namespace string) RainbondVolumeNamespaceLister {
	return rainbondVolumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RainbondVolumeNamespaceLister helps list and get RainbondVolumes.
type RainbondVolumeNamespaceLister interface {
	// List lists all RainbondVolumes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RainbondVolume, err error)
	// Get retrieves the RainbondVolume from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RainbondVolume, error)
	RainbondVolumeNamespaceListerExpansion
}

// rainbondVolumeNamespaceLister implements the RainbondVolumeNamespaceLister
// interface.
type rainbondVolumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RainbondVolumes in the indexer for a given namespace.
func (s rainbondVolumeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RainbondVolume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RainbondVolume))
	})
	return ret, err
}

// Get retrieves the RainbondVolume from the indexer for a given namespace and name.
func (s rainbondVolumeNamespaceLister) Get(name string) (*v1alpha1.RainbondVolume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rainbondvolume"), name)
	}
	return obj.(*v1alpha1.RainbondVolume), nil
}