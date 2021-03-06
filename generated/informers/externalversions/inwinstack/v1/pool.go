/*
Copyright © 2018 inwinSTACK Inc

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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	inwinstackv1 "github.com/inwinstack/blended/apis/inwinstack/v1"
	versioned "github.com/inwinstack/blended/generated/clientset/versioned"
	internalinterfaces "github.com/inwinstack/blended/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/inwinstack/blended/generated/listers/inwinstack/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PoolInformer provides access to a shared informer and lister for
// Pools.
type PoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PoolLister
}

type poolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPoolInformer constructs a new informer for Pool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPoolInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPoolInformer constructs a new informer for Pool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InwinstackV1().Pools().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InwinstackV1().Pools().Watch(options)
			},
		},
		&inwinstackv1.Pool{},
		resyncPeriod,
		indexers,
	)
}

func (f *poolInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPoolInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *poolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&inwinstackv1.Pool{}, f.defaultInformer)
}

func (f *poolInformer) Lister() v1.PoolLister {
	return v1.NewPoolLister(f.Informer().GetIndexer())
}
