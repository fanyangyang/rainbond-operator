// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	rainbondv1alpha1 "github.com/GLYASAI/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	versioned "github.com/GLYASAI/rainbond-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/GLYASAI/rainbond-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/GLYASAI/rainbond-operator/pkg/generated/listers/rainbond/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RbdComponentInformer provides access to a shared informer and lister for
// RbdComponents.
type RbdComponentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RbdComponentLister
}

type rbdComponentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRbdComponentInformer constructs a new informer for RbdComponent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRbdComponentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRbdComponentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRbdComponentInformer constructs a new informer for RbdComponent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRbdComponentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RainbondV1alpha1().RbdComponents(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RainbondV1alpha1().RbdComponents(namespace).Watch(options)
			},
		},
		&rainbondv1alpha1.RbdComponent{},
		resyncPeriod,
		indexers,
	)
}

func (f *rbdComponentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRbdComponentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rbdComponentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rainbondv1alpha1.RbdComponent{}, f.defaultInformer)
}

func (f *rbdComponentInformer) Lister() v1alpha1.RbdComponentLister {
	return v1alpha1.NewRbdComponentLister(f.Informer().GetIndexer())
}
